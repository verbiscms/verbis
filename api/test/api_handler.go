// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// controllerTest represents the suite of testing methods for controllers.
type controllerTest struct {
	testing  *testing.T
	recorder *httptest.ResponseRecorder
	ctx      *gin.Context
	engine   *gin.Engine
	respond  api.RespondJson
	got interface{}
}

// APITestSuite - New recorder for testing
// controllers, initialises gin & sets gin mode.
func APITestSuite(t *testing.T) *controllerTest {
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)

	return &controllerTest{
		testing:  t,
		recorder: rr,
		ctx:      ctx,
		engine:   engine,
	}
}

func (c *controllerTest) unmarshal(i interface{}) {
	responder := &api.RespondJson{}
	err := json.NewDecoder(c.recorder.Body).Decode(responder)
	if err != nil {
		c.testing.Error(err)
	}
	c.respond = *responder

	out, err := json.Marshal(responder.Data)
	if err != nil {
		c.testing.Error(err)
	}

	err = json.Unmarshal(out, i)
	if err != nil {
		c.testing.Fatal(err, "Unable to unmarshal")
	}

	c.got = i
}

// Run the API test.
func (c *controllerTest) Run(typ interface{}, want interface{}, status int, message string) {
	c.unmarshal(typ)
	assert.Equal(c.testing, status, c.recorder.Code)
	assert.Equal(c.testing, message, c.respond.Message)
	assert.Equal(c.testing, c.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")

	if reflect.ValueOf(want).IsZero() {
		assert.Equal(c.testing, typ, c.got)
		return
	}

	assert.Equal(c.testing, want, c.got)
}

// NewRequest makes a new http.Request and assigns the gin testing
// the request.
func (c *controllerTest) NewRequest(method string, url string, body io.Reader) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		c.testing.Fatal(err)
	}
	c.ctx.Request = req
}

// RequestAndServe makes a new http.Request and assigns the gin testing
// the request, serves HTTP.
func (c *controllerTest) RequestAndServe(method string, url string, engineUrl string, body io.Reader, handler func(ctx *gin.Context)) {
	switch method {
	case "GET":
		c.engine.GET(engineUrl, handler)
	case "POST":
		c.engine.POST(engineUrl, handler)
	case "PUT":
		c.engine.PUT(engineUrl, handler)
	case "DELETE":
		c.engine.DELETE(engineUrl, handler)
	}
	c.NewRequest(method, url, body)
	c.engine.ServeHTTP(c.recorder, c.ctx.Request)
}
