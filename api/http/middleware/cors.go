package middleware

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(g *gin.Context) {
		if api.SuperAdmin {
			g.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8090")
		} else {
			g.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		g.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		g.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token")
		g.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if g.Request.Method == "OPTIONS" {
			g.AbortWithStatus(200)
			return
		}

		g.Next()
	}
}

