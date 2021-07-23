// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/store/forms"
	"net/http"
)

func (t *FormsTestSuite) TestForms_Send() {
	t.T().Skip("Skipping, not implemented Service yet")

	uniq := "9fc52ef0-914d-11eb-a8b3-0242ac130003"
	uniqParsed, err := uuid.Parse(uniq)
	t.NoError(err)

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully sent form with ID: 123",
			formBody,
			func(m *mocks.Repository) {
				m.On("FindByUUID", uniqParsed).Return(form, nil)
				m.On("Send", &form, mock.Anything, mock.Anything).Return(nil)
			},
			"/forms/" + uniq,
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "name", Message: "SizeName is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			formBadValidation,
			func(m *mocks.Repository) {
				m.On("FindByUUID", uniqParsed).Return(formBadValidation, nil)
			},
			"/forms/" + uniq,
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			form,
			func(m *mocks.Repository) {
				m.On("FindByUUID", uniqParsed).Return(domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/forms/" + uniq,
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			form,
			func(m *mocks.Repository) {
				m.On("FindByUUID", uniqParsed).Return(domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/forms/" + uniq,
		},
		"Send Error": {
			nil,
			http.StatusInternalServerError,
			"error",
			form,
			func(m *mocks.Repository) {
				m.On("FindByUUID", uniqParsed).Return(form, nil)
				m.On("Send", &form, mock.Anything, mock.Anything).Return(&errors.Error{Code: errors.INTERNAL, Message: "error"})
			},
			"/forms/" + uniq,
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
