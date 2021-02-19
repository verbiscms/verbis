// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

const (
	JsonHeader = "application/json; charset=utf-8"
)

// HandlerSuite represents the suite of testing methods for controllers.
type HandlerSuite struct {
	suite.Suite
	Recorder *httptest.ResponseRecorder
	Context  *gin.Context
	Engine   *gin.Engine
}


// APITestSuite - New recorder for testing
// controllers, initialises gin & sets gin mode.
func APITestSuite() HandlerSuite {
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)

	return HandlerSuite{
		Recorder: rr,
		Context:      ctx,
		Engine:   engine,
	}
}

// Run the API test.
func (t *HandlerSuite) RunT(want interface{}, status int, message string) {
	defer func() {
		t.Reset()
	}()

	got, data := t.decode()
	t.Equal(message, got.Message)
	t.Equal(status, t.Status())
	t.Equal(JsonHeader, t.ContentType())
	t.JSONEq(t.marshalInput(want), data)
}

func (t *HandlerSuite) Status() int {
	return t.Recorder.Code
}

func (t *HandlerSuite) ContentType() string {
	return t.Recorder.Header().Get("Content-Type")
}

// NewRequest makes a new http.Request and assigns the gin testing
// the request.
func (t *HandlerSuite) NewRequest(method string, url string, body io.Reader) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fail("error creating http request", err)
	}
	t.Context.Request = req
}

// RequestAndServe makes a new http.Request and assigns the gin testing
// the request, serves HTTP.
func (t *HandlerSuite) RequestAndServe(method string, url string, engineUrl string, body io.Reader, handler func(ctx *gin.Context)) {
	switch method {
	case "GET":
		t.Engine.GET(engineUrl, handler)
	case "POST":
		t.Engine.POST(engineUrl, handler)
	case "PUT":
		t.Engine.PUT(engineUrl, handler)
	case "DELETE":
		t.Engine.DELETE(engineUrl, handler)
	}
	t.NewRequest(method, url, body)
	t.Engine.ServeHTTP(t.Recorder, t.Context.Request)
}

func (t *HandlerSuite) Reset() {
	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	t.Recorder = rr
	t.Context = ctx
	t.Engine = engine
}

func (t *HandlerSuite) marshalInput(i interface{}) string {
	str, ok := i.(string)
	if ok {
		return str
	}
	if i == nil {
		return "{}"
	}
	out, err := json.Marshal(i)
	if err != nil {
		t.Fail("error marshalling input", err)
	}
	return string(out)
}

// Run the API test.
func (t *HandlerSuite) decode() (api.RespondJson, string) {
	responder := &api.RespondJson{}
	err := json.NewDecoder(t.Recorder.Body).Decode(responder)
	if err != nil {
		t.Fail("error decoding body", err)
	}

	out, err := json.Marshal(responder.Data)
	if err != nil {
		t.Fail("error marshalling data", err)
	}

	return *responder, string(out)
}