// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (t *MiddlewareTestSuite) TestEmptyBody() {
	tt := map[string]struct {
		method  string
		input   string
		status  int
		message string
		header  string
		content string
		want    string
	}{
		"Valid": {
			http.MethodDelete,
			`{verbis: "cms"}`,
			http.StatusOK,
			"",
			"application/json",
			"text/plain; charset=utf-8",
			"verbis",
		},
		"Not JSON": {
			http.MethodGet,
			"",
			http.StatusOK,
			"",
			"text/plain; charset=utf-8",
			"text/plain; charset=utf-8",
			"verbis",
		},
		"Empty Body": {
			http.MethodPost,
			"",
			http.StatusUnauthorized,
			"Empty JSON body",
			"application/json; charset=utf-8",
			"application/json; charset=utf-8",
			"",
		},
		"Invalid JSON": {
			http.MethodPost,
			"notjson",
			http.StatusUnauthorized,
			"Invalid JSON",
			"application/json; charset=utf-8",
			"application/json; charset=utf-8",
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Engine.Use(EmptyBody())

			t.Engine.GET("/test", t.DefaultHandler)
			t.Engine.PUT("/test", t.DefaultHandler)
			t.Engine.POST("/test", t.DefaultHandler)
			t.Engine.DELETE("/test", t.DefaultHandler)

			t.NewRequest(test.method, "/test", bytes.NewBuffer([]byte(test.input)))
			t.Context.Request.Header.Set("Content-Type", test.header)
			t.ServeHTTP()

			t.Equal(test.status, t.Recorder.Code)
			t.Equal(test.content, t.Recorder.Header().Get("content-type"))

			if test.message != "" {
				var body map[string]interface{}
				err := json.Unmarshal(t.Recorder.Body.Bytes(), &body)
				t.NoError(err)
				t.Equal(test.message, body["message"])
			} else {
				t.Equal(test.want, t.Recorder.Body.String())
			}

			t.Reset()
		})
	}
}

func (t *MiddlewareTestSuite) Test_isEmpty() {
	tt := map[string]struct {
		want  bool
		input interface{}
	}{
		"Empty": {
			true,
			nil,
		},
		"With Body": {
			false,
			"{}",
		},
		"With Body JSON": {
			false,
			`{body: "verbis"}`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			var got bool
			t.RequestAndServe(http.MethodGet, "/test", "/test", test.input, func(ctx *gin.Context) {
				body, err := ioutil.ReadAll(ctx.Request.Body)
				t.NoError(err)
				got = isEmpty(ctx, body)
			})
			t.Equal(test.want, got)
			t.Reset()
		})
	}
}

func (t *MiddlewareTestSuite) Test_isJSON() {
	tt := map[string]struct {
		want  bool
		input string
	}{
		"Empty": {
			false,
			"invalidjson",
		},
		"With Body": {
			true,
			"{}",
		},
		"With Body JSON": {
			true,
			`{"body": "verbis"}`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := isJSON(test.input)
			t.Equal(test.want, got)
		})
	}
}
