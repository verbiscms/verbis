// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"fmt"
	validation "github.com/ainsleyclark/verbis/api/common/vaidation"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/forms"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *FormsTestSuite) TestForms_Create() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
	}{
		"Success": {
			form,
			http.StatusOK,
			"Successfully created form with ID: 123",
			form,
			func(m *mocks.Repository) {
				m.On("Create", form).Return(form, nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "name", Message: "Name is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			formBadValidation,
			func(m *mocks.Repository) {
				m.On("Create", formBadValidation).Return(domain.Form{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			nil,
			http.StatusBadRequest,
			"invalid",
			form,
			func(m *mocks.Repository) {
				m.On("Create", form).Return(domain.Form{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			http.StatusBadRequest,
			"conflict",
			form,
			func(m *mocks.Repository) {
				m.On("Create", form).Return(domain.Form{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			form,
			func(m *mocks.Repository) {
				m.On("Create", form).Return(domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/forms", "/forms", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Create(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
