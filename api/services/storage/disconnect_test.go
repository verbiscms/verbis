// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/mocks/services/storage/mocks"
	repo "github.com/verbiscms/verbis/api/mocks/store/files"
	options "github.com/verbiscms/verbis/api/mocks/store/options"
)

func (t *StorageTestSuite) TestStorage_Disconnect() {
	tt := map[string]struct {
		mock func(m *mocks.Service, r *repo.Repository, o *options.Repository)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Service, r *repo.Repository, o *options.Repository) {
				m.On("Config").
					Return(domain.StorageConfig{Provider: domain.StorageAWS})
				o.On("Insert", domain.OptionsDBMap{
					"storage_provider": domain.StorageLocal,
					"storage_bucket":   "",
				}).Return(nil)
			},
			nil,
		},
		"Already Disconnected": {
			func(m *mocks.Service, r *repo.Repository, o *options.Repository) {
				m.On("Config").
					Return(domain.StorageConfig{Provider: domain.StorageLocal})
			},
			ErrAlreadyDisconnected.Error(),
		},
		"Repo Error": {
			func(m *mocks.Service, r *repo.Repository, o *options.Repository) {
				m.On("Config").
					Return(domain.StorageConfig{Provider: domain.StorageAWS})
				o.On("Insert", domain.OptionsDBMap{
					"storage_provider": domain.StorageLocal,
					"storage_bucket":   "",
				}).Return(fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.SetupOptions(test.mock)
			got := s.Disconnect()
			if got != nil {
				t.Contains(got.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
