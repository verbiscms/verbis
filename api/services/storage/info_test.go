// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/graymeta/stow"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	mockCache "github.com/verbiscms/verbis/api/mocks/cache"
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
	tt := map[string]struct {
		mock  func(m *mocks.Service, r *repo.Repository)
		cache func(c *mockCache.Cacher)
		want  interface{}
	}{
		"Success": {
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageAWS, TestBucket, nil)
			},
			func(c *mockCache.Cacher) {
				c.On("Get", mock.Anything, MigrationCacheKey).Return(MigrationInfo{}, nil)
			},
			Configuration{
				ActiveProvider: domain.StorageAWS,
				ActiveBucket:   TestBucket,
				Providers: domain.StorageProviders{
					"test": domain.StorageProviderInfo{},
				},
				IsMigrating:   false,
				MigrationInfo: MigrationInfo{},
			},
		},
		"Error": {
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageAWS, "", fmt.Errorf("error"))
			},
			func(c *mockCache.Cacher) {
				c.On("Get", mock.Anything, MigrationCacheKey).Return(MigrationInfo{}, nil)
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			orig := internal.Providers
			defer func() { internal.Providers = orig }()
			internal.Providers = internal.ProviderMap{"test": &mockProviderErr{}}

			c := &mockCache.Cacher{}
			test.cache(c)
			cache.SetDriver(c)

			s := t.Setup(test.mock)
			got, err := s.Info()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
