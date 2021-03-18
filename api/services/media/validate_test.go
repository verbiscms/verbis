// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"mime/multipart"
)

func (t *MediaTestSuite) TestClient_Validate() {
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
			t.mediaPath + "/test.txt",
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
			t.mediaPath + "/gopher.png",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"image/png"},
				},
			},
			domain.Options{},
			nil,
		},
		"Mime": {
			t.mediaPath + "/gopher.png",
			domain.ThemeConfig{},
			domain.Options{},
			"The file is not permitted to be uploaded",
		},
		"File Size": {
			t.mediaPath + "/gopher.png",
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
			t.mediaPath + "/gopher.png",
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
			t.mediaPath + "/gopher.png",
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
