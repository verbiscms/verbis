// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
)

func (t *FormsTestSuite) TestForms_Send() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.FormRepository)
		url     string
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully sent form with ID: 123",
			formBody,
			func(m *mocks.FormRepository) {
				m.On("GetByUUID", "test").Return(form, nil)
				m.On("Send", &form, mock.Anything, mock.Anything).Return(nil)
			},
			"/forms/test",
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "name", Message: "Name is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			formBadValidation,
			func(m *mocks.FormRepository) {
				m.On("GetByUUID", "test").Return(formBadValidation, nil)
			},
			"/forms/test",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			form,
			func(m *mocks.FormRepository) {
				m.On("GetByUUID", "test").Return(domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/forms/test",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			form,
			func(m *mocks.FormRepository) {
				m.On("GetByUUID", "test").Return(domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/forms/test",
		},
		"Send Error": {
			nil,
			http.StatusInternalServerError,
			"error",
			form,
			func(m *mocks.FormRepository) {
				m.On("GetByUUID", "test").Return(form, nil)
				m.On("Send", &form, mock.Anything, mock.Anything).Return(&errors.Error{Code: errors.INTERNAL, Message: "error"})
			},
			"/forms/test",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, test.url, "/forms/:uuid", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Send(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
