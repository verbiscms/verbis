// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	publisher "github.com/verbiscms/verbis/api/mocks/publisher"
	theme "github.com/verbiscms/verbis/api/mocks/services/theme"
	"github.com/verbiscms/verbis/api/test"
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
		HandlerSuite: test.NewHandlerSuite(),
		bytes:        &b,
	})
}

// Setup
//
// A helper to obtain a public handler for testing.
func (t *PublicTestSuite) Setup(mf func(m *publisher.Publisher, ctx *gin.Context), ctx *gin.Context) *Public {
	m := &publisher.Publisher{}
	if mf != nil {
		mf(m, ctx)
	}
	return &Public{
		Deps:      &deps.Deps{},
		publisher: m,
	}
}

// SetupTheme
//
// A helper to obtain a public handler for testing.
func (t *PublicTestSuite) SetupTheme(mf func(m *publisher.Publisher, t *theme.Service, ctx *gin.Context), ctx *gin.Context) *Public {
	m := &publisher.Publisher{}
	mt := &theme.Service{}
	if mf != nil {
		mf(m, mt, ctx)
	}
	return &Public{
		Deps: &deps.Deps{
			Theme: mt,
		},
		publisher: m,
	}
}

var (
	// The default string used for testing.
	testString = "test"
)
