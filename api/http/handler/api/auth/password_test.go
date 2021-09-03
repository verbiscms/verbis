// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
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
		mock    func(m *mocks.Repository, c *cache.Store)
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully reset password",
			rp,
			func(m *mocks.Repository, c *cache.Store) {
				m.On("ResetPassword", user.Email, rp.NewPassword).Return(nil)
				c.On("Get", mock.Anything, rp.Token, &domain.User{}).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*domain.User)
					arg.Password = user.Password
					arg.Email = user.Email
				}).Return(nil)
			},
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "confirm_password", Message: "Confirm Password must equal the New Password.", Type: "eqfield"}}},
			http.StatusBadRequest,
			"Validation failed",
			rpdBadValidation,
			func(m *mocks.Repository, c *cache.Store) {
				m.On("ResetPassword", rpdBadValidation.Token, rpdBadValidation.NewPassword).Return(nil)
			},
		},
		"Cache Get Error": {
			nil,
			http.StatusBadRequest,
			"No user exists with the token: " + rp.Token,
			rp,
			func(m *mocks.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, rp.Token, &domain.User{}).Return(fmt.Errorf("error"))
			},
		},
		"Repo Error": {
			nil,
			http.StatusInternalServerError,
			"error",
			rp,
			func(m *mocks.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, rp.Token, &domain.User{}).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*domain.User)
					arg.Password = user.Password
					arg.Email = user.Email
				}).Return(nil)
				m.On("ResetPassword", user.Email, rp.NewPassword).Return(&errors.Error{Message: "error"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/reset", "/reset", test.input, func(ctx *gin.Context) {
				t.SetupCache(test.mock).ResetPassword(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
