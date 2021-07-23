// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/mocks/services/storage/mocks"
	repo "github.com/ainsleyclark/verbis/api/mocks/store/files"
	options "github.com/ainsleyclark/verbis/api/mocks/store/options"
)

func (t *StorageTestSuite) TestStorage_Save() {
	tt := map[string]struct {
		input domain.StorageChange
		mock  func(m *mocks.Service, r *repo.Repository, o *options.Repository)
		want  interface{}
	}{
		"Success": {
			domain.StorageChange{Provider: domain.StorageLocal, Bucket: TestBucket},
			func(m *mocks.Service, r *repo.Repository, o *options.Repository) {
				mockValidateSuccess(m, r)
				o.On("Update", "storage_provider", domain.StorageLocal).Return(nil)
				o.On("Update", "storage_bucket", "").Return(nil)
			},
			nil,
		},
		"Validation Failed": {
			domain.StorageChange{Provider: domain.StorageAWS},
			nil,
			"Validation failed",
		},
		"Provider Error": {
			domain.StorageChange{Provider: domain.StorageLocal, Bucket: TestBucket},
			func(m *mocks.Service, r *repo.Repository, o *options.Repository) {
				mockValidateSuccess(m, r)
				o.On("Update", "storage_provider", domain.StorageLocal).Return(&errors.Error{Message: "provider error"})
			},
			"provider error",
		},
		"Bucket Error": {
			domain.StorageChange{Provider: domain.StorageLocal, Bucket: TestBucket},
			func(m *mocks.Service, r *repo.Repository, o *options.Repository) {
				mockValidateSuccess(m, r)
				o.On("Update", "storage_provider", domain.StorageLocal).Return(nil)
				o.On("Update", "storage_bucket", "").Return(&errors.Error{Message: "bucket error"})
			},
			"bucket error",
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
