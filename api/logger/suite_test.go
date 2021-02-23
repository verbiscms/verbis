// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// LoggerTestSuite defines the helper used for logger
// testing.
type LoggerTestSuite struct {
	suite.Suite
}

// TestLogger
//
// Assert testing has begun.
func TestLogger(t *testing.T) {
	suite.Run(t, new(LoggerTestSuite))
}

// TearDownTestSuite
//
// Teardown logging after testing.
func (t *LoggerTestSuite) TearDownTestSuite() {
	Init(&environment.Env{
		AppDebug: "true",
	})
}

// Setup
//
// Helper function for setting up the logger suite.
func (t *LoggerTestSuite) Setup() *bytes.Buffer {
	buf := &bytes.Buffer{}
	logger.SetLevel(logrus.TraceLevel)
	logger.SetOutput(buf)
	logger.SetFormatter(&Formatter{
		Colours: false,
		Debug: true,
	})
	return buf
}

// SetupHandler
//
// Helper function for setting up the handler
// testing.
func (t *LoggerTestSuite) SetupHandler(fn func(ctx *gin.Context)) *bytes.Buffer {
	buf := t.Setup()

	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	engine.Use(Middleware())

	engine.GET("/test", func(ctx *gin.Context) {
		fn(ctx)
		return
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	ctx.Request = req

	engine.ServeHTTP(rr, req)

	return buf
}

// SetupHooks
//
// Helper function for setting up the hooks
// testing.
func (t *LoggerTestSuite) SetupHooks(writer io.Writer) WriterHook {
	return WriterHook{
		Writer: writer,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
		},
	}
}
