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
	"github.com/ainsleyclark/verbis/api/logger"
	"sync"
	"time"
)

// MigrationInfo represents the current state of the
// migration.
type MigrationInfo struct {
	// The total amount of files being process.
	Total int `json:"total"`
	// The percentage amount of total files processed.
	Progress int `json:"progress"`
	// How many files have succeeded migration.
	Succeeded int `json:"succeeded"`
	// How many files have failed migration.
	Failed int `json:"failed"`
	// How many files have already been processed (succeeded + failed)
	FilesProcessed int `json:"files_processed"`
	// When the migration started.
	MigratedAt time.Time `json:"migrated_at"`
	// Any errors that have occurred during the migration.
	Errors []FailedMigrationFile `json:"errors"`
}

// FailedMigrationFile represents an error when migrating.
// It includes an error.Error as well as a file for
// debugging.
type FailedMigrationFile struct {
	Error *errors.Error `json:"error"`
	File  domain.File   `json:"file"`
}

var (
	// ErrAlreadyMigrating is returned by Migrate() when a
	// migration is already in progress.
	ErrAlreadyMigrating = errors.New("migration is already in progress")
	// migrateTrackChan is the channel used for sending and
	// processing migrations.
	migrateTrackChan = make(chan migration, 20)
	// migrateWg is the wait group for migrations.
	migrateWg = sync.WaitGroup{}
)

// fail appends an error to the migration stack and adds
// one to failed files and files processed retrospectively.
func (m *MigrationInfo) fail(file domain.File, err error) {
	m.Failed += 1
	m.FilesProcessed += 1
	m.Errors = append(m.Errors, FailedMigrationFile{
		Error: errors.ToError(err),
		File:  file,
	})
	m.storeMigration()
	logger.WithError(err).Error()
}

// succeed adds a succeeded file to the migration stack as
// well as adding one to the files processed.
func (m *MigrationInfo) succeed(file domain.File) {
	m.Succeeded += 1
	m.FilesProcessed += 1
	m.storeMigration()
	logger.Info("Successfully migrated file: " + file.Name)
}

// calculateProcessed
func (m *MigrationInfo) storeMigration() {
	// TODO - we need to override the cache with the updated MigrationInfo here!
	m.Progress = (m.FilesProcessed * 100) / m.Total
}

// migration is an entity used to help to process file
// migrations.
type migration struct {
	file domain.File
	from domain.StorageChange
	to   domain.StorageChange
}

// Migrate satisfies the Provider interface by accepting a
// from and to StorageChange to migrate files to the
// remote provider or local storage.
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
		return 0, &errors.Error{Code: errors.INVALID, Message: "Validation failed", Operation: op, Err: err}
	}

	filters := params.Filters{
		"provider": {
			{Operator: "=", Value: from.Provider.String()},
		},
		"bucket": {
			{Operator: "=", Value: from.Bucket},
		},
	}

	ff, total, err := s.filesRepo.List(params.Params{LimitAll: true, Filters: filters})
	if err != nil {
		return 0, err
	}

	// TODO, this needs to be stored in the cache, if there are multiple
	// migrations on a stateless platform, there will be
	// inconsistencies
	s.isMigrating = true
	s.migration = MigrationInfo{
		Total:      total,
		MigratedAt: time.Now(),
	}

	go s.processMigration(ff, from, to)

	return total, nil
}

// processMigration ranges over the given files and adds a
// migration to the migrateTrackChan.
func (s *Storage) processMigration(files domain.Files, from, to domain.StorageChange) {
	for _, file := range files {
		migrateTrackChan <- migration{
			file: file,
			from: from,
			to:   to,
		}
		go s.migrateBackground()
	}

	migrateWg.Wait()
	s.isMigrating = false

	logger.Info(fmt.Sprintf("Storage: %d files migrated successfully", s.migration.Succeeded))
	logger.Info(fmt.Sprintf("Storage: %d files encountered an error during migration", s.migration.Failed))
}

// migrateBackground processes the migration by finding the
// original bytes, uploading to the new destination
// and deleting the original file.
func (s *Storage) migrateBackground() {
	migrateWg.Add(1)
	m := <-migrateTrackChan

	defer func() {
		migrateWg.Done()
	}()

	buf, _, err := s.Find(m.file.Url)
	if err != nil {
		s.migration.fail(m.file, err)
		return
	}

	u := domain.Upload{
		UUID:       m.file.UUID,
		Path:       m.file.Url,
		Size:       m.file.FileSize,
		Contents:   bytes.NewReader(buf),
		Private:    bool(m.file.Private),
		SourceType: m.file.SourceType,
	}

	_, err = s.upload(m.to.Provider, m.to.Bucket, u)
	if err != nil {
		s.migration.fail(m.file, err)
		return
	}

	err = s.Delete(m.file.Id)
	if err != nil {
		s.migration.fail(m.file, err)
		return
	}

	s.migration.succeed(m.file)
}
