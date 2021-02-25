// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// SiteTestSuite defines the helper used for site
// testing.
type SiteTestSuite struct {
	test.HandlerSuite
}

// TestSite
//
// Assert testing has begun.
func TestSite(t *testing.T) {
	suite.Run(t, &SiteTestSuite{
		HandlerSuite: test.TestSuite(),
	})
}

// Setup
//
// A helper to obtain a mock categories handler
// for testing.
func (t *SiteTestSuite) Setup(mf func(m *mocks.SiteRepository)) *Site {
	m := &mocks.SiteRepository{}
	if mf != nil {
		mf(m)
	}
	return &Site{
		Deps: &deps.Deps{
			Store: &models.Store{
				Site: m,
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
		Template: []map[string]interface{}{
			{
				"test": "testing",
			},
		},
	}
	// The default layouts used for testing.
	layouts = domain.Layouts{
		Layout: []map[string]interface{}{
			{
				"test": "testing",
			},
		},
	}
	// The default theme used for testing.
	theme = domain.ThemeConfig{
		Theme: domain.Theme{
			Title:       "Verbis",
			Description: "VerbisCMS",
			Version:     "0.1",
		},
	}
)
