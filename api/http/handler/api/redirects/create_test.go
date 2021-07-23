// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/store/redirects"
	"net/http"
)

func (t *RedirectsTestSuite) TestRedirects_Create() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
	}{
		"Success": {
			&redirect,
			http.StatusOK,
			"Successfully created redirect with ID: 123",
			redirect,
			func(m *mocks.Repository) {
				m.On("Create", redirect).Return(redirect, nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "to", Message: "To is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			redirectBadValidation,
			func(m *mocks.Repository) {
				m.On("Create", redirectBadValidation).Return(domain.Redirect{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			redirect,
			func(m *mocks.Repository) {
				m.On("Create", redirect).Return(domain.Redirect{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			redirect,
			func(m *mocks.Repository) {
				m.On("Create", redirect).Return(domain.Redirect{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			redirect,
			func(m *mocks.Repository) {
				m.On("Create", redirect).Return(domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/redirects", "/redirects", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Create(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
