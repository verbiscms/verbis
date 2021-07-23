// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/gin-gonic/gin"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/store/auth"
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
		mock    func(m *mocks.Repository)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully reset password",
			rp,
			func(m *mocks.Repository) {
				m.On("ResetPassword", rp.Token, rp.NewPassword).Return(nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "confirm_password", Message: "Confirm Password must equal the New Password.", Type: "eqfield"}}},
			http.StatusBadRequest,
			"Validation failed",
			rpdBadValidation,
			func(m *mocks.Repository) {
				m.On("ResetPassword", rpdBadValidation.Token, rpdBadValidation.NewPassword).Return(nil)
			},
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			rp,
			func(m *mocks.Repository) {
				m.On("ResetPassword", rp.Token, rp.NewPassword).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			rp,
			func(m *mocks.Repository) {
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
