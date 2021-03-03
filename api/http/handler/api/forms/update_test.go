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

func (t *FormsTestSuite) TestForms_Update() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.FormRepository)
		url     string
	}{
		"Success": {
			form,
			200,
			"Successfully updated form with ID: 123",
			form,
			func(m *mocks.FormRepository) {
				m.On("Update", &form).Return(form, nil)
			},
			"/forms/123",
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "name", Message: "Name is required.", Type: "required"}}},
			400,
			"Validation failed",
			formBadValidation,
			func(m *mocks.FormRepository) {
				m.On("Update", formBadValidation).Return(domain.Form{}, fmt.Errorf("error"))
			},
			"/forms/123",
		},
		"Invalid ID": {
			nil,
			400,
			"A valid ID is required to update the form",
			form,
			func(m *mocks.FormRepository) {
				m.On("Update", form).Return(domain.Form{}, fmt.Errorf("error"))
			},
			"/forms/wrongid",
		},
		"Not Found": {
			nil,
			400,
			"not found",
			form,
			func(m *mocks.FormRepository) {
				m.On("Update", &form).Return(domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/forms/123",
		},
		"Internal": {
			nil,
			500,
			"internal",
			form,
			func(m *mocks.FormRepository) {
				m.On("Update", &form).Return(domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/forms/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPut, test.url, "/forms/:id", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Update(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
