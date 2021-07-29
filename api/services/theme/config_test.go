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
)

var (
	cfg1 = domain.ThemeConfig{
		Theme: domain.Theme{
			Name:        "verbis",
			Title:       "verbis",
			Description: "description",
		},
	}
	cfg2 = domain.ThemeConfig{
		Theme: domain.Theme{
			Name:        "verbis2",
			Title:       "verbis2",
			Description: "description",
		},
	}
)

func (t *ThemeTestSuite) TestTheme_Config() {
	tt := map[string]struct {
		mock func(o *options.Repository, c *cache.Store, cf *config.Provider)
		want interface{}
	}{
		"From Config": {
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				c.On("Get", mock.Anything, configCacheKey).
					Return(nil, fmt.Errorf("error"))
				o.On("GetTheme").
					Return("theme", nil)
				cf.On("Get", "theme").
					Return(cfg1, nil)
				c.On("Set", mock.Anything, configCacheKey, cfg1, mock.Anything)
			},
			cfg1,
		},
		"From Cache": {
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				c.On("Get", mock.Anything, configCacheKey).
					Return(cfg1, nil)
			},
			cfg1,
		},
		"Cast Error": {
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				c.On("Get", mock.Anything, configCacheKey).
					Return(100, nil)
				o.On("GetTheme").
					Return("", fmt.Errorf("error"))
			},
			"error",
		},
		"Options Error": {
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				c.On("Get", mock.Anything, configCacheKey).
					Return(domain.ThemeConfig{}, fmt.Errorf("error"))
				o.On("GetTheme").
					Return("", fmt.Errorf("error"))
			},
			"error",
		},
		"Config Error": {
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				c.On("Get", mock.Anything, configCacheKey).
					Return(nil, fmt.Errorf("error"))
				o.On("GetTheme").
					Return("theme", nil)
				cf.On("Get", "theme").
					Return(domain.ThemeConfig{}, fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			th := t.SetupMock(test.mock)
			got, err := th.Config()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *ThemeTestSuite) TestTheme_Set() {
	tt := map[string]struct {
		input string
		mock  func(o *options.Repository, c *cache.Store, cf *config.Provider)
		want  interface{}
	}{
		"Success": {
			"verbis",
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				o.On("SetTheme", "verbis").
					Return(nil)
				c.On("Delete", mock.Anything, configCacheKey)
				c.On("Get", mock.Anything, configCacheKey).
					Return(cfg1, nil)
			},
			cfg1,
		},
		"No Theme": {
			"wrong",
			nil,
			"no theme found",
		},
		"Set Error": {
			"verbis",
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				o.On("SetTheme", "verbis").
					Return(fmt.Errorf("error"))
			},
			"error",
		},
		"Config Error": {
			"verbis",
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				o.On("SetTheme", "verbis").
					Return(nil)
				c.On("Delete", mock.Anything, configCacheKey)
				c.On("Get", mock.Anything, configCacheKey).
					Return(nil, fmt.Errorf("error"))
				o.On("GetTheme").Return("", fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			th := t.SetupMock(test.mock)
			got, err := th.Set(test.input)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *ThemeTestSuite) TestTheme_Exists() {
	tt := map[string]struct {
		input string
		want  bool
	}{
		"True": {
			"verbis",
			true,
		},
		"False": {
			"wrong",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			th := t.Setup("")
			got := th.Exists(test.input)
			t.Equal(test.want, got)
		})
	}
}

func (t *ThemeTestSuite) TestTheme_Find() {
	tt := map[string]struct {
		input string
		mock  func(o *options.Repository, c *cache.Store, cf *config.Provider)
		want  interface{}
	}{
		"Success": {
			"theme",
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				cf.On("Get", "theme").
					Return(cfg1, nil)
			},
			cfg1,
		},
		"Error": {
			"theme",
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				cf.On("Get", "theme").
					Return(domain.ThemeConfig{}, fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			th := t.SetupMock(test.mock)
			got, err := th.Find(test.input)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *ThemeTestSuite) TestTheme_List() {
	tt := map[string]struct {
		path string
		mock func(o *options.Repository, c *cache.Store, cf *config.Provider)
		want interface{}
	}{
		"Wrong Path": {
			"theme",
			nil,
			"no such file or directory",
		},
		"Get Theme Error": {
			t.TestPath,
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				o.On("GetTheme").
					Return("", fmt.Errorf("error"))
			},
			"error",
		},
		"Config Error": {
			t.TestPath,
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				o.On("GetTheme").
					Return(mock.Anything, nil)
				cf.On("Get", mock.Anything).Return("verbis").
					Return(cfg1, fmt.Errorf("error"))
				cf.On("Get", mock.Anything).Return("verbis").
					Return(cfg1, fmt.Errorf("error"))
			},
			ErrNoThemes.Error(),
		},
		"Success": {
			t.TestPath,
			func(o *options.Repository, c *cache.Store, cf *config.Provider) {
				o.On("GetTheme").
					Return("verbis", nil)
				cf.On("Get", "empty").
					Return(domain.ThemeConfig{}, fmt.Errorf("error"))
				cf.On("Get", "verbis").
					Return(cfg1, nil)
				cf.On("Get", "verbis2").
					Return(cfg2, nil)
			},
			[]domain.ThemeConfig{
				{Theme: domain.Theme{Name: "verbis", Title: "verbis", Description: "description", Active: true}},
				{Theme: domain.Theme{Name: "verbis2", Title: "verbis2", Description: "description"}},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			th := t.SetupMock(test.mock)
			th.themesPath = test.path
			got, err := th.List()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
