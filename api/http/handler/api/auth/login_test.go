// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (t *AuthTestSuite) TestAuth_Login() {

	var (
		login              = Login{Email: "info@verbiscms.com", Password: "password"}
		loginBadValidation = Login{Password: "password"}
	)

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		cookie  bool
		mock    func(m *mocks.AuthRepository)
	}{
		"Success": {
			user,
			200,
			"Successfully logged in & session started",
			login,
			true,
			func(m *mocks.AuthRepository) {
				m.On("Authenticate", login.Email, login.Password).Return(user, nil)
			},
		},
		"Validation Failed": {
			api.ErrorJson{Errors: validation.Errors{{Key: "email", Message: "Email is required.", Type: "required"}}},
			400,
			"Validation failed",
			loginBadValidation,
			false,
			func(m *mocks.AuthRepository) {
				m.On("Authenticate", loginBadValidation.Email, loginBadValidation.Email).Return(domain.User{}, fmt.Errorf("error"))
			},
		},
		"Not Authorised": {
			nil,
			401,
			"unauthorised",
			login,
			false,
			func(m *mocks.AuthRepository) {
				m.On("Authenticate", login.Email, login.Password).Return(domain.User{}, &errors.Error{Code: errors.INVALID, Message: "unauthorised"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/login", "/login", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Login(ctx)
			})

			if test.cookie {
				cookie := http.Cookie{
					Name:     "verbis-session",
					Expires:  time.Time{},
					MaxAge:   172800,
					Path:     "/",
					Raw:      "verbis-session=; Path=/; Max-Age=172800; HttpOnly",
					HttpOnly: true,
				}
				t.Equal(t.Recorder.Result().Cookies()[0], &cookie)
			}

			t.RunT(test.want, test.status, test.message)
		})
	}
}
