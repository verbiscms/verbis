// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	app "github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/gin-gonic/gin"
	"strings"
)

// Redirects
//
// Determines if the path includes the admin path or the
// api routes. If it doesnt, the redirect repository
// will be called with the path, if there is a
// match, the user will be redirect with the
// appropriate code.
func Redirects(d *deps.Deps) gin.HandlerFunc {
	return func(g *gin.Context) {
		path := g.Request.URL.String()

		if strings.Contains(path, d.Config.Admin.Path) || strings.Contains(path, app.APIRoute) {
			g.Next()
			return
		}

		redirect, err := d.Store.Redirects.GetByFrom(path)
		if err != nil {
			g.Next()
			return
		}

		g.Redirect(redirect.Code, redirect.To)
	}
}
