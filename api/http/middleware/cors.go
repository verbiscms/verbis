// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORS
//
// Sets up CORS headers for the API, if the request
// method is options, the request will be aborted
// with a status of 200.
func CORS() gin.HandlerFunc {
	return func(g *gin.Context) {
		g.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		g.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		g.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Origin, Cache-Control, X-Requested-With, token")
		g.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if g.Request.Method == "OPTIONS" {
			g.AbortWithStatus(200)
			return
		}

		g.Next()
	}
}
