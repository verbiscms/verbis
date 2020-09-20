package middleware

import (
	"cms/api/domain"
	"cms/api/http/controllers"
	"cms/api/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Administrator middleware
func AdminTokenCheck(userModel models.UserRepository, sessionModel models.SessionRepository) gin.HandlerFunc {
	return func(g *gin.Context) {

		if err := checkTokenExists(g); err != nil {
			return
		}

		u, err := checkUserToken(g, userModel)
		if err != nil {
			return
		}

		if err := checkSession(g, u.Id, sessionModel); err != nil {
			return
		}

		if u.Role.Id > 1 {
			g.Next()
		} else {
			controllers.AbortJSON(g, 403, "You must have access level of administrator to access this endpoint.", nil)
			return
		}
	}
}

// Operator middleware
func OperatorTokenCheck(userModel models.UserRepository, sessionModel models.SessionRepository) gin.HandlerFunc {
	return func(g *gin.Context) {

		if err := checkTokenExists(g); err != nil {
			return
		}

		u, err := checkUserToken(g, userModel)
		if err != nil {
			return
		}

		if err := checkSession(g, u.Id, sessionModel); err != nil {
			return
		}

		if u.Role.Id > 0 {
			g.Next()
		} else {
			controllers.AbortJSON(g, 403, "You must have access level of operator to access this endpoint.", nil)
			return
		}
	}
}

// Check if the token exists in the header
func checkTokenExists(g *gin.Context) error  {
	token := g.Request.Header.Get("token")

	if token == "" {
		controllers.AbortJSON(g, 401, "Missing token in the request header", nil)
		return fmt.Errorf("Missing token")
	}

	return nil
}

// Check the user token and return the user if passes
func checkUserToken(g *gin.Context, m models.UserRepository) (*domain.User, error)  {
	token := g.Request.Header.Get("token")

	u, err := m.CheckToken(token)
	if err != nil {
		controllers.AbortJSON(g, 401,"Invalid token in the request header", nil)
		return &domain.User{}, err
	}

	if u.Role.Id == 0 {
		controllers.AbortJSON(g, 403, "Your account has been suspended by the administration team", nil)
		return &domain.User{}, err
	}

	return &u, nil
}

// Check to see if the session has expired
func checkSession(g *gin.Context, userId int, m models.SessionRepository) error {

	if hasSession := m.Has(userId); !hasSession {
		return nil
	}

	if err := m.Check(userId); err != nil {
		controllers.AbortJSON(g, 401, err.Error(), gin.H{
			"reason": "Timed out",
		})
		return err
	}

	return nil
}