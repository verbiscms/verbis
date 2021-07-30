// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"fmt"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	storage "github.com/verbiscms/verbis/api/mocks/services/storage"
	theme "github.com/verbiscms/verbis/api/mocks/services/theme"
	repo "github.com/verbiscms/verbis/api/mocks/store/media"
	"mime/multipart"
	"os"
	"path/filepath"
)

func (t *MediaServiceTestSuite) TestClient_Validate() {
	tt := map[string]struct {
		input string
		mock  func(r *repo.Repository, s *storage.Bucket, t *theme.Service)
		opts  domain.Options
		want  interface{}
	}{
		"Config Error": {
			"",
			func(r *repo.Repository, s *storage.Bucket, t *theme.Service) {
				t.On("Config").Return(domain.ThemeConfig{}, &errors.Error{Message: "config error"})
			},
			domain.Options{},
			"config error",
		},
		"Nil File": {
			"",
			func(r *repo.Repository, s *storage.Bucket, t *theme.Service) {
				t.On("Config").Return(domain.ThemeConfig{
					Media: domain.MediaConfig{
						AllowedFileTypes: []string{"text/plain; charset=utf-8"},
					},
				}, nil)
			},
			domain.Options{
				MediaUploadMaxHeight: 1,
			},
			"Error opening file with the name",
		},
		"Text File": {
			filepath.Join(t.TestDataPath, "/test.txt"),
			func(r *repo.Repository, s *storage.Bucket, t *theme.Service) {
				t.On("Config").Return(domain.ThemeConfig{
					Media: domain.MediaConfig{
						AllowedFileTypes: []string{"text/plain; charset=utf-8"},
					},
				}, nil)
			},
			domain.Options{
				MediaUploadMaxHeight: 1,
			},
			nil,
		},
		"Image": {
			filepath.Join(t.TestDataPath, "/gopher.png"),
			func(r *repo.Repository, s *storage.Bucket, t *theme.Service) {
				t.On("Config").Return(domain.ThemeConfig{
					Media: domain.MediaConfig{
						AllowedFileTypes: []string{"image/png"},
					},
				}, nil)
			},
			domain.Options{},
			nil,
		},
		"Mime": {
			filepath.Join(t.TestDataPath, "/gopher.png"),
			func(r *repo.Repository, s *storage.Bucket, t *theme.Service) {
				t.On("Config").Return(domain.ThemeConfig{}, nil)
			},
			domain.Options{},
			"The file is not permitted to be uploaded",
		},
		"File Size": {
			filepath.Join(t.TestDataPath, "/gopher.png"),
			func(r *repo.Repository, s *storage.Bucket, t *theme.Service) {
				t.On("Config").Return(domain.ThemeConfig{
					Media: domain.MediaConfig{
						AllowedFileTypes: []string{"image/png"},
					},
				}, nil)
			},
			domain.Options{
				MediaUploadMaxSize: 1,
			},
			"The file exceeds the maximum size restriction",
		},
		"Image Width": {
			filepath.Join(t.TestDataPath, "/gopher.png"),
			func(r *repo.Repository, s *storage.Bucket, t *theme.Service) {
				t.On("Config").Return(domain.ThemeConfig{
					Media: domain.MediaConfig{
						AllowedFileTypes: []string{"image/png"},
					},
				}, nil)
			},
			domain.Options{
				MediaUploadMaxWidth: 1,
			},
			"The image exceeds the width/height restrictions",
		},
		"Image Height": {
			filepath.Join(t.TestDataPath, "/gopher.png"),
			func(r *repo.Repository, s *storage.Bucket, t *theme.Service) {
				t.On("Config").Return(domain.ThemeConfig{
					Media: domain.MediaConfig{
						AllowedFileTypes: []string{"image/png"},
					},
				}, nil)
			},
			domain.Options{
				MediaUploadMaxHeight: 1,
			},
			"The image exceeds the width/height restrictions",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.opts, test.mock)

			var mt = &multipart.FileHeader{}
			if test.input != "" {
				multi, err := t.ToMultiPartE(test.input)
				if err != nil {
					fmt.Println(err)
				}
				mt = multi
			}

			got := c.Validate(mt)
			if got != nil {
				t.Contains(errors.Message(got), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *MediaServiceTestSuite) TestValidator_Mime() {
	tt := map[string]struct {
		input string
		cfg   domain.ThemeConfig
		want  interface{}
	}{
		"Success": {
			filepath.Join(t.TestDataPath, "/gopher.png"),
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"image/png"},
				},
			},
			nil,
		},
		"Bad mime": {
			filepath.Join(t.TestDataPath, "/gopher.png"),
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"text/plain; charset=utf-8"},
				},
			},
			ErrMimeType,
		},
		"File Error": {
			"",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"text/plain; charset=utf-8"},
				},
			},
			ErrMimeType,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := validator{
				Config:  test.cfg,
				Options: domain.Options{},
			}

			file, _ := os.Open(test.input) // Ignore on purpose
			v.File = file

			got := v.Mime()
			t.Equal(test.want, got)
		})
	}
}

func (t *MediaServiceTestSuite) TestValidator_Image_Error() {
	file, _ := os.Open("") // Ignore on purpose
	v := validator{
		File: file,
	}
	err := v.Image()
	t.Error(err)
}
