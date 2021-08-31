// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/test"
	"net/http"
	"testing"
)

// MiddlewareTestSuite defines the helper used for middleware
// testing.
type MiddlewareTestSuite struct {
	test.HandlerSuite
	LogWriter bytes.Buffer
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

// BeforeTest assign the logger to a buffer.
func (t *MiddlewareTestSuite) BeforeTest(suiteName, testName string) {
	b := bytes.Buffer{}
	t.LogWriter = b
	logger.SetOutput(&t.LogWriter)
	logger.SetLevel(logrus.TraceLevel)
}

// Reset the log writer and handler suite.
func (t *MiddlewareTestSuite) Reset() {
	t.HandlerSuite.Reset()
	t.LogWriter.Reset()
}
