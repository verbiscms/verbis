// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"bytes"
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

const (
	// The header to be expected to be received from the API.
	JSONHeader = "application/json; charset=utf-8"
)

// RespondJSON is an abstraction of the api.RespondJSON
type RespondJSON struct {
	Status  int    `json:"status"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Meta    struct {
		RequestTime  string      `json:"request_time"`
		ResponseTime string      `json:"response_time"`
		LatencyTime  string      `json:"latency_time"`
		Pagination   interface{} `json:"pagination,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

// HandlerSuite represents the suite of testing methods for controllers.
type HandlerSuite struct {
	suite.Suite
	Recorder *httptest.ResponseRecorder
	Context  *gin.Context
	Engine   *gin.Engine
}

// NewHandlerSuite
//
// New recorder for testing
// controllers, initialises gin & sets gin mode.
func NewHandlerSuite() HandlerSuite {
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	cache.Init()

	return HandlerSuite{
		Recorder: rr,
		Context:  ctx,
		Engine:   engine,
	}
}

// RunT
//
// Run the API test.
func (t *HandlerSuite) RunT(want interface{}, status int, message string) {
	defer func() {
		t.Reset()
	}()

	got, data := t.decode()
	t.Equal(message, got.Message)
	t.Equal(status, t.Status())
	t.Equal(JSONHeader, t.ContentType())
	t.JSONEq(t.marshalWant(want), data)
}

// RespondData
//
// Returns the RespondJSON and the data decoded in
// string form.
func (t *HandlerSuite) RespondData() (RespondJSON, string) {
	return t.decode()
}

// Status
//
// Returns the recorder's status code.
func (t *HandlerSuite) Status() int {
	return t.Recorder.Code
}

// ContentType
//
// Returns the recorder's content type.
func (t *HandlerSuite) ContentType() string {
	return t.Recorder.Header().Get("Content-Type")
}

// NewRequest
//
// Creates a new http.Request and assigns the gin testing
// the request.
func (t *HandlerSuite) NewRequest(method string, url string, body io.Reader) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fail("error creating http request", err)
	}
	t.Context.Request = req
}

func (t *HandlerSuite) ServeHTTP() {
	t.Engine.ServeHTTP(t.Recorder, t.Context.Request)
}

// RequestAndServe
//
// Makes a new http.Request and assigns the gin testing
// the request, serves HTTP.
func (t *HandlerSuite) RequestAndServe(method string, url string, engineURL string, body interface{}, handler func(ctx *gin.Context)) {
	switch method {
	case http.MethodGet:
		t.Engine.GET(engineURL, handler)
	case http.MethodPost:
		t.Engine.POST(engineURL, handler)
	case http.MethodPut:
		t.Engine.PUT(engineURL, handler)
	case http.MethodDelete:
		t.Engine.DELETE(engineURL, handler)
	}
	t.NewRequest(method, url, t.marshalInput(body))
	t.Engine.ServeHTTP(t.Recorder, t.Context.Request)
}

// Reset
//
// Sets up a new recorder, engine and context upon
// test completion.
func (t *HandlerSuite) Reset() {
	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	t.Recorder = rr
	t.Context = ctx
	t.Engine = engine
}

// marshalInput
//
// Convert the interface{} to a new bytes.Buffer
// for comparison of the input.
func (t *HandlerSuite) marshalInput(i interface{}) *bytes.Buffer {
	if i == nil {
		return &bytes.Buffer{}
	}

	body, err := json.Marshal(i)
	if err != nil {
		t.Fail("error marshalling input")
	}

	return bytes.NewBuffer(body)
}

// marshalWant
//
// Marshal the test want, if the test want arg
// is is nil, return an empty JSON object.
func (t *HandlerSuite) marshalWant(i interface{}) string {
	str, ok := i.(string)
	if ok {
		return str
	}
	if i == nil {
		return "{}"
	}
	out, err := json.Marshal(i)
	if err != nil {
		t.Fail("error marshalling want", err)
	}
	return string(out)
}

// decode
//
// Unmarshal the result of the recorder into a
// RespondJSON ready for testing.
func (t *HandlerSuite) decode() (RespondJSON, string) {
	responder := RespondJSON{}
	err := json.NewDecoder(t.Recorder.Body).Decode(&responder)
	if err != nil {
		t.Fail("error decoding body", err)
	}

	out, err := json.Marshal(responder.Data)
	if err != nil {
		t.Fail("error marshalling data", err)
	}

	return responder, string(out)
}
