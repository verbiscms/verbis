// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

func (t *ThemeTestSuite) TestTheme_Screenshot() {
	tt := map[string]struct {
		theme string
		file  string
		want  interface{}
	}{
		"SVG": {
			"verbis",
			"screenshot.svg",
			domain.Mime("image/svg+xml"),
		},
		"PNG": {
			"verbis2",
			"screenshot.png",
			domain.Mime("image/png"),
		},
		"Not Found": {
			"wrong",
			"screenshot.png",
			"Error finding screenshot with the path",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup()
			_, mime, err := s.Screenshot(test.theme, test.file)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, mime)
		})
	}
}
