// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"bytes"
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

// SetupHandler
//
// Helper function for setting up the handler
// testing.
func (t *LoggerTestSuite) SetupHandler(fn func(ctx *gin.Context)) *bytes.Buffer {
	buf := &bytes.Buffer{}
	logrus.SetOutput(buf)
	logrus.SetFormatter(&Formatter{
		Colours: false,
	})

	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	engine.Use(Handler())

	ctx.Set("test", "fuck")

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
		Writer:    writer,
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