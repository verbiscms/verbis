// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/domain"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	config "github.com/verbiscms/verbis/api/mocks/config"
	options "github.com/verbiscms/verbis/api/mocks/store/options"
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
			filepath.Join(t.TestPath, ThemeName),
			"\\",
			filepath.ErrBadPattern.Error(),
		},
		"Is Directory": {
			filepath.Join(t.TestPath, "empty"),
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
			s := t.Setup("")
			got, err := s.walkMatch(test.root, test.pattern)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *ThemeTestSuite) TestTheme_GetActiveTheme() {
	tt := map[string]struct {
		mock func(o *options.Repository, c *cache.Store, cf *config.Provider)
		want interface{}
	}{
		"Success": {
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				o.On("GetTheme").
					Return("theme", nil).
					Once()
				c.On("Get", mock.Anything, configCacheKey, &domain.ThemeConfig{}).
					Return(nil).
					Run(func(args mock.Arguments) {
						arg := args.Get(2).(*domain.ThemeConfig)
						arg.Theme = cfg1.Theme
					})
			},
			filepath.Join(t.TestPath, "theme"),
		},
		"Options Error": {
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				o.On("GetTheme").
					Return("", fmt.Errorf("error"))
			},
			"error",
		},
		"Config Error": {
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				o.On("GetTheme").
					Return("", nil).
					Once()
				c.On("Get", mock.Anything, configCacheKey, &domain.ThemeConfig{}).
					Return(fmt.Errorf("error"))
				o.On("GetTheme").
					Return("", fmt.Errorf("error")).
					Once()
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.SetupMock(test.mock)
			got, _, err := s.getActiveTheme()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
