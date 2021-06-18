// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/redirects"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *RedirectsTestSuite) TestRedirects_Delete() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully deleted redirect with ID: 123",
			func(m *mocks.Repository) {
				m.On("Delete", 123).Return(nil)
			},
			"/redirects/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to delete a redirect",
			func(m *mocks.Repository) {
				m.On("Delete", 123).Return(nil)
			},
			"/redirects/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			func(m *mocks.Repository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/redirects/123",
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			func(m *mocks.Repository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
			"/redirects/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Repository) {
				m.On("Delete", 123).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/redirects/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodDelete, test.url, "/redirects/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Delete(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
