// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

// Authorise - TODO Comments
func Authorise(group string, method string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//user, err := ctx.Get("user")
		//if err != nil {
		//	api.AbortJSON(ctx, http.StatusForbidden, "Your account has been suspended", nil)
		//	return
		//}

		user := domain.User{
			UserPart: domain.UserPart{
				Role: domain.Role{
					ID: domain.AuthorRoleID,
				},
			},
		}

		permission, exists := domain.Permissions[user.Role.ID][group][method]
		if !exists {
			api.AbortJSON(ctx, http.StatusBadRequest, "Invalid request", nil)
			return
		}

		if !permission.Allow {
			api.AbortJSON(ctx, http.StatusForbidden, fmt.Sprintf("Forbidden, you do not have access to the %s, group with the method %s", group, method), nil)
			return
		}

		fmt.Println("carry on son")
	}
}
