// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/mocks/services/storage/mocks"
	repo "github.com/ainsleyclark/verbis/api/mocks/store/files"
	"github.com/graymeta/stow"
	"github.com/stretchr/testify/mock"
)

var mockValidateSuccess = func(m *mocks.Service, r *repo.Repository) {
	m.On("Config").Return(domain.StorageLocal, TestBucket, nil)

	c1 := &mocks.StowContainer{}
	c1.On("ID").Return(TestBucket)
	c1.On("Name").Return("name-1")

	loc := &mocks.StowLocation{}
	loc.On("Containers", mock.Anything, mock.Anything, pageSize).Return([]stow.Container{c1}, "", nil)

	m.On("Provider", domain.StorageLocal).Return(loc, nil)
}

func (t *StorageTestSuite) TestStorage_Validate() {
	tt := map[string]struct {
		input domain.StorageChange
		mock  func(m *mocks.Service, r *repo.Repository)
		want  error
	}{
		"Success": {
			domain.StorageChange{Provider: domain.StorageLocal, Bucket: TestBucket},
			mockValidateSuccess,
			nil,
		},
		"Empty Bucket": {
			domain.StorageChange{Provider: domain.StorageAWS, Bucket: ""},
			nil,
			fmt.Errorf("bucket cannot be empty"),
		},
		"Info Error": {
			domain.StorageChange{Provider: domain.StorageAWS, Bucket: TestBucket},
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageLocal, "", fmt.Errorf("error"))
			},
			fmt.Errorf("error"),
		},
		"Connect Error": {
			domain.StorageChange{Provider: domain.StorageAWS, Bucket: TestBucket},
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageAWS, TestBucket, nil)
			},
			fmt.Errorf("Configuration not set for: Aws"),
		},
		"List Buckets Error": {
			domain.StorageChange{Provider: domain.StorageLocal, Bucket: TestBucket},
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageLocal, TestBucket, nil)
				m.On("Provider", domain.StorageLocal).Return(nil, fmt.Errorf("error"))
			},
			fmt.Errorf("error"),
		},
		"No Bucket Found": {
			domain.StorageChange{Provider: domain.StorageAWS, Bucket: TestBucket},
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageAWS, TestBucket, nil)

				c1 := &mocks.StowContainer{}
				c1.On("ID").Return("id-1")
				c1.On("Name").Return("name-1")

				loc := &mocks.StowLocation{}
				loc.On("Containers", mock.Anything, mock.Anything, pageSize).Return([]stow.Container{c1}, "", nil)

				m.On("Provider", domain.StorageAWS).Return(loc, nil)
			},
			fmt.Errorf("invalid storage bucket: %s", TestBucket),
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			if name == "No Bucket Found" {
				s.env = &environment.Env{AWSAccessKey: "key", AWSSecret: "secret"}
			}
			got := s.validate(test.input)
			t.Equal(test.want, got)
		})
	}
}
