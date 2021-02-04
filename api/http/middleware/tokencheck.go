// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
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

		if u.Role.Id > 1 {
			g.Next()
		} else {
			api.AbortJSON(g, 403, "You must have access level of administrator to access this endpoint.", nil)
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

		if u.Role.Id > 0 {
			g.Next()
		} else {
			api.AbortJSON(g, 403, "You must have access level of operator to access this endpoint.", nil)
			return
		}
	}
}

// Check if the token exists in the header
func checkTokenExists(g *gin.Context) error {
	token := g.Request.Header.Get("token")
	if token == "" {
		api.AbortJSON(g, 401, "Missing token in the request header", nil)
		return fmt.Errorf("Missing token")
	}
	return nil
}

// Check the user token and return the user if passes
func checkUserToken(d *deps.Deps, g *gin.Context) (*domain.User, error) {
	token := g.Request.Header.Get("token")

	u, err := d.Store.User.CheckToken(token)
	if err != nil {
		api.AbortJSON(g, 401, "Invalid token in the request header", nil)
		return &domain.User{}, err
	}

	if u.Role.Id == 0 {
		api.AbortJSON(g, 403, "Your account has been suspended by the administration team", nil)
		return &domain.User{}, err
	}

	return &u, nil
}

// Check to see if the session has expired
func checkSession(g *gin.Context, userId int) error {

	//if hasSession := m.Has(userId); !hasSession {
	//	return nil
	//}

	//err := m.Check(userId);
	//if err != nil {
	//	controllers.AbortJSON(g, 401, errors.Message(err), err)
	//	return err
	//}

	return nil
}
