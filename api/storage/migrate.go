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
	"sync"
)


// Need to add JSON info here!, Perhaps move it to domain?
// Makes it easier for storage config instead of the
// interface{} were using at the moment.



type MigrationInfo struct {
	Total          int
	Progress       int
	Failed         int
	Succeeded      int
	FilesProcessed int
	Errors         []FailedMigrationFile
}

type FailedMigrationFile struct {
	Error error
	File  domain.File
}

var (
	ErrAlreadyMigrating = errors.New("migration is already in progress")
	migrateTrackChan    = make(chan int, 20)
	migrateWg           = sync.WaitGroup{}
)

func (m *MigrationInfo) fail(file domain.File, err error) {
	m.Failed += 1
	m.FilesProcessed += 1
	m.Errors = append(m.Errors, FailedMigrationFile{
		Error: err,
		File:  file,
	})
}

func (m *MigrationInfo) succeed() {
	m.Succeeded += 1
	m.FilesProcessed += 1
	m.Progress = (m.FilesProcessed * 100) / m.Total
}

func (s *Storage) Migrate(from, to domain.StorageChange) (int, error) {
	const op = "Storage.Migrate"

	if s.isMigrating {
		return 0, &errors.Error{Code: errors.INVALID, Message: "Error migration is already in progress", Operation: op, Err: ErrAlreadyMigrating}
	}

	if from.Provider == to.Provider {
		return 0, &errors.Error{Code: errors.INVALID, Message: "Error providers cannot be the same", Operation: op, Err: fmt.Errorf("providers are the same")}
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

func (s *Storage) migrateBackground(file domain.File, from, to domain.StorageChange) {
	migrateWg.Add(1)
	defer func() {
		migrateWg.Done()
		<-migrateTrackChan
	}()

	if from.Provider == file.Provider {
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
		Size:       0,
		Contents:   bytes.NewReader(buf),
		Private:    bool(file.Private),
		SourceType: file.SourceType,
	}

	_, err = s.upload(from.Provider, from.Bucket, u)
	if err != nil {
		s.migration.fail(file, err)
		return
	}

	err = s.filesRepo.Update(file.Id, to)
	if err != nil {
		s.migration.fail(file, err)
		return
	}

	s.migration.succeed()
}
