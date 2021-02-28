// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const (
	// The configuration test path.
	TestPath = "/test/testdata/config/"
)

// ConfigTestSuite defines the helper used for config
// testing.
type ConfigTestSuite struct {
	suite.Suite
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
	logger.SetOutput(ioutil.Discard)
	wd, err := os.Getwd()
	t.NoError(err)
	t.apiPath = filepath.Join(filepath.Dir(wd), "")

	d := DefaultTheme
	d.Theme.Title = "test"
	t.config = d
}

func (t *ConfigTestSuite) TestInit() {
	got := Init(t.apiPath + TestPath)
	t.NotNil(cfg)
	t.Equal(t.config, *got)
}

func (t *ConfigTestSuite) TestGet() {
	Init(t.apiPath + TestPath)
	got := Get()
	t.NotNil(cfg)
	t.Equal(t.config, *got)
}

func (t *ConfigTestSuite) TestSet() {
	want := domain.ThemeConfig{
		AssetsPath: "assets",
	}
	Set(want)
	t.Equal(&want, cfg)
}

func (t *ConfigTestSuite) TestFetch() {

	tt := map[string]struct {
		path     string
		filename string
		want     interface{}
	}{
		"Success": {
			TestPath,
			FileName,
			t.config,
		},
		"Wrong Path": {
			"wrong",
			FileName,
			DefaultTheme,
		},
		"Bad Unmarshal": {
			TestPath,
			"/badconfig.yml",
			DefaultTheme,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := getThemeConfig(t.apiPath+test.path, test.filename)
			t.NotNil(cfg)
			t.Equal(test.want, *got)
		})
	}
}
