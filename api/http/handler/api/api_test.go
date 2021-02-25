// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

// ApiTestSuite defines the helper used for api
// testing.
type ApiTestSuite struct {
	test.HandlerSuite
}

// TestApi
//
// Assert testing has begun.
func TestApi(t *testing.T) {
	suite.Run(t, &ApiTestSuite{
		HandlerSuite: test.TestSuite(),
	})
}

func Test_Params(t *testing.T) {
	ctx := &gin.Context{}
	want := &params.Params{
		Stringer: &apiParams{ctx: ctx},
	}
	got := Params(ctx)
	assert.Equal(t, want.Stringer, got.Stringer)
}

func (t *ApiTestSuite) TestApiParams_Param() {

	tt := map[string]struct {
		query string
		param string
		want  string
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
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/test?"+test.query, "/login", nil, func(ctx *gin.Context) {
				got := Params(ctx).Param(test.param)
				t.Equal(test.want, got)
			})
			t.Reset()
		})
	}
}
