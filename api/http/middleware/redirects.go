// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/gin-gonic/gin"
)

func Redirects(d *deps.Deps) gin.HandlerFunc {
	return func(g *gin.Context) {
		path := g.Request.URL.String()
		redirect, err := d.Store.Redirects.GetByFrom(path)

		if err != nil {
			g.Next()
			return
		}

		g.Redirect(redirect.Code, redirect.To)
	}
}
