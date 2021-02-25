// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import (
	"github.com/ainsleyclark/verbis/api/deps"
	mocks "github.com/ainsleyclark/verbis/api/mocks/publisher"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"testing"
)

// SEOTestSuite defines the helper used for seo
// testing.
type SEOTestSuite struct {
	test.HandlerSuite
	bytes *[]byte
}

// TestSEO
//
// Assert testing has begun.
func TestSEO(t *testing.T) {
	b := []byte(testString)
	suite.Run(t, &SEOTestSuite{
		HandlerSuite: test.TestSuite(),
		bytes:        &b,
	})
}

// Setup
//
// A helper to obtain a seo handler for testing.
func (t *SEOTestSuite) Setup(mf func(ms *mocks.Renderer, ctx *gin.Context), ctx *gin.Context) *SEO {
	m := &mocks.Renderer{}
	if mf != nil {
		mf(m, ctx)
	}
	return &SEO{
		Deps:      &deps.Deps{},
		Publisher: m,
	}
}

// SetupSitemap
//
// A helper to obtain a seo handler for testing
// for sitemap handle funcs.
func (t *SEOTestSuite) SetupSitemap(mf func(m *mocks.Renderer, ms *mocks.SiteMapper, ctx *gin.Context), ctx *gin.Context) *SEO {
	ms := &mocks.SiteMapper{}
	m := &mocks.Renderer{}
	m.On("SiteMap").Return(ms)
	if mf != nil {
		mf(m, ms, ctx)
	}
	return &SEO{
		Deps:      &deps.Deps{},
		Publisher: m,
	}
}

var (
	// The default string used for testing.
	testString = "test"
)
