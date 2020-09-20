package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		g.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		g.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		g.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token")
		g.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if g.Request.Method == "OPTIONS" {
			g.AbortWithStatus(200)
			return
		}

		g.Next()
	}
}

