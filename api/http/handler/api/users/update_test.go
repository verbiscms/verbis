// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/verbiscms/verbis/api/common/vaidation"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/store/users"
	"net/http"
)

func (t *UsersTestSuite) TestUser_Update() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			user,
			http.StatusOK,
			"Successfully updated user with ID: 123",
			user,
			func(m *mocks.Repository) {
				m.On("Update", user).Return(user, nil)
			},
			"/users/123",
		},
		"Validation Failed": {
			api.ErrorJSON{Errors: validation.Errors{{Key: "role_id", Message: "Role Id is required.", Type: "required"}}},
			http.StatusBadRequest,
			"Validation failed",
			userBadValidation,
			func(m *mocks.Repository) {
				m.On("Update", userBadValidation).Return(domain.User{}, fmt.Errorf("error"))
			},
			"/users/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"A valid ID is required to update the user",
			user,
			func(m *mocks.Repository) {
				m.On("Update", userBadValidation).Return(domain.User{}, fmt.Errorf("error"))
			},
			"/users/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusBadRequest,
			"not found",
			user,
			func(m *mocks.Repository) {
				m.On("Update", user).Return(domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/users/123",
		},
		"Internal": {
			nil,
			http.StatusInternalServerError,
			"internal",
			user,
			func(m *mocks.Repository) {
				m.On("Update", user).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/users/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodPut, test.url, "/users/:id", test.input, func(ctx *gin.Context) {
				t.Setup(test.mock).Update(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
