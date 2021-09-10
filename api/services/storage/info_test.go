// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"context"
	"fmt"
	"github.com/graymeta/stow"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	"github.com/verbiscms/verbis/api/mocks/services/storage/mocks"
	repo "github.com/verbiscms/verbis/api/mocks/store/files"
	"github.com/verbiscms/verbis/api/services/storage/internal"
)

type mockProviderErr struct{}

func (m *mockProviderErr) Dial(env *environment.Env) (stow.Location, error) {
	return nil, fmt.Errorf("error")
}

func (m *mockProviderErr) Info(env *environment.Env) domain.StorageProviderInfo {
	return domain.StorageProviderInfo{}
}

func (t *StorageTestSuite) TestStorage_Info() {
	mi := &MigrationInfo{Total: 10}
	tt := map[string]struct {
		mock  func(m *mocks.Service, r *repo.Repository)
		cache func(c *cache.Store)
		want  interface{}
	}{
		"Not Migrating": {
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageConfig{Provider: domain.StorageAWS, Bucket: TestBucket})
			},
			func(c *cache.Store) {
				c.On("Get", mock.Anything, migrationIsMigrating, mock.Anything).Return(fmt.Errorf("error"))
			},
			Configuration{
				Info: domain.StorageConfig{
					Provider: domain.StorageAWS,
					Bucket:   TestBucket,
				},
				Providers: domain.StorageProviders{
					"test": domain.StorageProviderInfo{},
				},
				IsMigrating:   false,
				MigrationInfo: nil,
			},
		},
		"Is Migrating": {
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageConfig{Provider: domain.StorageAWS, Bucket: TestBucket})
			},
			func(c *cache.Store) {
				c.On("Get", mock.Anything, migrationIsMigrating, mock.Anything).Return(nil)
				c.On("Get", mock.Anything, migrationKey, &MigrationInfo{}).Return(nil).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*MigrationInfo)
					arg.Total = 10
				})
			},
			Configuration{
				Info: domain.StorageConfig{
					Provider: domain.StorageAWS,
					Bucket:   TestBucket,
				},
				Providers: domain.StorageProviders{
					"test": domain.StorageProviderInfo{},
				},
				IsMigrating:   true,
				MigrationInfo: mi,
			},
		},
		"Migration Error": {
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageConfig{Provider: domain.StorageAWS, Bucket: TestBucket})
			},
			func(c *cache.Store) {
				c.On("Get", mock.Anything, migrationIsMigrating, mock.Anything).Return(nil)
				c.On("Get", mock.Anything, migrationKey, mock.Anything).Return(fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			orig := internal.Providers
			defer func() { internal.Providers = orig }()
			internal.Providers = internal.ProviderMap{"test": &mockProviderErr{}}

			s := t.Setup(test.mock)
			c := &cache.Store{}
			if test.cache != nil {
				test.cache(c)
			}
			s.cache = c

			got, err := s.Info(context.Background())
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
