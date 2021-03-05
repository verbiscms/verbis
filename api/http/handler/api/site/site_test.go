// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	mocks "github.com/ainsleyclark/verbis/api/mocks/verbis/site"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
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
)

func (t *SiteTestSuite) TestSite_Global() {
	t.RequestAndServe(http.MethodGet, "/site", "/site", nil, func(ctx *gin.Context) {
		t.Setup(func(m *mocks.Repository) {
			m.On("Global").Return(site)
		}).Global(ctx)
	})
	t.RunT(site, http.StatusOK, "Successfully obtained site config")
}
