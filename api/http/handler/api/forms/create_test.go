// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

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

func (t *FormsTestSuite) TestForms_Create() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.FormRepository)
	}{
		"Success": {
			form,
			200,
			"Successfully created form with ID: 123",
			form,
			func(m *mocks.FormRepository) {
				m.On("Create", &form).Return(form, nil)
			},
		},
		"Validation Failed": {
			api.ValidationErrJson{Errors: validation.Errors{{Key: "name", Message: "Name is required.", Type: "required"}}},
			400,
			"Validation failed",
			formBadValidation,
			func(m *mocks.FormRepository) {
				m.On("Create", &formBadValidation).Return(domain.Form{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			nil,
			400,
			"invalid",
			form,
			func(m *mocks.FormRepository) {
				m.On("Create", &form).Return(domain.Form{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			400,
			"conflict",
			form,
			func(m *mocks.FormRepository) {
				m.On("Create", &form).Return(domain.Form{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			form,
			func(m *mocks.FormRepository) {
				m.On("Create", &form).Return(domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
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