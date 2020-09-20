package middleware

import (
	"cms/api/models"
	"github.com/gin-gonic/gin"
)

func SessionCheck(m models.SessionRepository) gin.HandlerFunc {
	return func(g *gin.Context) {

		// Get the Verbis session cookie
		cookie, err := g.Cookie("verbis-session")
		if err != nil {
			g.Next()
		}

		// Update the session table
		if err := m.Update(cookie); err != nil {
			g.Next()
		}

		g.Next()
	}
}