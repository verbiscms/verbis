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
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/forms"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *FormsTestSuite) TestForms_Update() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			form,
			http.StatusOK,
			"Successfully updated form with ID: 123",
			form,
			func(m *mocks.Repository) {
				m.On("Update", form).Return(form, nil)
			},
			"/forms/123",
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "name", Message: "Name is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			formBadValidation,
			func(m *mocks.Repository) {
				m.On("Update", formBadValidation).Return(domain.Form{}, fmt.Errorf("error"))
			},
			"/forms/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to update the form",
			form,
			func(m *mocks.Repository) {
				m.On("Update", form).Return(domain.Form{}, fmt.Errorf("error"))
			},
			"/forms/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			form,
			func(m *mocks.Repository) {
				m.On("Update", form).Return(domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/forms/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"config",
			form,
			func(m *mocks.Repository) {
				m.On("Update", form).Return(domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "config"})
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
