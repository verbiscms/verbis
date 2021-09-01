// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	app "github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/deps"
	"strings"
)

// Redirects
//
// Determines if the path includes the admin path or the
// api routes. If it doesn't, the redirect repository
// will be called with the path, if there is a
// match, the user will be redirected with the
// appropriate code.
func Redirects(d *deps.Deps) gin.HandlerFunc {
	return func(g *gin.Context) {
		path := g.Request.URL.String()

		if strings.Contains(path, app.AdminPath) || strings.Contains(path, app.HTTPAPIRoute) {
			g.Next()
			return
		}

		redirect, err := d.Store.Redirects.FindByFrom(path)
		if err != nil {
			g.Next()
			return
		}

		g.Redirect(redirect.Code, redirect.To)
	}
}
