// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"github.com/verbiscms/verbis/api/store/users"
	"net/http"
)

// Administrator middleware
func AdminTokenCheck(d *deps.Deps) gin.HandlerFunc {
	return func(g *gin.Context) {
		if err := checkTokenExists(g); err != nil {
			return
		}

		u, err := checkUserToken(d, g)
		if err != nil {
			return
		}

		if u.Role.ID > 1 {
			g.Next()
		} else {
			api.AbortJSON(g, http.StatusForbidden, "You must have access level of administrator to access this endpoint.", nil)
			return
		}
	}
}

// Operator middleware
func OperatorTokenCheck(d *deps.Deps) gin.HandlerFunc {
	return func(g *gin.Context) {
		if err := checkTokenExists(g); err != nil {
			return
		}

		u, err := checkUserToken(d, g)
		if err != nil {
			return
		}

		if u.Role.ID > 0 {
			g.Next()
		} else {
			api.AbortJSON(g, http.StatusForbidden, "You must have access level of operator to access this endpoint.", nil)
			return
		}
	}
}

// Check if the token exists in the header
func checkTokenExists(g *gin.Context) error {
	token := g.Request.Header.Get("token")
	if token == "" {
		api.AbortJSON(g, http.StatusUnauthorized, "Missing token in the request header", nil)
		return fmt.Errorf("missing token")
	}
	return nil
}

// Check the user token and return the user if passes
func checkUserToken(d *deps.Deps, g *gin.Context) (*domain.User, error) {
	token := g.Request.Header.Get("token")

	u, err := d.Store.User.FindByToken(token)
	if err != nil {
		api.AbortJSON(g, http.StatusUnauthorized, "Invalid token in the request header", nil)
		return nil, err
	}

	err = d.Store.User.CheckSession(token)

	if err == users.ErrSessionExpired {
		g.SetCookie("verbis-session", "", -1, "/", "", false, true)
		api.AbortJSON(g, http.StatusUnauthorized, "Session expired, please login again", gin.H{
			"errors": gin.H{
				"session": "expired",
			},
		})
		return &domain.User{}, err
	}

	if err != nil {
		api.AbortJSON(g, http.StatusUnauthorized, "Invalid token in the request header", nil)
		return &domain.User{}, err
	}

	if u.Role.ID == domain.BannedRoleID {
		api.AbortJSON(g, http.StatusForbidden, "Your account has been suspended by the administration team", nil)
		return &domain.User{}, err
	}

	return &u, nil
}
