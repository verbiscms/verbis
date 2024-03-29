// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/environment"
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

// TestLogger asserts testing has begun.
func TestLogger(t *testing.T) {
	suite.Run(t, new(LoggerTestSuite))
}

// TearDownTestSuite - Teardown logging after testing.
func (t *LoggerTestSuite) TearDownTestSuite() {
	Init(&environment.Env{
		AppDebug: "true",
	})
}

// Setup is a helper function for setting up the logger
// suite.
func (t *LoggerTestSuite) Setup() *bytes.Buffer {
	buf := &bytes.Buffer{}
	logger.SetLevel(logrus.TraceLevel)
	logger.SetOutput(buf)
	logger.SetFormatter(&Formatter{
		Colours: false,
		Debug:   true,
	})
	return buf
}

// SetupHandler is a helper function for setting up the
// http handler.
func (t *LoggerTestSuite) SetupHandler(fn func(ctx *gin.Context), url string) *bytes.Buffer {
	buf := t.Setup()

	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	engine.Use(Middleware())

	engine.GET(url, func(ctx *gin.Context) {
		fn(ctx)
	})

	req := httptest.NewRequest(http.MethodGet, url, nil)
	ctx.Request = req

	engine.ServeHTTP(rr, req)

	return buf
}

// SetupHooks is a helper function function for setting up
// the hooks for testing.
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
