// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *UsersTestSuite) TestUser_ResetPassword() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(u *mocks.UserRepository)
		url     string
	}{
		"Success": {
			nil,
			200,
			"Successfully updated password for the user with ID: 123",
			reset,
			func(m *mocks.UserRepository) {
				m.On("GetByID", 123).Return(domain.User{}, nil)
				m.On("ResetPassword", 123, reset).Return(nil)
			},
			"/users/reset/123",
		},
		"Invalid ID": {
			nil,
			400,
			"A valid ID is required to update a user's password",
			reset,
			func(m *mocks.UserRepository) {
				m.On("GetByID", 123).Return(domain.User{}, nil)
				m.On("ResetPassword", 123, reset).Return(nil)
			},
			"/users/reset/wrongid",
		},
		"Not found": {
			nil,
			400,
			"No user has been found with the ID: 123",
			reset,
			func(m *mocks.UserRepository) {
				m.On("GetByID", 123).Return(domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
				m.On("ResetPassword", 123, reset).Return(nil)
			},
			"/users/reset/123",
		},
		"Validation Failed": {
			`{"errors":[{"key":"confirm_password", "message":"Confirm Password must equal the New Password.", "type":"eqfield"}]}`,
			400,
			"Validation failed",
			resetBadValidation,
			func(m *mocks.UserRepository) {
				m.On("GetByID", 123).Return(domain.User{}, nil)
				m.On("ResetPassword", 123, reset).Return(nil)
			},
			"/users/reset/123",
		},
		"Invalid": {
			nil,
			400,
			"invalid",
			reset,
			func(m *mocks.UserRepository) {
				m.On("GetByID", 123).Return(domain.User{}, nil)
				m.On("ResetPassword", 123, reset).Return(&errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
			"/users/reset/123",
		},
		"Internal": {
			nil,
			500,
			"internal",
			reset,
			func(m *mocks.UserRepository) {
				m.On("GetByID", 123).Return(domain.User{}, nil)
				m.On("ResetPassword", 123, reset).Return(&errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/users/reset/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, test.url, "/users/reset/:id", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).ResetPassword(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
