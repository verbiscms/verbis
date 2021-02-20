// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *RedirectsTestSuite) TestRedirects_Find() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.RedirectRepository)
		url     string
	}{
		"Success": {
			redirect,
			200,
			"Successfully obtained redirect with ID: 123",
			func(m *mocks.RedirectRepository) {
				m.On("GetById", int64(123)).Return(redirect, nil)
			},
			"/redirects/123",
		},
		"Invalid ID": {
			nil,
			400,
			"Pass a valid number to obtain the redirect by ID",
			func(m *mocks.RedirectRepository) {
				m.On("GetById", int64(123)).Return(domain.Redirect{}, fmt.Errorf("error"))
			},
			"/redirects/wrongid",
		},
		"Not Found": {
			nil,
			200,
			"no redirects found",
			func(m *mocks.RedirectRepository) {
				m.On("GetById", int64(123)).Return(domain.Redirect{}, &errors.Error{Code: errors.NOTFOUND, Message: "no redirects found"})
			},
			"/redirects/123",
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			func(m *mocks.RedirectRepository) {
				m.On("GetById", int64(123)).Return(domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/redirects/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, test.url, "/redirects/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Find(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
