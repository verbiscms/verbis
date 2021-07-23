// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/store/redirects"
	"net/http"
)

func (t *RedirectsTestSuite) TestRedirects_Find() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			redirect,
			http.StatusOK,
			"Successfully obtained redirect with ID: 123",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(redirect, nil)
			},
			"/redirects/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"Pass a valid number to obtain the redirect by ID",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(domain.Redirect{}, fmt.Errorf("error"))
			},
			"/redirects/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"no redirects found",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(domain.Redirect{}, &errors.Error{Code: errors.NOTFOUND, Message: "no redirects found"})
			},
			"/redirects/123",
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
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
