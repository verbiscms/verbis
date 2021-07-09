// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/mocks/storage/mocks"
	options "github.com/ainsleyclark/verbis/api/mocks/store/options"
	"github.com/graymeta/stow"
)

// bucket is the default bucket for testing.
const bucket = "verbis-bucket"

func (t *StorageTestSuite) SetupContainer(local bool, mock func(s *mocks.StowLocation, o *options.Repository)) *Storage {
	m := &mocks.StowLocation{}
	o := &options.Repository{}

	if mock != nil {
		mock(m, o)
	}

	provider := domain.StorageAWS
	if local {
		provider = domain.StorageLocal
	}

	s := &Storage{
		optionsRepo:  o,
		stowLocation: m,
		options: &domain.Options{
			StorageProvider: provider,
		},
	}

	return s
}

func (t *StorageTestSuite) TestContainer_SetBucket() {
	tt := map[string]struct {
		local bool
		mock  func(s *mocks.StowLocation, o *options.Repository)
		want  interface{}
	}{
		"Local": {
			true,
			nil,
			"Error setting bucket",
		},
		"Remote": {
			false,
			func(s *mocks.StowLocation, o *options.Repository) {
				s.On("Container", bucket).Return(&mocks.StowContainer{}, nil)
				o.On("Update", "storage_bucket", bucket).Return(nil)
			},
			nil,
		},
		"Stow Error": {
			false,
			func(s *mocks.StowLocation, o *options.Repository) {
				s.On("Container", bucket).Return(&mocks.StowContainer{}, fmt.Errorf("error"))
			},
			"Error setting bucket",
		},
		"Repo Error": {
			false,
			func(s *mocks.StowLocation, o *options.Repository) {
				s.On("Container", bucket).Return(&mocks.StowContainer{}, nil)
				o.On("Update", "storage_bucket", bucket).Return(fmt.Errorf("error"))
			},
			"Error updating options table with new bucket",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.SetupContainer(test.local, test.mock)
			orig := s.stowContainer
			err := s.SetBucket(bucket)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.NotEqual(s.stowContainer, orig)
			t.Equal(test.want, err)
		})
	}
}

func (t *StorageTestSuite) TestContainer_ListBuckets() {
	tt := map[string]struct {
		local bool
		mock  func(s *mocks.StowLocation, o *options.Repository)
		want  interface{}
	}{
		"Local": {
			true,
			nil,
			"Error listing buckets",
		},
		"Remote": {
			false,
			func(s *mocks.StowLocation, o *options.Repository) {
				c1 := &mocks.StowContainer{}
				c1.On("ID").Return("id-1")
				c1.On("Name").Return("name-1")
				c2 := &mocks.StowContainer{}
				c2.On("ID").Return("id-2")
				c2.On("Name").Return("name-2")
				s.On("Containers", "", "", pageSize).Return([]stow.Container{c1, c2}, "", nil)
			},
			domain.Buckets{
				domain.Bucket{Id: "id-1", Name: "name-1"},
				domain.Bucket{Id: "id-2", Name: "name-2"},
			},
		},
		"Error": {
			false,
			func(s *mocks.StowLocation, o *options.Repository) {
				s.On("Containers", "", "", pageSize).Return(nil, "", fmt.Errorf("error"))
			},
			"Error obtaining buckets",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.SetupContainer(test.local, test.mock)
			got, err := s.ListBuckets()
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
		local bool
		mock  func(s *mocks.StowLocation, o *options.Repository)
		want  interface{}
	}{
		"Success": {
			false,
			func(s *mocks.StowLocation, o *options.Repository) {
				s.On("CreateContainer", bucket).Return(nil, nil)
			},
			nil,
		},
		"Local": {
			true,
			nil,
			"Error creating bucket",
		},
		"Error": {
			false,
			func(s *mocks.StowLocation, o *options.Repository) {
				s.On("CreateContainer", bucket).Return(nil, fmt.Errorf("error"))
			},
			"Error creating bucket",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.SetupContainer(test.local, test.mock)
			err := s.CreateBucket(bucket)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}

func (t *StorageTestSuite) TestContainer_Delete() {
	tt := map[string]struct {
		local bool
		mock  func(s *mocks.StowLocation, o *options.Repository)
		want  interface{}
	}{
		"Success": {
			false,
			func(s *mocks.StowLocation, o *options.Repository) {
				s.On("RemoveContainer", bucket).Return(nil)
			},
			nil,
		},
		"Local": {
			true,
			nil,
			"Error deleting bucket",
		},
		"Error": {
			false,
			func(s *mocks.StowLocation, o *options.Repository) {
				s.On("RemoveContainer", bucket).Return(fmt.Errorf("error"))
			},
			"Error deleting bucket with the name",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.SetupContainer(test.local, test.mock)
			err := s.DeleteBucket(bucket)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}
