// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"os"
	"path/filepath"
	"testing"
)

// ConfigTestSuite defines the helper used for config
// testing.
type ConfigTestSuite struct {
	suite.Suite
	TestPath string
	Config   domain.ThemeConfig
}

// TestConfig asserts testing has begun.
func TestConfig(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

// SetupSuite assigns the test path and default
// configuration.
func (t *ConfigTestSuite) SetupSuite() {
	wd, err := os.Getwd()
	t.NoError(err)
	t.TestPath = filepath.Join(wd, "testdata")

	d := DefaultTheme
	d.Theme = domain.Theme{
		Title:      "test",
		Name:       "testdata",
		Screenshot: "/themes/testdata/screenshot.png",
		Active:     false,
	}
	t.Config = d
}

func (t *ConfigTestSuite) Test_Get() {
	c := Config{ThemePath: t.TestPath}
	cfg, err := c.Get("")
	t.NoError(err)
	t.Equal(t.Config, cfg)
}

func (t *ConfigTestSuite) Test_GetThemeConfig() {
	tt := map[string]struct {
		filename string
		want     interface{}
	}{
		"Success": {
			FileName,
			t.Config,
		},
		"Wrong Path": {
			"wrong",
			"Error retrieving theme config file",
		},
		"Bad Unmarshal": {
			"badconfig.yml",
			"Syntax error in theme config file",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got, err := getThemeConfig(t.TestPath, test.filename)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
