// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import "github.com/ainsleyclark/verbis/api/domain"

func (t *MediaTestSuite) TestClient_Validate() {
	tt := map[string]struct {
		input string
		cfg   domain.ThemeConfig
		opts  domain.Options
		want  interface{}
	}{
		"Success": {
			"/gopher.png",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"image/png"},
				},
			},
			domain.Options{},
			nil,
		},
		"Bad Mime": {
			"/gopher.png",
			domain.ThemeConfig{},
			domain.Options{},
			ErrMimeType.Error(),
		},
		"Bad File Size": {
			"/gopher.png",
			domain.ThemeConfig{
				Media: domain.MediaConfig{
					AllowedFileTypes: []string{"image/png"},
				},
			},
			domain.Options{
				MediaUploadMaxSize: 120,
			},
			"mimetype is not permitted",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			c := t.Setup(test.cfg, test.opts)
			mt := t.File(test.input)
			got := c.Validate(mt)
			if got != nil {
				t.Contains(got.Error(), test.want)
			}
			t.Nil(got)
		})
	}
}
