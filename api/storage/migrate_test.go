// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/common/params"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/mocks/storage/mocks"
	repo "github.com/ainsleyclark/verbis/api/mocks/store/files"
)

func (t *StorageTestSuite) TestMigrationInfo_Fail() {
	mi := MigrationInfo{
		Failed: 0,
		Errors: nil,
	}
	mi.fail(fileRemote, fmt.Errorf("error"))
	t.Equal(1, mi.FilesProcessed)
	t.Equal(1, mi.Failed)
	t.Equal(mi.Errors[0].File, fileRemote)
}

func (t *StorageTestSuite) TestMigrationInfo_Succeed() {
	tt := map[string]struct {
		input MigrationInfo
		want  MigrationInfo
	}{
		"Simple": {
			MigrationInfo{Total: 100, Succeeded: 0, FilesProcessed: 0},
			MigrationInfo{Total: 100, Succeeded: 1, FilesProcessed: 1, Progress: 1},
		},
		"Half": {
			MigrationInfo{Total: 100, Succeeded: 50, FilesProcessed: 50},
			MigrationInfo{Total: 100, Succeeded: 51, FilesProcessed: 51, Progress: 51},
		},
		"100": {
			MigrationInfo{Total: 100, Succeeded: 99, FilesProcessed: 99},
			MigrationInfo{Total: 100, Succeeded: 100, FilesProcessed: 100, Progress: 100},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			test.input.succeed()
			t.Equal(test.want, test.input)
		})
	}
}

func (t *StorageTestSuite) TestStorage_Migrate() {
	tt := map[string]struct {
		migrating bool
		from      domain.StorageChange
		to        domain.StorageChange
		mock      func(m *mocks.Service, r *repo.Repository)
		want      interface{}
	}{
		"Already Migrating": {
			true,
			domain.StorageChange{},
			domain.StorageChange{},
			nil,
			"Error migration is already in progress",
		},
		"Same Providers": {
			false,
			domain.StorageChange{Provider: domain.StorageAWS},
			domain.StorageChange{Provider: domain.StorageAWS},
			nil,
			"Error providers cannot be the same",
		},
		"Validation Failed": {
			false,
			domain.StorageChange{Provider: domain.StorageLocal},
			domain.StorageChange{Provider: domain.StorageAWS},
			nil,
			"Validation failed",
		},
		"Repo Error": {
			false,
			domain.StorageChange{Provider: domain.StorageAWS},
			domain.StorageChange{Provider: domain.StorageLocal, Bucket: TestBucket},
			func(m *mocks.Service, r *repo.Repository) {
				mockValidateSuccess(m, r)
				r.On("List", params.Params{LimitAll: false}).Return(nil, 0, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Test": {
			false,
			domain.StorageChange{Provider: domain.StorageAWS},
			domain.StorageChange{Provider: domain.StorageLocal, Bucket: TestBucket},
			func(m *mocks.Service, r *repo.Repository) {
				mockValidateSuccess(m, r)
				r.On("List", params.Params{LimitAll: false}).Return(nil, 0, &errors.Error{Message: "error"})
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			s.env = &environment.Env{AWSAccessKey: "key", AWSSecret: "secret"}
			if test.migrating {
				s.isMigrating = true
			}
			total, err := s.Migrate(test.from, test.to)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, total)
		})
	}
}

//func (t *StorageTestSuite) TestStorage_MigrateBackground() {
//	tt := map[string]struct {
//		file domain.File
//		from domain.StorageChange
//		to   domain.StorageChange
//		mock func(m *mocks.Service, r *repo.Repository)
//		want MigrationInfo
//	}{
//		"Same Provider": {
//			domain.File{Provider: domain.StorageAWS},
//			domain.StorageChange{Provider: domain.StorageLocal},
//			domain.StorageChange{Provider: domain.StorageAWS},
//			nil,
//			MigrationInfo{},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup(test.mock)
//			s.migrateBackground(test.file, test.from, test.to)
//			t.Equal(test.want, s.migration)
//		})
//	}
//}
