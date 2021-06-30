// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"mime/multipart"
	"os"
)

func (t *MediaServiceTestSuite) TestClient_Validate() {
	tt := map[string]struct {
		input string
		cfg   domain.ThemeConfig
		opts  domain.Options
		want  interface{}
	}{
		"Nil File": {
			"",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"text/plain; charset=utf-8"},
				},
			},
			domain.Options{
				MediaUploadMaxHeight: 1,
			},
			"Error opening file with the name",
		},
		"Text File": {
			t.MediaPath + "/test.txt",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"text/plain; charset=utf-8"},
				},
			},
			domain.Options{
				MediaUploadMaxHeight: 1,
			},
			nil,
		},
		"Image": {
			t.MediaPath + "/gopher.png",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"image/png"},
				},
			},
			domain.Options{},
			nil,
		},
		"mime": {
			t.MediaPath + "/gopher.png",
			domain.ThemeConfig{},
			domain.Options{},
			"The file is not permitted to be uploaded",
		},
		"File Size": {
			t.MediaPath + "/gopher.png",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"image/png"},
				},
			},
			domain.Options{
				MediaUploadMaxSize: 1,
			},
			"The file exceeds the maximum size restriction",
		},
		"Image Width": {
			t.MediaPath + "/gopher.png",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"image/png"},
				},
			},
			domain.Options{
				MediaUploadMaxWidth: 1,
			},
			"The image exceeds the width/height restrictions",
		},
		"Image Height": {
			t.MediaPath + "/gopher.png",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"image/png"},
				},
			},
			domain.Options{
				MediaUploadMaxHeight: 1,
			},
			"The image exceeds the width/height restrictions",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.cfg, test.opts)

			var mt = &multipart.FileHeader{}
			if test.input != "" {
				mt = t.File(test.input)
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
			t.MediaPath + "/gopher.png",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"image/png"},
				},
			},
			nil,
		},
		"Bad mime": {
			t.MediaPath + "/gopher.png",
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
				Config:  &test.cfg,
				Options: &domain.Options{},
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
