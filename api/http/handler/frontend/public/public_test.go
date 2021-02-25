// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	mocks "github.com/ainsleyclark/verbis/api/mocks/publisher"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"testing"
)

// PublicTestSuite defines the helper used for public
// testing.
type PublicTestSuite struct {
	test.HandlerSuite
	bytes *[]byte
}

// TestPublic
//
// Assert testing has begun.
func TestPublic(t *testing.T) {
	b := []byte(testString)
	suite.Run(t, &PublicTestSuite{
		HandlerSuite: test.TestSuite(),
		bytes:        &b,
	})
}

// Setup
//
// A helper to obtain a public handler for testing.
func (t *PublicTestSuite) Setup(mf func(m *mocks.Renderer, ctx *gin.Context), ctx *gin.Context) *Public {
	m := &mocks.Renderer{}
	if mf != nil {
		mf(m, ctx)
	}
	return &Public{
		Publisher: m,
	}
}

var (
	// The default string used for testing.
	testString = "test"
)
