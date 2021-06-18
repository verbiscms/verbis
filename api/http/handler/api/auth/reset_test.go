// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	events "github.com/ainsleyclark/verbis/api/mocks/events"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/auth"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
)

func (t *AuthTestSuite) TestAuth_SendResetPassword() {
	var (
		srp = SendResetPassword{
			Email: "info@verbiscms.com",
		}
		srpBadValidation = SendResetPassword{}
		dispatchSuccess  = func(m *events.Dispatcher) {
			m.On("Dispatch", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		}
	)

	tt := map[string]struct {
		want       interface{}
		status     int
		message    string
		input      interface{}
		dispatcher func(m *events.Dispatcher)
		mock       func(m *mocks.Repository)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"A fresh verification link has been sent to your email",
			srp,
			dispatchSuccess,
			func(m *mocks.Repository) {
				m.On("SendResetPassword", srp.Email).Return(user.UserPart, "token", nil)
			},
		},
		"Validation Failed": {
			`{"errors":[{"key":"email","message":"Email is required.","type":"required"}]}`,
			http.StatusBadRequest,
			"Validation failed",
			srpBadValidation,
			dispatchSuccess,
			func(m *mocks.Repository) {
				m.On("SendResetPassword", srp.Email).Return(user.UserPart, "token", nil)
			},
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			srp,
			dispatchSuccess,
			func(m *mocks.Repository) {
				m.On("SendResetPassword", srp.Email).Return(domain.UserPart{}, "", &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			srp,
			dispatchSuccess,
			func(m *mocks.Repository) {
				m.On("SendResetPassword", srp.Email).Return(domain.UserPart{}, "", &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
		"Dispatch Error": {
			nil,
			http.StatusInternalServerError,
			"dispatch",
			srp,
			func(m *events.Dispatcher) {
				m.On("Dispatch", mock.Anything, mock.Anything, mock.Anything).Return(&errors.Error{Code: errors.INTERNAL, Message: "dispatch"})
			},
			func(m *mocks.Repository) {
				m.On("SendResetPassword", srp.Email).Return(user.UserPart, "token", nil)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/sendreset", "/sendreset", test.input, func(ctx *gin.Context) {
				t.SetupDispatcher(test.mock, test.dispatcher).SendResetPassword(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
