// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"bytes"
	"context"
	"fmt"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/common/params"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
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
	// Avoid data race for go routine
	mtx *sync.Mutex
}

// FailedMigrationFile represents an error when migrating.
// It includes an error.Error as well as a file for
// debugging.
type FailedMigrationFile struct {
	Error *errors.Error `json:"error"`
	File  domain.File   `json:"file"`
}

const (
	// migrateConcurrentAllowance is the amount of files that
	// are allowed to be migrated concurrently.
	migrateConcurrentAllowance = 1
	// migrationKey is the key used in the cache used for
	// retrieving migration information.
	migrationKey = "storage_migration"
	// migrationIsMigrating is the key used in the cache used for
	// determining if there is a migration taking place.
	migrationIsMigrating = "storage_is_migrating"
)

var (
	// ErrAlreadyMigrating is returned by Migrate() when a
	// migration is already in progress.
	ErrAlreadyMigrating = errors.New("migration is already in progress")
	// ErrNoFilesToMigrate is returned by Migrate() when no
	// files have been found to process.
	ErrNoFilesToMigrate = errors.New("no files to migrate")
)

// fail appends an error to the migration stack and adds
// one to failed files and files processed retrospectively.
func (m *MigrationInfo) fail(file domain.File, err error) {
	m.Failed++
	m.FilesProcessed++
	m.Errors = append(m.Errors, FailedMigrationFile{
		Error: errors.ToError(err),
		File:  file,
	})
	m.Progress = (m.FilesProcessed * 100) / m.Total
	logger.WithError(err).Error()
}

// succeed adds a succeeded file to the migration stack as
// well as adding one to the files processed.
func (m *MigrationInfo) succeed(file domain.File) {
	m.Succeeded++
	m.FilesProcessed++
	m.Progress = (m.FilesProcessed * 100) / m.Total
	logger.Debug("Successfully migrated file: " + file.Name)
}

// migration is an entity used to help to process file
// migrations.
type migration struct {
	file domain.File
	from domain.StorageChange
	to   domain.StorageChange
	wg   *sync.WaitGroup
}

// Migrate satisfies the Provider interface by accepting a
// from and to StorageChange to migrate files to the
// remote provider or local storage.
func (s *Storage) Migrate(ctx context.Context, from, to domain.StorageChange, deleteFiles bool) (int, error) {
	const op = "Storage.Migrate"

	if s.isMigrating(ctx) {
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

	if total == 0 {
		return 0, &errors.Error{Code: errors.NOTFOUND, Message: "Error no files found with provider: " + from.Provider.String(), Operation: op, Err: ErrNoFilesToMigrate}
	}

	logger.Debug(fmt.Sprintf("Starting storage migration with %d files being processed", total))

	go s.processMigration(ctx, ff, from, to, deleteFiles)

	return total, nil
}

// isMigrating retrieves the migrationIsMigrating key from the
// cache and returns true if the app is already migrating
// files.
func (s *Storage) isMigrating(ctx context.Context) bool {
	_, err := s.cache.Get(ctx, migrationIsMigrating, nil)
	return err == nil
}

// getMigration returns the current migration information in
// the background.
func (s *Storage) getMigration() (*MigrationInfo, error) {
	const op = "Storage.GetMigration"
	mi := &MigrationInfo{}
	_, err := s.cache.Get(context.Background(), migrationKey, mi)
	if err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: "Error getting migration", Operation: op, Err: err}
	}
	return mi, nil
}

// processMigration ranges over the given files and adds a
// migration to the migrateTrackChan.
func (s *Storage) processMigration(ctx context.Context, files domain.Files, from, to domain.StorageChange, deleteFiles bool) {
	mi := &MigrationInfo{
		Total:      len(files),
		MigratedAt: time.Now(),
		mtx:        &sync.Mutex{},
	}

	s.cache.Set(ctx, migrationIsMigrating, true, cache.Options{Expiration: cache.RememberForever})
	s.cache.Set(ctx, migrationKey, mi, cache.Options{Expiration: cache.RememberForever})

	var wg sync.WaitGroup

	// migrateTrackChan is the channel used for sending and
	// processing migrations.
	var c = make(chan migration, migrateConcurrentAllowance)

	for _, file := range files {
		wg.Add(1)
		c <- migration{
			file: file,
			from: from,
			to:   to,
			wg:   &wg,
		}
		go s.migrateBackground(ctx, c, deleteFiles, mi)
	}

	wg.Wait()

	logger.Info(fmt.Sprintf("Storage: %d files migrated successfully", mi.Succeeded))
	logger.Info(fmt.Sprintf("Storage: %d files encountered an error during migration", mi.Failed))

	s.cache.Delete(context.Background(), migrationIsMigrating)
	s.cache.Delete(context.Background(), migrationKey)
}

// migrateBackground processes the migration by finding the
// original bytes, uploading to the new destination
// and deleting the original file.
func (s *Storage) migrateBackground(ctx context.Context, channel chan migration, deleteFiles bool, info *MigrationInfo) {
	m := <-channel

	info.mtx.Lock()

	defer func() {
		s.cache.Set(ctx, migrationKey, info, cache.Options{Expiration: cache.RememberForever})
		info.mtx.Unlock()
		m.wg.Done()
	}()

	buf, _, err := s.Find(m.file.URL)
	if err != nil {
		info.fail(m.file, err)
		return
	}

	u := domain.Upload{
		UUID:       m.file.UUID,
		Path:       m.file.URL,
		Size:       m.file.FileSize,
		Contents:   bytes.NewReader(buf),
		Private:    bool(m.file.Private),
		SourceType: m.file.SourceType,
	}

	file, err := s.upload(m.to.Provider, m.to.Bucket, u, false)
	if err != nil {
		info.fail(m.file, err)
		return
	}

	if deleteFiles {
		err = s.deleteFile(false, m.file.ID)
		if err != nil {
			info.fail(m.file, err)
			return
		}
	}

	file.ID = m.file.ID

	updated, err := s.filesRepo.Update(file)
	if err != nil {
		info.fail(m.file, err)
		return
	}

	info.succeed(updated)
}
