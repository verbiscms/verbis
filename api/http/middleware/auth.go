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

// ContextUser defines the key of obtaining the user
// from the context.
const ContextUser = "user"

// Authorise defines permissions middleware and checks to
// see if the user can access the group and method.
// If the user could not be obtained or cast to a
// domain.User, or the user does not have access
// to the endpoint http.StatusForbidden will be
// returned.
func Authorise(group, method string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtain the user from the context, if the
		// user is nil or there is no match, return.
		u, ok := ctx.Get(ContextUser)
		if !ok || u == nil {
			api.AbortJSON(ctx, http.StatusForbidden, "User not found", nil)
			return
		}

		// Try and convert the user to of type domain.User
		user, ok := u.(domain.User)
		if !ok {
			api.AbortJSON(ctx, http.StatusForbidden, "Error converting to type user", nil)
			return
		}

		// Check if the user has access to the group and method.
		// If they don't, abort and return.
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
