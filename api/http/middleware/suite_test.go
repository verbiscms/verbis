// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/test"
	"net/http"
	"testing"
)

// MiddlewareTestSuite defines the helper used for middleware
// testing.
type MiddlewareTestSuite struct {
	test.HandlerSuite
}

// TestMiddleware asserts testing has begun.
func TestMiddleware(t *testing.T) {
	suite.Run(t, &MiddlewareTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// DefaultHandler is a helper func for returning data for testing.
func (t *MiddlewareTestSuite) DefaultHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "verbis")
}
