// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"os"
	"path/filepath"
	"runtime"
)

func (t *ThemeTestSuite) TestTheme_Util() {
	if runtime.GOOS == "windows" {
		t.T().Skip("Skipping for pattern matches on windows.")
	}

	var s []string

	tt := map[string]struct {
		root    string
		pattern string
		want    interface{}
	}{
		"Bad Pattern": {
			t.apiPath + ThemesPath + string(os.PathSeparator) + "verbis",
			"\\",
			filepath.ErrBadPattern.Error(),
		},
		"Is Directory": {
			t.apiPath + ThemesPath + string(os.PathSeparator) + "empty",
			"*.cms",
			s,
		},
		"No Files": {
			"wrong",
			"",
			"no such file or directory",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup()
			got, err := s.walkMatch(test.root, test.pattern)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
