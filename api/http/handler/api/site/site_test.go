// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	mocks "github.com/ainsleyclark/verbis/api/mocks/site"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// SiteTestSuite defines the helper used for site
// testing.
type SiteTestSuite struct {
	test.HandlerSuite
	ThemePath string
}

// TestSite
//
// Assert testing has begun.
func TestSite(t *testing.T) {
	suite.Run(t, &SiteTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
		ThemePath:    "/themes/",
	})
}

// Setup
//
// A helper to obtain a mock categories handler
// for testing.
func (t *SiteTestSuite) Setup(mf func(m *mocks.Repository)) *Site {
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}
	return &Site{
		Deps: &deps.Deps{
			Site:   m,
			Config: &config.DefaultTheme,
			Paths: paths.Paths{
				Base: "",
			},
			Options: &domain.Options{
				ActiveTheme: "",
			},
		},
	}
}

var (
	// The default site config used for testing.
	site = domain.Site{
		Title:       "Verbis",
		Description: "VerbisCMS",
		Logo:        "/logo.svg",
		Url:         "verbiscms.com",
		Version:     "0.1",
	}
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
	// The default themes used for testing.
	themes = domain.Themes{
		{
			Title:       "Verbis",
			Description: "VerbisCMS",
			Version:     "0.1",
		},
		{
			Title:       "Verbis2",
			Description: "VerbisCMS2",
			Version:     "0.1",
		},
	}
)
