// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/errors"
)

func (t *SiteTestSuite) TestSite_Screenshot() {
	tt := map[string]struct {
		theme string
		file  string
		want  interface{}
	}{
		"SVG": {
			"verbis",
			"screenshot.svg",
			"image/svg+xml",
		},
		"PNG": {
			"verbis2",
			"screenshot.png",
			"image/png",
		},
		"Not Found": {
			"wrong",
			"screenshot.png",
			"Error finding screenshot with the path",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(ThemesPath, test.theme)
			_, mime, err := s.Screenshot(test.theme, test.file)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, mime)
		})
	}
}
