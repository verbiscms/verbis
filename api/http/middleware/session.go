// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

func SessionCheck(d *deps.Deps) gin.HandlerFunc {
	return func(g *gin.Context) {
		token := g.Request.Header.Get("token")

		if err := d.Store.User.CheckSession(token); err != nil {
			g.SetCookie("verbis-session", "", -1, "/", "", false, true)
			api.AbortJSON(g, 401, "Session expired, please login again", gin.H{
				"errors": gin.H{
					"session": "expired",
				},
			})
			return
		}

		g.Next()
	}
}
