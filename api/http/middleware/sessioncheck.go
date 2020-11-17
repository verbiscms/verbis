package middleware

import (
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
)

func SessionCheck(m models.UserRepository) gin.HandlerFunc {
	return func(g *gin.Context) {

		token := g.Request.Header.Get("token")

		if err := m.CheckSession(token); err != nil {
			g.SetCookie("verbis-session", "", -1, "/", "", false, true)
			controllers.AbortJSON(g, 401, "Session expired, please login again.", gin.H{
				"errors": gin.H{
					"session": "expired",
				},
			})
			return
		}

		g.Next()
	}
}
