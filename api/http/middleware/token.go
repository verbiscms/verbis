// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"github.com/verbiscms/verbis/api/store/users"
	"net/http"
)

// TokenCheck - TODO Comments
func TokenCheck(u users.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		// Check if the user token is valid
		user, err := u.FindByToken(token)
		if err != nil {
			api.AbortJSON(ctx, http.StatusUnauthorized, "Invalid token in the request header", nil)
			return
		}

		// Check if the session is expired or there was
		// an error obtaining the session data.
		err = u.CheckSession(token)
		if err == users.ErrSessionExpired {
			ctx.SetCookie("verbis-session", "", -1, "/", "", false, true)
			api.AbortJSON(ctx, http.StatusUnauthorized, "Session expired, please login again", gin.H{
				"errors": gin.H{
					"session": "expired",
				},
			})
			return
		} else if err != nil {
			api.AbortJSON(ctx, http.StatusUnauthorized, "Invalid token in the request header", nil)
			return
		}

		// Check if the user is banned, no routes can
		// be accessed.
		if user.Role.ID == domain.BannedRoleID {
			api.AbortJSON(ctx, http.StatusForbidden, "Your account has been suspended", nil)
			return
		}

		// Bind the user to Gin's Context for processing
		// down the request chain.
		ctx.Set("user", user)
	}
}
