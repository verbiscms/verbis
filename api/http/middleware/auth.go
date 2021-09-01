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

const ContextUser = "user"

// Authorise - TODO Comments
func Authorise(group string, method string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, ok := ctx.Get(ContextUser)
		if !ok || u == nil {
			api.AbortJSON(ctx, http.StatusForbidden, "User not found", nil)
			return
		}

		user, ok := u.(domain.User)
		if !ok {
			api.AbortJSON(ctx, http.StatusForbidden, "Error converting to type user", nil)
			return
		}

		err := user.Role.Permissions.Enforce(group, method)
		if err == domain.ErrPermissionDenied {
			api.AbortJSON(ctx, http.StatusForbidden, fmt.Sprintf("Forbidden, you do not have access to the %s, group with the method %s", group, method), nil)
			return
		} else if err != nil {
			api.AbortJSON(ctx, http.StatusBadRequest, "Invalid request", nil)
			return
		}

		ctx.Next()
	}
}
