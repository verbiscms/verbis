// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *AuthTestSuite) TestAuth_SendResetPassword() {

	var (
		srp = SendResetPassword{
			Email: "info@verbiscms.com",
		}
		srpBadValidation = SendResetPassword{}
	)

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.AuthRepository)
	}{
		"Success": {
			nil,
			200,
			"A fresh verification link has been sent to your email",
			srp,
			func(m *mocks.AuthRepository) {
				m.On("SendResetPassword", srp.Email).Return(nil)
			},
		},
		"Validation Failed": {
			`{"errors":[{"key":"email","message":"Email is required.","type":"required"}]}`,
			400,
			"Validation failed",
			srpBadValidation,
			func(m *mocks.AuthRepository) {
				m.On("SendResetPassword", srpBadValidation.Email).Return(nil)
			},
		},
		"Not Found": {
			nil,
			400,
			"not found",
			srp,
			func(m *mocks.AuthRepository) {
				m.On("SendResetPassword", srp.Email).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			srp,
			func(m *mocks.AuthRepository) {
				m.On("SendResetPassword", srp.Email).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/sendreset", "/sendreset", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).SendResetPassword(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}