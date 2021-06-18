// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *UsersTestSuite) TestUser_Find() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			user,
			http.StatusOK,
			"Successfully obtained user with ID: 123",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(user, nil)
			},
			"/users/123",
		},
		"Invalid ID": {
			nil,
			http.StatusBadRequest,
			"Pass a valid number to obtain the user by ID",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(domain.User{}, fmt.Errorf("error"))
			},
			"/users/wrongid",
		},
		"Not Found": {
			nil,
			http.StatusOK,
			"no users found",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(domain.User{}, &errors.Error{Code: errors.NOTFOUND, Message: "no users found"})
			},
			"/users/123",
		},
		"Internal Error": {
			nil,
			http.StatusInternalServerError,
			"internal",
			func(m *mocks.Repository) {
				m.On("Find", 123).Return(domain.User{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
			"/users/123",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, test.url, "/users/:id", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Find(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
