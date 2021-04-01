// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	mocks "github.com/ainsleyclark/verbis/api/mocks/services/theme"
	options "github.com/ainsleyclark/verbis/api/mocks/store/options"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
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
func (t *ThemesTestSuite) Setup(mf func(m *mocks.Repository)) *Themes {
	logger.SetOutput(ioutil.Discard)

	m := &mocks.Repository{}
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

// SetupOptions
//
// A helper to obtain a mock themes handler
// with options for testing.
func (t *ThemesTestSuite) SetupOptions(mf func(m *mocks.Repository, mo *options.Repository)) *Themes {
	s := t.Setup(nil)

	m := &mocks.Repository{}
	mo := &options.Repository{}

	if mf != nil {
		mf(m, mo)
	}

	s.Store = &store.Repository{
		Options: mo,
	}
	s.Theme = m

	return s
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
	themes = []*domain.ThemeConfig{
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
