// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	"github.com/verbiscms/verbis/api/mocks/services/storage/mocks"
	repo "github.com/verbiscms/verbis/api/mocks/store/files"
	"io/ioutil"
	"strings"
	"sync"
)

func (t *StorageTestSuite) TestMigrationInfo_Fail() {
	mi := MigrationInfo{
		Failed: 0,
		Total:  100,
		Errors: nil,
		mtx:    &sync.Mutex{},
	}
	mi.fail(fileRemote, fmt.Errorf("error"))
	t.Equal(1, mi.FilesProcessed)
	t.Equal(1, mi.Failed)
	t.Equal(mi.Errors[0].File, fileRemote)
}

func (t *StorageTestSuite) TestMigrationInfo_Succeed() {
	mtx := &sync.Mutex{}

	tt := map[string]struct {
		input MigrationInfo
		want  MigrationInfo
	}{
		"Simple": {
			MigrationInfo{Total: 100, Succeeded: 0, FilesProcessed: 0, mtx: mtx},
			MigrationInfo{Total: 100, Succeeded: 1, FilesProcessed: 1, Progress: 1, mtx: mtx},
		},
		"Half": {
			MigrationInfo{Total: 100, Succeeded: 50, FilesProcessed: 50, mtx: mtx},
			MigrationInfo{Total: 100, Succeeded: 51, FilesProcessed: 51, Progress: 51, mtx: mtx},
		},
		"100": {
			MigrationInfo{Total: 100, Succeeded: 99, FilesProcessed: 99, mtx: mtx},
			MigrationInfo{Total: 100, Succeeded: 100, FilesProcessed: 100, Progress: 100, mtx: mtx},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			test.input.succeed(domain.File{})
			t.Equal(test.want, test.input)
		})
	}
}

func (t *StorageTestSuite) TestStorage_Migrate() {
	mockCacheSuccess := func(m *cache.Store) {
		m.On("Get", mock.Anything, migrationIsMigratingKey, mock.Anything).Return(fmt.Errorf("error"))
		m.On("Get", mock.Anything, migrationKey, mock.Anything).Return(fmt.Errorf("error"))
		m.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Times(100)
		m.On("Delete", mock.Anything, migrationIsMigratingKey)
		m.On("Delete", mock.Anything, migrationKey)
	}

	tt := map[string]struct {
		mock  func(m *mocks.Service, r *repo.Repository)
		cache func(m *cache.Store)
		want  interface{}
	}{
		"Success": {
			func(m *mocks.Service, r *repo.Repository) {
				mockValidateSuccess(m, r)
				m.On("Config").Return(domain.StorageConfig{Provider: domain.StorageLocal, Bucket: ""})
				r.On("List", mock.Anything).Return(filesSlice, 2, nil)
				r.On("FindByURL", filesSlice[0].URL).Return(domain.File{}, fmt.Errorf("error"))
				r.On("FindByURL", filesSlice[1].URL).Return(domain.File{}, fmt.Errorf("error"))
			},
			mockCacheSuccess,
			2,
		},
		"Already Migrating": {
			nil,
			func(m *cache.Store) {
				m.On("Get", mock.Anything, migrationIsMigratingKey, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						arg := args.Get(2).(*bool)
						*arg = true
					})
			},
			"Error migration is already in progress",
		},
		"Validation Failed": {
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageConfig{Provider: domain.StorageAWS, Bucket: ""})
			},
			mockCacheSuccess,
			"Validation failed",
		},
		"Repo Error": {
			func(m *mocks.Service, r *repo.Repository) {
				mockValidateSuccess(m, r)
				m.On("Config").Return(domain.StorageConfig{Provider: domain.StorageLocal, Bucket: ""})
				r.On("List", mock.Anything).Return(nil, 0, &errors.Error{Message: "error"})
			},
			mockCacheSuccess,
			"error",
		},
		"Zero Length": {
			func(m *mocks.Service, r *repo.Repository) {
				mockValidateSuccess(m, r)
				m.On("Config").Return(domain.StorageConfig{Provider: domain.StorageLocal, Bucket: ""})
				r.On("List", mock.Anything).Return(nil, 0, nil)
			},
			mockCacheSuccess,
			"Error, no files found to migrate",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			c := &cache.Store{}
			test.cache(c)
			s.cache = c

			s.env = &environment.Env{AWSAccessKey: "key", AWSSecret: "secret"}
			total, err := s.Migrate(context.Background(), true, true)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, total)
		})
	}
}

func (t *StorageTestSuite) TestStorage_GetMigration() {
	tt := map[string]struct {
		mock func(c *cache.Store)
		want interface{}
		nil  bool
	}{
		"Success": {
			func(c *cache.Store) {
				c.On("Get", mock.Anything, migrationKey, &MigrationInfo{}).Return(nil).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*MigrationInfo)
					arg.Total = 10
				})
			},
			MigrationInfo{Total: 10},
			false,
		},
		"Find Error": {
			func(c *cache.Store) {
				c.On("Get", mock.Anything, migrationKey, &MigrationInfo{}).Return(fmt.Errorf("error"))
			},
			nil,
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := &cache.Store{}
			test.mock(c)
			s := t.Setup(nil)
			s.cache = c
			got := s.getMigration(context.Background())
			if test.nil {
				t.Nil(got)
				return
			}
			t.Equal(test.want, *got)
		})
	}
}

func (t *StorageTestSuite) TestStorage_MigrateBackground() {
	tt := map[string]struct {
		mock func(m *mocks.Service, r *repo.Repository)
		want MigrationInfo
	}{
		"Find Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", fileRemote.URL).Return(domain.File{}, fmt.Errorf("error"))
			},
			MigrationInfo{Failed: 1, Succeeded: 0},
		},
		"Upload Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)

				c := &mocks.StowContainer{}
				m.On("BucketByFile", domain.File{}).Return(c, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(ioutil.NopCloser(strings.NewReader("test")), nil)
				c.On("Item", mock.Anything).Return(item, nil)

				m.On("Bucket", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("Error"))
			},
			MigrationInfo{Failed: 1, Succeeded: 0},
		},
		"Delete Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(ioutil.NopCloser(strings.NewReader("test")), nil)
				item.On("ID").Return("item")

				c := &mocks.StowContainer{}
				c.On("Item", mock.Anything).Return(item, nil)
				c.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(item, nil)
				c.On("ID").Return("bucket")

				m.On("BucketByFile", domain.File{}).Return(c, nil)
				m.On("Bucket", mock.Anything, mock.Anything).Return(c, nil)
				r.On("Create", mock.Anything).Return(domain.File{}, nil)

				r.On("Find", mock.Anything).Return(domain.File{}, fmt.Errorf("error"))
			},
			MigrationInfo{Failed: 1, Succeeded: 0},
		},
		"Repo Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(ioutil.NopCloser(strings.NewReader("test")), nil)
				item.On("ID").Return("item")

				c := &mocks.StowContainer{}
				c.On("Item", mock.Anything).Return(item, nil)
				c.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(item, nil)
				c.On("ID").Return("bucket")
				c.On("RemoveItem", mock.Anything).Return(nil)

				m.On("BucketByFile", domain.File{}).Return(c, nil).Times(2)
				m.On("Bucket", mock.Anything, mock.Anything).Return(c, nil)

				r.On("Find", mock.Anything).Return(domain.File{}, nil)
				r.On("Update", mock.Anything).Return(domain.File{}, fmt.Errorf("error"))
			},
			MigrationInfo{Failed: 1, Succeeded: 0},
		},
		"Success": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(ioutil.NopCloser(strings.NewReader("test")), nil)
				item.On("ID").Return("item")

				c := &mocks.StowContainer{}
				c.On("Item", mock.Anything).Return(item, nil)
				c.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(item, nil)
				c.On("ID").Return("bucket")
				c.On("RemoveItem", mock.Anything).Return(nil)

				m.On("BucketByFile", domain.File{}).Return(c, nil).Times(2)
				m.On("Bucket", mock.Anything, mock.Anything).Return(c, nil)

				r.On("Find", mock.Anything).Return(domain.File{}, nil)
				r.On("Update", mock.Anything).Return(domain.File{}, nil)
			},
			MigrationInfo{Failed: 0, Succeeded: 1},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)

			wg := sync.WaitGroup{}
			wg.Add(1)
			c := make(chan migration, 1)
			c <- migration{
				file: fileRemote,
				wg:   &wg,
			}

			mi := &MigrationInfo{
				mtx:   &sync.Mutex{},
				Total: 2,
			}
			s.migrateBackground(context.Background(), c, true, mi)

			t.Equal(test.want.Failed, mi.Failed)
			t.Equal(test.want.Succeeded, mi.Succeeded)
		})
	}
}
