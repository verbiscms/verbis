// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/common/paths"
	"github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/services/site"
	"github.com/verbiscms/verbis/api/test"
	"net/http"
	"testing"
)

// SiteTestSuite defines the helper used for site
// testing.
type SiteTestSuite struct {
	test.HandlerSuite
	ThemePath string
}

// TestSite asserts testing has begun.
func TestSite(t *testing.T) {
	suite.Run(t, &SiteTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
		ThemePath:    "/themes/",
	})
}

// Setup is a helper to obtain a mock site handler
// for testing.
func (t *SiteTestSuite) Setup(mf func(m *mocks.Service)) *Site {
	m := &mocks.Service{}
	if mf != nil {
		mf(m)
	}
	d := &deps.Deps{
		Site:   m,
		Config: &config.DefaultTheme,
		Paths: paths.Paths{
			Base: "",
		},
		Options: &domain.Options{
			ActiveTheme: "",
		},
	}
	return New(d)
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
		t.Setup(func(m *mocks.Service) {
			m.On("Global").Return(site)
		}).Global(ctx)
	})
	t.RunT(site, http.StatusOK, "Successfully obtained site config")
}
