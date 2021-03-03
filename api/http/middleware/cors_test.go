// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"net/http"
)

func (t *MiddlewareTestSuite) TestCORS() {
	tt := map[string]struct {
		origin string
		want   string
	}{
		"Access Control": {
			"Access-Control-Allow-Origin",
			"*",
		},
		"Allow Credentials": {
			"Access-Control-Allow-Credentials",
			"true",
		},
		"Allow Headers": {
			"Access-Control-Allow-Headers",
			"access-control-allow-origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Origin, Cache-Control, X-Requested-With, token",
		},
		"Allow Methods": {
			"Access-Control-Allow-Methods",
			"POST, OPTIONS, GET, PUT, DELETE",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Engine.Use(CORS())

			t.NewRequest(http.MethodGet, "/test", nil)
			t.Context.Request.Header.Set("Origin", test.origin)
			t.ServeHTTP()

			got := t.Recorder.Header().Get(test.origin)
			t.Equal(test.want, got)
			t.Reset()
		})
	}
}

func (t *MiddlewareTestSuite) TestCORS_AbortOptions() {
	t.Engine.Use(CORS())

	t.NewRequest(http.MethodOptions, "/", nil)
	t.ServeHTTP()

	t.Empty(t.Recorder.Body)
	t.Equal(http.StatusOK, t.Recorder.Code)
}
