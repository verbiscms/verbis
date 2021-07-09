// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/mocks/storage/mocks"
	repo "github.com/ainsleyclark/verbis/api/mocks/store/files"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/url"
	"strings"
)

const (
	TestFileUrl = "/file.txt"
)

func (t *StorageTestSuite) TestBucket_Find() {
	tt := map[string]struct {
		mock func(s *mocks.Service, r *repo.Repository)
		want interface{}
	}{
		"Success": {
			func(s *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)

				c := &mocks.StowContainer{}
				s.On("Bucket", domain.File{}).Return(c, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(ioutil.NopCloser(strings.NewReader("test")), nil)
				c.On("Item", mock.Anything).Return(item, nil)
			},
			"test",
		},
		"Repo Error": {
			func(s *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Bucket Error": {
			func(s *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)
				s.On("Bucket", mock.Anything).Return(nil, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Item Error": {
			func(s *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)
				c := &mocks.StowContainer{}
				s.On("Bucket", domain.File{}).Return(c, nil)
				c.On("Item", mock.Anything).Return(nil, fmt.Errorf("error"))
			},
			"Error obtaining file with the ID",
		},
		"Open Error": {
			func(s *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)

				c := &mocks.StowContainer{}
				s.On("Bucket", domain.File{}).Return(c, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(nil, fmt.Errorf("error"))
				c.On("Item", mock.Anything).Return(item, nil)
			},
			"Error opening file",
		},
		"Read Error": {
			func(s *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)

				c := &mocks.StowContainer{}
				s.On("Bucket", domain.File{}).Return(c, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(&MockIOReaderError{}, nil)
				c.On("Item", mock.Anything).Return(item, nil)
			},
			"Error reading file",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := &mocks.Service{}
			r := &repo.Repository{}

			test.mock(m, r)

			s := Storage{
				filesRepo: r,
				service:   m,
			}

			got, _, err := s.Find(TestFileUrl)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}

			t.Equal(test.want, string(got))
		})
	}
}

func (t *StorageTestSuite) TestBucket_Upload() {
	tt := map[string]struct {
		input domain.Upload
		mock  func(s *mocks.Service, c *mocks.StowContainer, r *repo.Repository)
		local bool
		want  interface{}
	}{
		"Local": {
			u,
			func(s *mocks.Service, c *mocks.StowContainer, r *repo.Repository) {
				item := &mocks.StowItem{}
				item.On("ID").Return("test.txt")
				item.On("URL").Return(&url.URL{Path: "/uploads/2020/01/test.txt"})
				c.On("ID").Return("bucket")
				c.On("Put", "/uploads/2020/01/"+key+".txt", u.Contents, u.Size, mock.Anything).Return(item, nil)
				r.On("Create", fileLocal).Return(fileLocal, nil)
			},
			true,
			fileLocal,
		},
		"Remote": {
			u,
			func(s *mocks.Service, c *mocks.StowContainer, r *repo.Repository) {
				item := &mocks.StowItem{}
				item.On("ID").Return("uploads/2020/01/test.txt")
				c.On("ID").Return("bucket")
				c.On("Put", mock.Anything, u.Contents, u.Size, mock.Anything).Return(item, nil)
				r.On("Create", fileRemote).Return(fileRemote, nil)
			},
			false,
			fileRemote,
		},
		"Validate Error": {
			domain.Upload{
				Path: "",
			},
			nil,
			true,
			"Validation failed",
		},
		"Put Error": {
			u,
			func(s *mocks.Service, c *mocks.StowContainer, r *repo.Repository) {
				c.On("Put", mock.Anything, u.Contents, u.Size, mock.Anything).Return(&mocks.StowItem{}, fmt.Errorf("error"))
			},
			true,
			"Error uploading file to storage provider",
		},
		"Seek Error": {
			domain.Upload{
				Path:       "/uploads/2020/01/test.txt",
				Size:       100,
				Contents:   &MockIOSeekerError{},
				Private:    false,
				SourceType: domain.MediaSourceType,
			},
			func(s *mocks.Service, c *mocks.StowContainer, r *repo.Repository) {
				c.On("Put", mock.Anything, mock.Anything, u.Size, mock.Anything).Return(&mocks.StowItem{}, nil)
			},
			true,
			"Error seeking bytes",
		},
		"Mime Error": {
			domain.Upload{
				Path:       "/uploads/2020/01/test.txt",
				Size:       100,
				Contents:   &MockIOReaderSeekerError{},
				Private:    false,
				SourceType: domain.MediaSourceType,
			},
			func(s *mocks.Service, c *mocks.StowContainer, r *repo.Repository) {
				c.On("Put", mock.Anything, mock.Anything, u.Size, mock.Anything).Return(&mocks.StowItem{}, nil)
			},
			true,
			"Error obtaining mime type",
		},
		"Repo Error": {
			u,
			func(s *mocks.Service, c *mocks.StowContainer, r *repo.Repository) {
				item := &mocks.StowItem{}
				item.On("ID").Return("uploads/2020/01/test.txt")
				c.On("ID").Return("bucket")
				c.On("Put", mock.Anything, u.Contents, u.Size, mock.Anything).Return(item, nil)
				r.On("Create", fileRemote).Return(domain.File{}, &errors.Error{Message: "error"})
			},
			false,
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := &mocks.Service{}
			r := &repo.Repository{}
			c := &mocks.StowContainer{}

			if test.mock != nil {
				test.mock(m, c, r)
			}

			s := Storage{
				filesRepo:     r,
				service:       m,
				stowContainer: c,
				ProviderName:  domain.StorageAWS,
			}

			if test.local {
				s.ProviderName = domain.StorageLocal
			}

			got, err := s.Upload(test.input)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *StorageTestSuite) TestBucket_Exists() {
	tt := map[string]struct {
		input string
		mock  func(r *repo.Repository)
		want  interface{}
	}{
		"True": {
			"test",
			func(r *repo.Repository) {
				r.On("Exists", "test").Return(true)
			},
			true,
		},
		"False": {
			"test",
			func(r *repo.Repository) {
				r.On("Exists", "test").Return(false)
			},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			r := &repo.Repository{}

			test.mock(r)

			s := Storage{
				filesRepo: r,
			}

			got := s.Exists(test.input)
			t.Equal(test.want, got)
		})
	}
}

func (t *StorageTestSuite) TestBucket_Delete() {
	tt := map[string]struct {
		mock func(s *mocks.Service, r *repo.Repository)
		want interface{}
	}{
		"Success": {
			func(s *mocks.Service, r *repo.Repository) {
				r.On("Find", mock.Anything).Return(domain.File{}, nil)
				c := &mocks.StowContainer{}
				s.On("Bucket", domain.File{}).Return(c, nil)
				c.On("RemoveItem", mock.Anything).Return(nil)
				r.On("Delete", mock.Anything).Return(nil)
			},
			nil,
		},
		"Find Error": {
			func(s *mocks.Service, r *repo.Repository) {
				r.On("Find", mock.Anything).Return(domain.File{}, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Bucket Error": {
			func(s *mocks.Service, r *repo.Repository) {
				r.On("Find", mock.Anything).Return(domain.File{}, nil)
				s.On("Bucket", mock.Anything).Return(nil, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Storage Remove Error": {
			func(s *mocks.Service, r *repo.Repository) {
				r.On("Find", mock.Anything).Return(domain.File{}, nil)
				c := &mocks.StowContainer{}
				s.On("Bucket", domain.File{}).Return(c, nil)
				c.On("RemoveItem", mock.Anything).Return(fmt.Errorf("error"))
			},
			"Error deleting file from storage",
		},
		"Repo Remove Error": {
			func(s *mocks.Service, r *repo.Repository) {
				r.On("Find", mock.Anything).Return(domain.File{}, nil)
				c := &mocks.StowContainer{}
				s.On("Bucket", domain.File{}).Return(c, nil)
				c.On("RemoveItem", mock.Anything).Return(nil)
				r.On("Delete", mock.Anything).Return(&errors.Error{Message: "error"})
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := &mocks.Service{}
			r := &repo.Repository{}

			test.mock(m, r)

			s := Storage{
				filesRepo: r,
				service:   m,
			}

			err := s.Delete(1)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}

			t.Equal(test.want, err)
		})
	}
}
