// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE downloadFile.

package storage

import (
	"fmt"
	"github.com/graymeta/stow"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/mocks/services/storage/mocks"
	repo "github.com/verbiscms/verbis/api/mocks/store/files"
)

func (t *StorageTestSuite) TestContainer_ListBuckets() {
	tt := map[string]struct {
		mock func(m *mocks.Service, r *repo.Repository)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Service, r *repo.Repository) {
				c1 := &mocks.StowContainer{}
				c1.On("ID").Return("id-1")
				c1.On("Name").Return("name-1")

				c2 := &mocks.StowContainer{}
				c2.On("ID").Return("id-2")
				c2.On("Name").Return("name-2")

				loc := &mocks.StowLocation{}
				loc.On("Containers", mock.Anything, mock.Anything, pageSize).Return([]stow.Container{c1, c2}, "", nil)

				m.On("Provider", domain.StorageLocal).Return(loc, nil)
			},
			domain.Buckets{
				domain.Bucket{ID: "id-1", Name: "name-1"},
				domain.Bucket{ID: "id-2", Name: "name-2"},
			},
		},
		"Service Error": {
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Provider", domain.StorageLocal).Return(nil, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Container Error": {
			func(m *mocks.Service, r *repo.Repository) {
				loc := &mocks.StowLocation{}
				loc.On("Containers", mock.Anything, mock.Anything, pageSize).Return(nil, "", fmt.Errorf("Error"))
				m.On("Provider", domain.StorageLocal).Return(loc, nil)
			},
			"Error obtaining buckets",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.ListBuckets(domain.StorageLocal)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *StorageTestSuite) TestContainer_CreateBucket() {
	tt := map[string]struct {
		mock func(m *mocks.Service, r *repo.Repository)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Service, r *repo.Repository) {
				loc := &mocks.StowLocation{}
				cont := &mocks.StowContainer{}
				cont.On("ID").Return("bucket-id")
				cont.On("Name").Return("bucket-name")
				loc.On("CreateContainer", TestBucket).Return(cont, nil)
				m.On("Provider", domain.StorageLocal).Return(loc, nil)
			},
			domain.Bucket{
				ID:   "bucket-id",
				Name: "bucket-name",
			},
		},
		"Service Error": {
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Provider", domain.StorageLocal).Return(nil, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Create Error": {
			func(m *mocks.Service, r *repo.Repository) {
				loc := &mocks.StowLocation{}
				loc.On("CreateContainer", TestBucket).Return(nil, fmt.Errorf("error"))
				m.On("Provider", domain.StorageLocal).Return(loc, nil)
			},
			"Error creating bucket",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			got, err := s.CreateBucket(domain.StorageLocal, TestBucket)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *StorageTestSuite) TestContainer_Delete() {
	tt := map[string]struct {
		mock func(m *mocks.Service, r *repo.Repository)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Service, r *repo.Repository) {
				loc := &mocks.StowLocation{}
				loc.On("RemoveContainer", TestBucket).Return(nil)
				m.On("Provider", domain.StorageLocal).Return(loc, nil)
			},
			nil,
		},
		"Service Error": {
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Provider", domain.StorageLocal).Return(nil, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Create Error": {
			func(m *mocks.Service, r *repo.Repository) {
				loc := &mocks.StowLocation{}
				loc.On("RemoveContainer", TestBucket).Return(fmt.Errorf("error"))
				m.On("Provider", domain.StorageLocal).Return(loc, nil)
			},
			"Error deleting bucket with the name",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.DeleteBucket(domain.StorageLocal, TestBucket)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}
