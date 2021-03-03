// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SessionCheck
//
// Obtains the token header and checks if the session is
// still active. If it isn't the request will be
// aborted with a status code of http.StatusUnauthorized.
func SessionCheck(d *deps.Deps) gin.HandlerFunc {
	return func(g *gin.Context) {
		token := g.Request.Header.Get("token")

		err := d.Store.User.CheckSession(token)
		if err != nil {
			g.SetCookie("verbis-session", "", -1, "/", "", false, true)
			api.AbortJSON(g, http.StatusUnauthorized, "Session expired, please login again", gin.H{
				"errors": gin.H{
					"session": "expired",
				},
			})
			return
		}

		g.Next()
	}
}
