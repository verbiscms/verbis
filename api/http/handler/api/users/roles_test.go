// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
)

func (t *UsersTestSuite) TestUser_Roles() {
	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		mock    func(m *mocks.UserRepository)
	}{
		"Success": {
			roles,
			200,
			"Successfully obtained user roles",
			func(m *mocks.UserRepository) {
				m.On("GetRoles").Return(roles, nil)
			},
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			func(m *mocks.UserRepository) {
				m.On("GetRoles").Return(nil, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe("GET", "/roles", "/roles", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).Roles(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
