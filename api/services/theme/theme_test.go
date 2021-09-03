// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	config "github.com/verbiscms/verbis/api/mocks/config"
	options "github.com/verbiscms/verbis/api/mocks/store/options"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// ThemeTestSuite defines the helper used for theme
// testing.
type ThemeTestSuite struct {
	suite.Suite
	TestPath string
}

// TestTheme asserts testing has begun.
func TestTheme(t *testing.T) {
	suite.Run(t, &ThemeTestSuite{})
}

const (
	ThemeName = "verbis"
)

// SetupSuite assigns the test path and default
// configuration.
func (t *ThemeTestSuite) SetupSuite() {
	wd, err := os.Getwd()
	t.NoError(err)
	t.TestPath = filepath.Join(wd, "testdata")
	logger.SetOutput(ioutil.Discard)
}

func (t *ThemeTestSuite) Setup(theme string) *Theme {
	o := &options.Repository{}
	c := &cache.Store{}
	cfg := domain.ThemeConfig{
		TemplateDir:   "templates",
		LayoutDir:     "layouts",
		FileExtension: ".cms",
	}
	c.On("Get", mock.Anything, configCacheKey, &domain.ThemeConfig{}).
		Return(nil).
		Run(func(args mock.Arguments) {
			arg := args.Get(2).(*domain.ThemeConfig)
			arg.TemplateDir = cfg.TemplateDir
			arg.LayoutDir = cfg.LayoutDir
			arg.FileExtension = cfg.FileExtension
		})
	o.On("GetTheme").Return(theme, nil)
	th := New(c, o)
	th.themesPath = t.TestPath
	return th
}

func (t *ThemeTestSuite) SetupMock(mf func(o *options.Repository, c *cache.Store, cf *config.Provider)) *Theme {
	o := &options.Repository{}
	c := &cache.Store{}
	cf := &config.Provider{}
	if mf != nil {
		mf(o, c, cf)
	}
	theme := New(c, o)
	theme.config = cf
	theme.themesPath = t.TestPath
	return theme
}

// SetupSuite assigns the test path and default
// configuration.
//func (t *ThemeTestSuite) Setup() *Theme {
//	return &Theme{
//		config: &domain.ThemeConfig{
//			FileExtension: ".cms",
//			LayoutDir:     "layouts",
//			TemplateDir:   "templates",
//		},
//		options:    nil,
//		themesPath: t.apiPath + ThemesPath,
//	}
//}

//func TestNew(t *testing.T) {
//	got := New()
//	assert.NotNil(t, got)
//}
//
//func (t *ThemeTestSuite) TestTheme_List() {
//	s := t.Setup()
//	got, _ := s.List("verbis")
//	want, _ := config.All(s.themesPath, "verbis")
//	t.Equal(got, want)
//}
//
//func (t *ThemeTestSuite) TestTheme_Find() {
//	s := t.Setup()
//	got, _ := s.Find("verbis")
//	want, _ := config.Find(s.themesPath + string(os.PathSeparator) + "verbis")
//	t.Equal(got, want)
//}
//
//func (t *ThemeTestSuite) TestTheme_Exists() {
//	tt := map[string]struct {
//		theme string
//		want  bool
//	}{
//		"True": {
//			"verbis",
//			true,
//		},
//		"False": {
//			"wrong",
//			false,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			s := t.Setup()
//			got := s.Exists(test.theme)
//			t.Equal(test.want, got)
//		})
//	}
//}
