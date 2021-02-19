// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *AuthTestSuite) TestAuth_ResetPassword() {

	var (
		rp = ResetPassword{
			NewPassword:     "password",
			ConfirmPassword: "password",
			Token:           "token",
		}
		rpdBadValidation = ResetPassword{
			NewPassword: "password",
			Token:       "token",
		}
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
			"Successfully reset password",
			rp,
			func(m *mocks.AuthRepository) {
				m.On("ResetPassword", rp.Token, rp.NewPassword).Return(nil)
			},
		},
		"Validation Failed": {
			api.ValidationErrJson{Errors: validation.Errors{{Key: "confirm_password", Message: "Confirm Password must equal the New Password.", Type: "eqfield"}}},
			400,
			"Validation failed",
			rpdBadValidation,
			func(m *mocks.AuthRepository) {
				m.On("ResetPassword", rpdBadValidation.Token, rpdBadValidation.NewPassword).Return(nil)
			},
		},
		"Not Found": {
			nil,
			400,
			"not found",
			rp,
			func(m *mocks.AuthRepository) {
				m.On("ResetPassword", rp.Token, rp.NewPassword).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			rp,
			func(m *mocks.AuthRepository) {
				m.On("ResetPassword", rp.Token, rp.NewPassword).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/reset", "/reset", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).ResetPassword(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}