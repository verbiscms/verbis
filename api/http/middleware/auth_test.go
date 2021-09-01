// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

func (t *MiddlewareTestSuite) TestAuthorise() {
	tt := map[string]struct {
		group  string
		method string
		user   interface{}
		status int
		want   string
	}{
		"No User": {
			domain.PermissionSettings,
			domain.ViewMethod,
			nil,
			http.StatusForbidden,
			"User not found",
		},
		"Bad Cast": {
			domain.PermissionSettings,
			domain.ViewMethod,
			10,
			http.StatusForbidden,
			"Error converting to type user",
		},
		"Permission Denied": {
			domain.PermissionSettings,
			domain.ViewMethod,
			domain.User{
				UserPart: domain.UserPart{
					Role: domain.Role{
						Permissions: domain.RbacGroup{domain.PermissionSettings: {
							domain.ViewMethod: {Allow: false},
						}},
					},
				},
			},
			http.StatusForbidden,
			"Forbidden, you do not have access",
		},
		"Invalid Request": {
			domain.PermissionSettings,
			domain.ViewMethod,
			domain.User{
				UserPart: domain.UserPart{
					Role: domain.Role{
						Permissions: domain.RbacGroup{"wrong": {
							domain.ViewMethod: {Allow: false},
						}},
					},
				},
			},
			http.StatusBadRequest,
			"Invalid request",
		},
		"Permitted": {
			domain.PermissionSettings,
			domain.ViewMethod,
			domain.User{
				UserPart: domain.UserPart{
					Role: domain.Role{
						Permissions: domain.RbacGroup{domain.PermissionSettings: {
							domain.ViewMethod: {Allow: true},
						}},
					},
				},
			},
			http.StatusOK,
			"success",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			defer t.Reset()

			t.Engine.GET("/test",
				func(ctx *gin.Context) { ctx.Set(ContextUser, test.user) },
				Authorise(test.group, test.method),
				func(ctx *gin.Context) { api.Respond(ctx, http.StatusOK, "success", nil) },
			)

			t.NewRequest(http.MethodGet, "/test", nil)
			t.ServeHTTP()

			respond, _ := t.RespondData()
			t.Equal(test.status, t.Status())
			t.Contains(respond.Message, test.want)
		})
	}
}
