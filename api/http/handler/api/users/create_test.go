// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *UsersTestSuite) TestUser_Create() {

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.UserRepository)
	}{
		"Success": {
			user,
			200,
			"Successfully created user with ID: 123",
			userCreate,
			func(m *mocks.UserRepository) {
				m.On("Create", &userCreate).Return(user, nil)
			},
		},
		"Validation Failed": {
			api.ValidationErrJson{Errors: validation.Errors{{Key: "role_id", Message: "Role Id is required.", Type: "required"}}},
			400,
			"Validation failed",
			userCreateBadValidation,
			func(m *mocks.UserRepository) {
				m.On("Create", &userCreateBadValidation).Return(domain.User{}, fmt.Errorf("error"))
			},
		},
		"Invalid": {
			nil,
			400,
			"invalid",
			userCreate,
			func(m *mocks.UserRepository) {
				m.On("Create", &userCreate).Return(domain.User{}, &errors.Error{Code: errors.INVALID, Message: "invalid"})
			},
		},
		"Conflict": {
			nil,
			400,
			"conflict",
			userCreate,
			func(m *mocks.UserRepository) {
				m.On("Create", &userCreate).Return(domain.User{}, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
			},
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			userCreate,
			func(m *mocks.UserRepository) {
				m.On("Create", &userCreate).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPost, "/users", "/users", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Create(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}