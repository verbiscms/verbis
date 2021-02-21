// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MiddlewareTestSuite struct {
	suite.Suite
}


func TestMiddleware(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}

func (t *MiddlewareTestSuite) Setup(request string) (*httptest.Server, *http.Request, func()) {
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()

	server := httptest.NewServer(r)

	req, err := http.NewRequest(http.MethodGet, request, nil)
	t.NoError(err)

	return server, req, server.Close
}

func (t *MiddlewareTestSuite) t() {

}

func (t *MiddlewareTestSuite) TestCORS() {

	tt := map[string]struct {
		origin string
		want   string
	}{
		"Access Control": {
			origin: "Access-Control-Allow-Origin",
			want:   "*",
		},
		"Allow Credentials": {
			origin: "Access-Control-Allow-Credentials",
			want:   "true",
		},
		"Allow Headers": {
			origin: "Access-Control-Allow-Headers",
			want:   "access-control-allow-origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Origin, Cache-Control, X-Requested-With, token",
		},
		"Allow Methods": {
			origin: "Access-Control-Allow-Methods",
			want:   "POST, OPTIONS, GET, PUT, DELETE",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			gin.SetMode(gin.TestMode)
			gin.DefaultWriter = ioutil.Discard
			r := gin.Default()
			r.Use(CORS())

			t.Setup()


			server := httptest.NewServer(r)
			defer server.Close()

			client := &http.Client{}
			req, err := http.NewRequest("GET", "http://"+server.Listener.Addr().String()+"/api", nil)
			t.NoError(err)
			req.Header.Add("Origin", test.origin)

			get, err := client.Do(req)
			t.NoError(err)

			o := get.Header.Get(test.origin)
			t.Equal(test.want, o)
		})
	}
}

// TestCORS_AbortOptions - Test Cors abort options
func TestCORS_AbortOptions(t *testing.T) {
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()
	r.Use(CORS())

	server := httptest.NewServer(r)
	defer server.Close()

	client := &http.Client{}
	req, err := http.NewRequest("OPTIONS", "http://"+server.Listener.Addr().String()+"/api", nil)
	assert.NoError(t, err)

	opts, err := client.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.NoBody, opts.Body)
	assert.Equal(t, 200, opts.StatusCode)
}
