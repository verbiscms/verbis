// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	storage "github.com/ainsleyclark/verbis/api/mocks/storage"
	repo "github.com/ainsleyclark/verbis/api/mocks/store/media"
	"github.com/stretchr/testify/mock"
)

func (t *MediaServiceTestSuite) TestService_Delete() {
	tt := map[string]struct {
		mock func(r *repo.Repository, s *storage.Bucket)
		want interface{}
	}{
		"Success": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("Find", MediaID).Return(testMedia, nil)
				r.On("Delete", MediaID).Return(nil)
				s.On("Delete", 1).Return(fmt.Errorf("error"))
				s.On("Find", TestFileURLWebP).Return(nil, domain.File{}, fmt.Errorf("error"))
			},
			nil,
		},
		"Find Error": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("Find", MediaID).Return(domain.Media{}, fmt.Errorf("error"))
			},
			"error",
		},
		"Delete Error": {
			func(r *repo.Repository, s *storage.Bucket) {
				r.On("Find", MediaID).Return(domain.Media{Id: MediaID}, nil)
				r.On("Delete", MediaID).Return(fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil, nil, test.mock)
			err := s.Delete(MediaID)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
		})
	}
}

func (t *MediaServiceTestSuite) TestService_DeleteFiles() {
	tt := map[string]struct {
		input domain.Media
		mock  func(r *repo.Repository, s *storage.Bucket)
		want  interface{}
	}{
		"Delete Success": {
			testMedia,
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Delete", testMedia.File.Id).Return(nil)
				s.On("Find", TestFileURLWebP).Return(nil, domain.File{}, fmt.Errorf("error"))
				r.On("Delete", testMedia.Id).Return(nil)
			},
			"Deleted original media item",
		},
		"Single Error": {
			testMedia,
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Delete", testMedia.File.Id).Return(fmt.Errorf("singular deleted"))
				s.On("Find", TestFileURLWebP).Return(nil, domain.File{}, fmt.Errorf("Error"))
				r.On("Delete", testMedia.Id).Return(nil)
			},
			"singular deleted",
		},
		"Size Error": {
			testMediaSizes,
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Delete", mock.Anything).Return(fmt.Errorf("size deleted"))
				s.On("Find", TestFileURLWebP).Return(nil, domain.File{}, fmt.Errorf("Error"))
				r.On("Delete", testMedia.Id).Return(nil)
			},
			"size deleted",
		},
		"Repo Error": {
			testMedia,
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Delete", testMedia.File.Id).Return(nil)
				s.On("Find", TestFileURLWebP).Return(nil, domain.File{}, fmt.Errorf("error"))
				r.On("Delete", testMedia.Id).Return(fmt.Errorf("repo error"))
			},
			"repo error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil, nil, test.mock)
			s.deleteFiles(test.input)
			got := t.LogWriter
			t.Contains(got.String(), test.want)
			t.Reset()
		})
	}
}

func (t *MediaServiceTestSuite) TestService_DeleteWebP() {
	file := domain.File{Id: 1, Url: "/uploads/test.jpg", Name: "test.jpg"}

	tt := map[string]struct {
		mock func(r *repo.Repository, s *storage.Bucket)
		want interface{}
	}{
		"Success": {
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Find", "/uploads/test.jpg.webp").Return(nil, file, nil)
				s.On("Delete", file.Id).Return(nil)
			},
			"Deleted WebP file",
		},

		"Find Error": {
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Find", "/uploads/test.jpg.webp").Return(nil, domain.File{}, fmt.Errorf("error"))
			},
			"",
		},
		"Delete Error": {
			func(r *repo.Repository, s *storage.Bucket) {
				s.On("Find", "/uploads/test.jpg.webp").Return(nil, file, nil)
				s.On("Delete", file.Id).Return(fmt.Errorf("webp error"))
			},
			"webp error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil, nil, test.mock)
			s.deleteWebP(file, true)
			got := t.LogWriter
			t.Contains(got.String(), test.want)
			t.Reset()
		})
	}
}
