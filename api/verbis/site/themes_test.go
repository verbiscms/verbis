// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"os"
)

func (t *SiteTestSuite) TestSite_Themes() {
	tt := map[string]struct {
		root  string
		theme string
		want  interface{}
	}{
		"Success": {
			ThemesPath,
			"verbis",
			domain.Themes{
				domain.Theme{
					Title:      "test",
					Screenshot: "/themes/verbis/screenshot.svg",
					Name:       "verbis",
					Active:     true,
				},
				domain.Theme{
					Title:      "test",
					Screenshot: "/themes/verbis2/screenshot.png",
					Name:       "verbis2",
					Active:     false,
				},
			},
		},
		"Wrong Path": {
			"wrong",
			"wrong",
			"Error finding themes",
		},
		"No Themes": {
			ThemesPath + string(os.PathSeparator) + "empty",
			"",
			"No themes available",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(test.root, test.theme)
			s.options = &domain.Options{
				ActiveTheme: test.theme,
			}
			got, err := s.Themes()
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
