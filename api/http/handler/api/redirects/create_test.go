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
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *RedirectsTestSuite) TestCategories_Create() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.RedirectRepository)
	}{
		"Success": {
			&redirect,
			200,
			"Successfully created redirect with ID: 123",
			redirect,
			func(m *mocks.RedirectRepository) {
				m.On("Create", &redirect).Return(redirect, nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "to", Message: "To is required.", Type: "required"}}},
			400,
			"Validation failed",
			redirectBadValidation,
			func(m *mocks.RedirectRepository) {
				m.On("Create", &redirectBadValidation).Return(domain.Redirect{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			nil,
			400,
			"invalid",
			redirect,
			func(m *mocks.RedirectRepository) {
				m.On("Create", &redirect).Return(domain.Redirect{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			400,
			"conflict",
			redirect,
			func(m *mocks.RedirectRepository) {
				m.On("Create", &redirect).Return(domain.Redirect{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			redirect,
			func(m *mocks.RedirectRepository) {
				m.On("Create", &redirect).Return(domain.Redirect{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
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
