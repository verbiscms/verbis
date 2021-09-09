// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/mocks/services/storage/mocks"
	repo "github.com/verbiscms/verbis/api/mocks/store/files"
	options "github.com/verbiscms/verbis/api/mocks/store/options"
)

func (t *StorageTestSuite) TestStorage_Save() {
	tt := map[string]struct {
		input domain.StorageConfig
		mock  func(m *mocks.Service, r *repo.Repository, o *options.Repository)
		want  interface{}
	}{
		"Success": {
			domain.StorageConfig{Provider: domain.StorageLocal, Bucket: TestBucket, LocalBackup: false},
			func(m *mocks.Service, r *repo.Repository, o *options.Repository) {
				mockValidateSuccess(m, r)
				o.On("Insert", domain.OptionsDBMap{
					"storage_provider":     domain.StorageLocal,
					"storage_bucket":       "",
					"storage_local_backup": false,
				}).Return(nil)
			},
			nil,
		},
		"Validation Failed": {
			domain.StorageConfig{Provider: domain.StorageAWS},
			nil,
			"Validation failed",
		},
		"Error": {
			domain.StorageConfig{Provider: domain.StorageLocal, Bucket: TestBucket, LocalBackup: false},
			func(m *mocks.Service, r *repo.Repository, o *options.Repository) {
				mockValidateSuccess(m, r)
				o.On("Insert", domain.OptionsDBMap{
					"storage_provider":     domain.StorageLocal,
					"storage_bucket":       "",
					"storage_local_backup": false,
				}).Return(&errors.Error{Message: "error"})
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.SetupOptions(test.mock)
			got := s.Save(test.input)
			if got != nil {
				t.Contains(errors.Message(got), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
