// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiParams(t *testing.T) {
	ctx := &gin.Context{}
	want := &Params{
		defaults:       Defaults{},
		Stringer:       &apiParams{ctx: ctx},
	}
	got := ApiParams(ctx, Defaults{})
	assert.Equal(t, want, got)
}

func TestApiParams_Param(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tt := map[string]struct {
		query  string
		param string
		want string
	}{
		"Page": {
			"page=2",
			"page",
			"2",
		},
		"Limit": {
			"limit=2",
			"limit",
			"2",
		},
		"Limit All": {
			"limit=all",
			"limit",
			"all",
		},
		"Order By": {
			"order_by=id",
			"order_by",
			"id",
		},
		"Order Direction": {
			"order_direction=name",
			"order_direction",
			"name",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			g, engine := gin.CreateTestContext(rr)

			req, err := http.NewRequest("GET", "/test?"+test.query, nil)
			assert.NoError(t, err)
			g.Request = req

			var got = ""
			engine.GET("/test", func(g *gin.Context) {
				got = ApiParams(g, Defaults{}).Param(test.param)
			})
			engine.ServeHTTP(rr, req)

			assert.Equal(t, test.want, got)
		})
	}
}