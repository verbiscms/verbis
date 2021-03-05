// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theme

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// ThemeTestSuite defines the helper used for site
// testing.
type ThemeTestSuite struct {
	suite.Suite
	apiPath string
}

// TestTheme
//
// Assert testing has begun.
func TestTheme(t *testing.T) {
	suite.Run(t, &ThemeTestSuite{})
}

const (
	// The default themes path used for testing.
	ThemesPath = "/test/testdata/themes"
)

// Setup
//
// Sets defaults and returns a new Theme repo.
func (t *ThemeTestSuite) Setup() *theme {
	return &theme{
		config: &domain.ThemeConfig{
			FileExtension: ".cms",
			LayoutDir:     "layouts",
			TemplateDir:   "templates",
		},
		options:    nil,
		themesPath: t.apiPath + ThemesPath,
	}
}

// SetupSuite
//
// Reassign API path for testing.
func (t *ThemeTestSuite) SetupSuite() {
	logger.SetOutput(ioutil.Discard)
	wd, err := os.Getwd()
	t.NoError(err)

	apiPath := filepath.Join(filepath.Dir(wd), "../")
	t.apiPath = apiPath
}

func (t *ThemeTestSuite) Test_New() {
	t.T().Skip()
}

func (t *ThemeTestSuite) TestTheme_List() {
	s := t.Setup()
	got, _ := s.List("verbis")
	want, _ := config.All(s.themesPath, "verbis")
	t.Equal(got, want)
}

func (t *ThemeTestSuite) TestTheme_Find() {
	s := t.Setup()
	got, _ := s.Find("verbis")
	want, _ := config.Find(s.themesPath + string(os.PathSeparator) + "verbis")
	t.Equal(got, want)
}

func (t *ThemeTestSuite) TestTheme_Exists() {
	tt := map[string]struct {
		theme string
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
			s := t.Setup()
			got := s.Exists(test.theme)
			t.Equal(test.want, got)
		})
	}
}
