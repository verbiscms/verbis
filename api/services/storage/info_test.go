// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/graymeta/stow"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/mocks/services/storage/mocks"
	repo "github.com/verbiscms/verbis/api/mocks/store/files"
	internal2 "github.com/verbiscms/verbis/api/services/storage/internal"
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
		mock func(m *mocks.Service, r *repo.Repository)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageAWS, TestBucket, nil)
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
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			orig := internal2.Providers
			defer func() { internal2.Providers = orig }()
			internal2.Providers = internal2.ProviderMap{"test": &mockProviderErr{}}

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
