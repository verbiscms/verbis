// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/stretchr/testify/suite"
	"os"
	"path/filepath"
	"testing"
)

const (
	// The singular configuration test path.
	TestSingularPath = "/test/testdata/themes/verbis"
	// The multiple themes path.
	TestMultiPath = "/test/testdata/themes"
)

// ConfigTestSuite defines the helper used for config
// testing.
type ConfigTestSuite struct {
	suite.Suite
	logger  *bytes.Buffer
	apiPath string
	config  domain.ThemeConfig
}

// TestConfig
//
// Assert testing has begun.
func TestConfig(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

// SetupSuite
//
// Reassign API path for testing.
func (t *ConfigTestSuite) SetupSuite() {
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	t.logger = buf

	wd, err := os.Getwd()
	t.NoError(err)
	t.apiPath = filepath.Join(filepath.Dir(wd), "")

	d := DefaultTheme
	d.Theme = domain.Theme{
		Title:      "test",
		Screenshot: "/themes/verbis/screenshot.svg",
		Name:       "verbis",
		Active:     true,
	}

	t.config = d
}

func (t *ConfigTestSuite) Test_Init() {
	got := Init(t.apiPath + TestSingularPath)
	t.NotNil(cfg)
	t.Equal(t.config, *got)
}

func (t *ConfigTestSuite) Test_Get() {
	Init(t.apiPath + TestSingularPath)
	got := Get()
	t.NotNil(cfg)
	t.Equal(t.config, *got)
}

func (t *ConfigTestSuite) Test_Set() {
	want := domain.ThemeConfig{
		AssetsPath: "assets",
	}
	Set(want)
	t.Equal(&want, cfg)
}

func (t *ConfigTestSuite) Test_Fetch() {
	got := Fetch("wrong")
	want := "no such file or directory"
	t.Contains(t.logger.String(), want)
	t.Equal(&DefaultTheme, got)
}

func (t *ConfigTestSuite) Test_Find() {
	tt := map[string]struct {
		path     string
		filename string
		want     interface{}
		err      string
	}{
		"Success": {
			TestSingularPath,
			FileName,
			t.config,
			"",
		},
		"Wrong Path": {
			"wrong",
			FileName,
			DefaultTheme,
			"Error retrieving theme config file",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got, err := Find(t.apiPath + test.path)
			if err != nil {
				t.Contains(errors.Message(err), test.err)
				return
			}
			t.NotNil(cfg)
			t.Equal(test.want, *got)
		})
	}
}

func (t *ConfigTestSuite) Test_All() {
	tt := map[string]struct {
		path  string
		theme string
		want  domain.Themes
		err   string
	}{
		"Success": {
			TestMultiPath,
			"verbis",
			[]domain.Theme{
				{
					Title:      "test",
					Screenshot: "/themes/verbis/screenshot.svg",
					Name:       "verbis",
					Active:     true,
				},
				{
					Title:      "test",
					Screenshot: "/themes/verbis2/screenshot.png",
					Name:       "verbis2",
					Active:     false,
				},
			},
			"",
		},
		"Wrong Path": {
			"wrong",
			"wrong",
			nil,
			"Error finding themes",
		},
		"No Themes": {
			TestMultiPath + string(os.PathSeparator) + "empty",
			"",
			nil,
			"No themes available",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got, err := All(t.apiPath+test.path, "verbis")
			if err != nil {
				t.Contains(errors.Message(err), test.err)
				return
			}
			for i, v := range got {
				t.Equal(test.want[i], v.Theme)
			}
		})
	}
}

func (t *ConfigTestSuite) Test_GetThemeConfig() {
	tt := map[string]struct {
		path     string
		filename string
		want     interface{}
		err      string
	}{
		"Success": {
			TestSingularPath,
			FileName,
			t.config,
			"",
		},
		"Wrong Path": {
			"wrong",
			FileName,
			DefaultTheme,
			"Error retrieving theme config file",
		},
		"Bad Unmarshal": {
			TestSingularPath,
			"/badconfig.yml",
			DefaultTheme,
			"Syntax error in theme config file",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got, err := getThemeConfig(t.apiPath+test.path, test.filename)
			if err != nil {
				t.Contains(errors.Message(err), test.err)
				return
			}
			t.NotNil(cfg)
			t.Equal(test.want, *got)
		})
	}
}
