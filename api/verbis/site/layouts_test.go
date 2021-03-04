// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
)

func (t *SiteTestSuite) TestSite_Layouts() {
	tt := map[string]struct {
		theme string
		want  interface{}
	}{
		"Success": {
			"verbis",
			domain.Layouts{
				{Key: "layout-hyphen", Name: "Layout Hyphen"},
				{Key: "layout", Name: "Layout"},
				{Key: "nested/layout-nested", Name: "Nested/Layout Nested"},
			},
		},
		"Wrong Path": {
			"wrong",
			"Error getting layouts with the path:",
		},
		"No Layouts": {
			"empty",
			"No layouts available",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(ThemesPath, test.theme)
			got, err := s.Layouts()
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
