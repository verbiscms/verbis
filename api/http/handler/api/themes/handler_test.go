// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	mocks "github.com/verbiscms/verbis/api/mocks/services/theme"
	"github.com/verbiscms/verbis/api/test"
	"io/ioutil"
	"testing"
)

// ThemesTestSuite defines the helper used for site
// testing.
type ThemesTestSuite struct {
	test.HandlerSuite
	ThemePath string
}

// TestThemes
//
// Assert testing has begun.
func TestThemes(t *testing.T) {
	suite.Run(t, &ThemesTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
		ThemePath:    "/themes/",
	})
}

const (
	// The default active theme used for testing.
	TestActiveTheme = "verbis"
)

// Setup
//
// A helper to obtain a mock themes handler
// for testing.
func (t *ThemesTestSuite) Setup(mf func(m *mocks.Service)) *Themes {
	logger.SetOutput(ioutil.Discard)

	m := &mocks.Service{}
	if mf != nil {
		mf(m)
	}

	d := &deps.Deps{
		Config: &config.DefaultTheme,
		Options: &domain.Options{
			ActiveTheme: TestActiveTheme,
		},
		Theme: m,
	}

	return New(d)
}

var (
	// The default templates used for testing.
	templates = domain.Templates{
		domain.Template{
			Key:  "test",
			Name: "testing",
		},
	}
	// The default layouts used for testing.
	layouts = domain.Layouts{
		domain.Layout{
			Key:  "test",
			Name: "testing",
		},
	}
	// The default themes configs used for testing.
	themes = []domain.ThemeConfig{
		{
			Theme: domain.Theme{
				Title:       "Verbis",
				Description: "VerbisCMS",
				Version:     "0.1",
			},
		},
		{
			Theme: domain.Theme{
				Title:       "Verbis2",
				Description: "VerbisCMS2",
				Version:     "0.1",
			},
		},
	}
)
