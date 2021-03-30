// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/redirects"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *RedirectsTestSuite) TestCategories_Update() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			redirect,
			http.StatusOK,
			"Successfully updated redirect with ID: 123",
			redirect,
			func(m *mocks.Repository) {
				m.On("Update", redirect).Return(redirect, nil)
			},
			"/redirects/123",
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "to", Message: "To is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			redirectBadValidation,
			func(m *mocks.Repository) {
				m.On("Update", redirectBadValidation).Return(domain.Redirect{}, fmt.Errorf("error"))
			},
			"/redirects/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to update the redirect",
			redirect,
			func(m *mocks.Repository) {
				m.On("Update", redirectBadValidation).Return(domain.Redirect{}, fmt.Errorf("error"))
			},
			"/redirects/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			redirect,
			func(m *mocks.Repository) {
				m.On("Update", redirect).Return(domain.Redirect{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/redirects/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"config",
			redirect,
			func(m *mocks.Repository) {
				m.On("Update", redirect).Return(domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: "config"})
			},
			"/redirects/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPut, test.url, "/redirects/:id", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Update(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
