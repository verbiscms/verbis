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
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/url"
	"strings"
)

func (t *StorageTestSuite) TestBucket_Find() {
	tt := map[string]struct {
		mock func(m *mocks.Service, r *repo.Repository)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)

				c := &mocks.StowContainer{}
				m.On("BucketByFile", domain.File{}).Return(c, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(ioutil.NopCloser(strings.NewReader("test")), nil)
				c.On("Item", mock.Anything).Return(item, nil)
			},
			"test",
		},
		"Repo Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, &errors.Error{Message: "error"})
			},
			"error",
		},
		"BucketByFile Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)
				m.On("BucketByFile", mock.Anything).Return(nil, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Item Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)
				c := &mocks.StowContainer{}
				m.On("BucketByFile", domain.File{}).Return(c, nil)
				c.On("Item", mock.Anything).Return(nil, fmt.Errorf("error"))
			},
			"Error obtaining file with the ID",
		},
		"Open Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)

				c := &mocks.StowContainer{}
				m.On("BucketByFile", domain.File{}).Return(c, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(nil, fmt.Errorf("error"))
				c.On("Item", mock.Anything).Return(item, nil)
			},
			"Error opening file",
		},
		"Read Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("FindByURL", mock.Anything).Return(domain.File{}, nil)

				c := &mocks.StowContainer{}
				m.On("BucketByFile", domain.File{}).Return(c, nil)

				item := &mocks.StowItem{}
				item.On("Open").Return(&mockIOReaderReadError{}, nil)
				c.On("Item", mock.Anything).Return(item, nil)
			},
			"Error reading file",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
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
		mock  func(m *mocks.Service, r *repo.Repository)
		local bool
		want  interface{}
	}{
		"Local": {
			upload,
			func(m *mocks.Service, r *repo.Repository) {
				item := &mocks.StowItem{}
				item.On("ID").Return("test.txt")
				item.On("URL").Return(&url.URL{Path: "/uploads/2020/01/test.txt"})

				cont := &mocks.StowContainer{}
				cont.On("ID").Return("bucket")
				cont.On("Put", "/uploads/2020/01/"+key+".txt", upload.Contents, upload.Size, mock.Anything).Return(item, nil)

				m.On("Config").Return(domain.StorageLocal, "", nil)
				m.On("Bucket", domain.StorageLocal, "").Return(cont, nil)
				r.On("Create", fileLocal).Return(fileLocal, nil)
			},
			true,
			fileLocal,
		},
		"Remote": {
			upload,
			func(m *mocks.Service, r *repo.Repository) {
				item := &mocks.StowItem{}
				item.On("ID").Return("uploads/2020/01/test.txt")

				cont := &mocks.StowContainer{}
				cont.On("ID").Return("bucket")
				cont.On("Put", mock.Anything, upload.Contents, upload.Size, mock.Anything).Return(item, nil)

				m.On("Config").Return(domain.StorageAWS, "", nil)
				m.On("Bucket", domain.StorageAWS, "").Return(cont, nil)
				r.On("Create", fileRemote).Return(fileRemote, nil)
			},
			false,
			fileRemote,
		},
		"Config Error": {
			upload,
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageAWS, "", fmt.Errorf("error"))
			},
			false,
			"error",
		},
		"Bucket Error": {
			upload,
			func(m *mocks.Service, r *repo.Repository) {
				m.On("Config").Return(domain.StorageAWS, "", nil)
				m.On("Bucket", domain.StorageAWS, "").Return(nil, fmt.Errorf("error"))
			},
			false,
			"error",
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
			upload,
			func(m *mocks.Service, r *repo.Repository) {
				cont := &mocks.StowContainer{}
				cont.On("Put", mock.Anything, upload.Contents, upload.Size, mock.Anything).Return(&mocks.StowItem{}, fmt.Errorf("error"))
				m.On("Config").Return(domain.StorageAWS, "", nil)
				m.On("Bucket", domain.StorageAWS, "").Return(cont, nil)
			},
			true,
			"Error uploading file to storage provider",
		},
		"Mime Error": {
			domain.Upload{
				UUID:       uuid.New(),
				Path:       "/uploads/2020/01/test.txt",
				Size:       100,
				Contents:   &mockIOReaderReadError{},
				Private:    false,
				SourceType: domain.MediaSourceType,
			},
			func(m *mocks.Service, r *repo.Repository) {
				cont := &mocks.StowContainer{}
				cont.On("Put", mock.Anything, mock.Anything, upload.Size, mock.Anything).Return(&mocks.StowItem{}, nil)
				m.On("Config").Return(domain.StorageAWS, "", nil)
				m.On("Bucket", domain.StorageAWS, "").Return(cont, nil)
			},
			true,
			"Error obtaining mime type",
		},
		"Repo Error": {
			upload,
			func(m *mocks.Service, r *repo.Repository) {
				item := &mocks.StowItem{}
				item.On("ID").Return("test.txt")
				item.On("URL").Return(&url.URL{Path: "/uploads/2020/01/test.txt"})

				cont := &mocks.StowContainer{}
				cont.On("ID").Return("bucket")
				cont.On("Put", "/uploads/2020/01/"+key+".txt", upload.Contents, upload.Size, mock.Anything).Return(item, nil)

				m.On("Config").Return(domain.StorageLocal, "", nil)
				m.On("Bucket", domain.StorageLocal, "").Return(cont, nil)
				r.On("Create", fileLocal).Return(domain.File{}, fmt.Errorf("error"))
			},
			false,
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
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
		mock func(m *mocks.Service, r *repo.Repository)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("Find", mock.Anything).Return(domain.File{}, nil)
				c := &mocks.StowContainer{}
				m.On("BucketByFile", domain.File{}).Return(c, nil)
				c.On("RemoveItem", mock.Anything).Return(nil)
				r.On("Delete", mock.Anything).Return(nil)
			},
			nil,
		},
		"Find Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("Find", mock.Anything).Return(domain.File{}, &errors.Error{Message: "error"})
			},
			"error",
		},
		"BucketByFile Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("Find", mock.Anything).Return(domain.File{}, nil)
				m.On("BucketByFile", mock.Anything).Return(nil, &errors.Error{Message: "error"})
			},
			"error",
		},
		"Storage Remove Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("Find", mock.Anything).Return(domain.File{}, nil)
				c := &mocks.StowContainer{}
				m.On("BucketByFile", domain.File{}).Return(c, nil)
				c.On("RemoveItem", mock.Anything).Return(fmt.Errorf("error"))
			},
			"Error deleting file from storage",
		},
		"Repo Remove Error": {
			func(m *mocks.Service, r *repo.Repository) {
				r.On("Find", mock.Anything).Return(domain.File{}, nil)
				c := &mocks.StowContainer{}
				m.On("BucketByFile", domain.File{}).Return(c, nil)
				c.On("RemoveItem", mock.Anything).Return(nil)
				r.On("Delete", mock.Anything).Return(&errors.Error{Message: "error"})
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.mock)
			err := s.Delete(1)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, err)
		})
	}
}
