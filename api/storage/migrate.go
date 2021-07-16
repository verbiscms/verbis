// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/common/params"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
)

type MigrationInfo struct {
	Total     int64
	Processed int64
	Failed    int
	Succeeded int
	Errors    []FailedMigrationFile
}

type FailedMigrationFile struct {
	Error error
	File  domain.File
}

func (m *MigrationInfo) fail(file domain.File, err error) {
	m.Failed += 1
	m.Errors = append(m.Errors, FailedMigrationFile{
		Error: err,
		File:  file,
	})
}

func (m *MigrationInfo) succeed() {
	m.Processed = m.Total / int64(m.processed()) * int64(100)
	m.Succeeded += 1
}

func (m *MigrationInfo) processed() int {
	return m.Succeeded + m.Failed
}

func (s *Storage) Migrate(from, to domain.StorageChange) (int, error) {
	const op = "Storage.Migrate"
	if s.isMigrating {
		return 0, &errors.Error{Code: errors.INVALID, Message: "Error, migration is already in progress", Operation: op, Err: nil}
	}

	files, total, err := s.filesRepo.List(params.Params{LimitAll: false})
	if err != nil {
		return 0, err
	}

	s.isMigrating = true
	s.migration = MigrationInfo{
		Total: int64(total),
	}

	go s.migrateBackground(files, from, to)

	return total, nil
}

func (s *Storage) migrateBackground(files domain.Files, from, to domain.StorageChange) {
	for _, f := range files {
		if from.Provider == f.Provider {
			continue
		}

		buf, _, err := s.Find(f.Url)
		if err != nil {
			s.migration.fail(f, err)
			continue
		}

		u := domain.Upload{
			UUID:       uuid.New(),
			Path:       f.Url,
			Size:       0,
			Contents:   bytes.NewReader(buf),
			Private:    bool(f.Private),
			SourceType: f.SourceType,
		}

		_, err = s.upload(from.Provider, from.Bucket, u)

		if err != nil {
			s.migration.fail(f, err)
			continue
		}

		err = s.filesRepo.Update(f.Id, to)
		if err != nil {
			s.migration.fail(f, err)
			continue
		}

		s.migration.succeed()
	}

	s.migration = MigrationInfo{}
}
