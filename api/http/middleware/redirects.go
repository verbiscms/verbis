// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

func Redirects(options *domain.Options) gin.HandlerFunc {
	return func(g *gin.Context) {
		path := location.Get(g).String() + g.Request.URL.String()

		// when storing redirects, string lowering database
		// and string lowercase here.

		for _, v := range options.SeoRedirects {
			if path == v.From {
				g.Redirect(v.Code, v.To)
				return
			}
		}

		g.Next()
	}
}
