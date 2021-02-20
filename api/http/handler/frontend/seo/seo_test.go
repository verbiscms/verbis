// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import (
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/render"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"testing"
)

// SEOTestSuite defines the helper used for seo
// testing.
type SEOTestSuite struct {
	api.HandlerSuite
	bytes *[]byte
}

// TestSEO
//
// Assert testing has begun.
func TestSEO(t *testing.T) {
	b := []byte(testString)
	suite.Run(t, &SEOTestSuite{
		HandlerSuite: api.TestSuite(),
		bytes: &b,
	})
}

// Setup
//
// A helper to obtain a seo handler for testing.
func (t *SEOTestSuite) Setup(mf func(m *mocks.Renderer, ctx *gin.Context), ctx *gin.Context) *SEO {
	m := &mocks.Renderer{}
	if mf != nil {
		mf(m, ctx)
	}
	return &SEO{
		Publisher: m,
	}
}

var (
	// The default string used for testing.
	testString = "test"
)