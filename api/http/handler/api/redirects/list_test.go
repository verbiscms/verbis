// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	mocks "github.com/verbiscms/verbis/api/mocks/store/redirects"
	"github.com/verbiscms/verbis/api/test/dummy"
	"net/http"
)

func (t *RedirectsTestSuite) TestRedirects_List() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Repository)
	}{
		"Success": {
			redirects,
			http.StatusOK,
			"Successfully obtained redirects",
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams).Return(redirects, 1, nil)
			},
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"no redirects found",
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no redirects found"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams).Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams).Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, "/redirects", "/redirects", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).List(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
