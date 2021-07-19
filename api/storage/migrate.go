// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/common/params"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
	"github.com/gookit/color"
	"sync"
)

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
type MigrationInfo struct {
	Total          int                   `json:"total"`
	Progress       int                   `json:"progress"`
	Failed         int                   `json:"failed"`
	Succeeded      int                   `json:"succeeded"`
	FilesProcessed int                   `json:"files_processed"`
	Errors         []FailedMigrationFile `json:"errors"`
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
type FailedMigrationFile struct {
	Error *errors.Error `json:"error"`
	File  domain.File   `json:"file"`
}

var (
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	ErrAlreadyMigrating = errors.New("migration is already in progress")
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	migrateTrackChan = make(chan int, 20)
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	migrateWg = sync.WaitGroup{}
)

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
func (m *MigrationInfo) fail(file domain.File, err error) {
	m.Failed += 1
	m.FilesProcessed += 1
	// NIL POINTER
	m.Errors = append(m.Errors, FailedMigrationFile{
		Error: errors.ToError(err),
		File:  file,
	})
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
func (m *MigrationInfo) succeed() {
	m.Succeeded += 1
	m.FilesProcessed += 1
	m.Progress = (m.FilesProcessed * 100) / m.Total
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
func (s *Storage) Migrate(from, to domain.StorageChange) (int, error) {
	const op = "Storage.Migrate"

	if s.isMigrating {
		return 0, &errors.Error{Code: errors.INVALID, Message: "Error migration is already in progress", Operation: op, Err: ErrAlreadyMigrating}
	}

	if from.Provider == to.Provider {
		return 0, &errors.Error{Code: errors.INVALID, Message: "Error providers cannot be the same", Operation: op, Err: fmt.Errorf("providers are the same")}
	}

	err := s.validate(to)
	if err != nil {
		color.Red.Println(err)
		return 0, &errors.Error{Code: errors.INVALID, Message: "Validation failed", Operation: op, Err: err}
	}

	ff, total, err := s.filesRepo.List(params.Params{LimitAll: false})
	if err != nil {
		return 0, err
	}

	s.isMigrating = true
	s.migration = MigrationInfo{
		Total: total,
	}

	go func() {
		for _, file := range ff {
			migrateTrackChan <- 1
			go s.migrateBackground(file, from, to)
		}
	}()

	return total, nil
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
func (s *Storage) migrateBackground(file domain.File, from, to domain.StorageChange) {
	migrateWg.Add(1)
	defer func() {
		migrateWg.Done()
		<-migrateTrackChan
	}()

	if to.Provider == file.Provider {
		return
	}

	buf, _, err := s.Find(file.Url)
	if err != nil {
		s.migration.fail(file, err)
		return
	}

	u := domain.Upload{
		UUID:       uuid.New(),
		Path:       file.Url,
		Size:       file.FileSize,
		Contents:   bytes.NewReader(buf),
		Private:    bool(file.Private),
		SourceType: file.SourceType,
	}

	_, err = s.upload(to.Provider, to.Bucket, u)
	if err != nil {
		s.migration.fail(file, err)
		return
	}

	err = s.filesRepo.Update(file.Id, to)
	if err != nil {
		s.migration.fail(file, err)
		return
	}

	err = s.Delete(file.Id)
	if err != nil {
		return
	}

	s.migration.succeed()
}
