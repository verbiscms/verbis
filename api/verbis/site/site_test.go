// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// SiteTestSuite defines the helper used for site
// testing.
type SiteTestSuite struct {
	suite.Suite
	apiPath string
}

// TestSite
//
// Assert testing has begun.
func TestSite(t *testing.T) {
	suite.Run(t, &SiteTestSuite{})
}

const (
	// The default themes path used for testing.
	ThemesPath = "/test/testdata/themes"
)

func (t *SiteTestSuite) Setup(themes, path string) *Site {
	return &Site{
		config: &domain.ThemeConfig{
			Theme: domain.Theme{
				Name: path,
			},
			FileExtension: ".cms",
			LayoutDir:     "layouts",
			TemplateDir:   "templates",
		},
		options: nil,
		theme:   t.apiPath + themes,
	}
}

// SetupSuite
//
// Reassign API path for testing.
func (t *SiteTestSuite) SetupSuite() {
	logger.SetOutput(ioutil.Discard)
	wd, err := os.Getwd()
	t.NoError(err)

	apiPath := filepath.Join(filepath.Dir(wd), "../")
	t.apiPath = apiPath
}
