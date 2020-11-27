package routes

import (
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
)

func api(s *server.Server, c *handler.Handler, m *models.Store) {

	// Auth routes outside of admin

	// API Routes
	api := s.Group("/api/v1")
	{
		// Set API Middleware
		api.Use(middleware.EmptyBody())

		// Site
		api.GET("/site", c.Site.GetSite)

		// Auth
		api.POST("/login", c.Auth.Login)
		api.POST("/logout", c.Auth.Logout)
		api.POST("/password/reset", c.Auth.ResetPassword)
		api.POST("/password/email", c.Auth.SendResetPassword)
		api.GET("/password/verify/:token", c.Auth.VerifyPasswordToken)
		api.GET("/email/verify/:token", c.Auth.VerifyEmail)

		// Operator
		operator := api.Group("")
		{
			operator.Use(middleware.OperatorTokenCheck(m.User))
			operator.Use(middleware.SessionCheck(m.User))

			// Theme
			operator.GET("/theme", c.Site.GetTheme)

			// Templates
			operator.GET("/templates", c.Site.GetTemplates)

			// Layouts
			operator.GET("/layouts", c.Site.GetLayouts)

			// Posts
			operator.GET("/posts", c.Posts.Get)
			operator.GET("/posts/:id", c.Posts.GetById)
			operator.POST("/posts", c.Posts.Create)
			operator.PUT("/posts/:id", c.Posts.Update)
			operator.DELETE("/posts/:id", c.Posts.Delete)

			// Fields
			operator.GET("/fields", c.Fields.Get)

			// Categories
			operator.GET("/categories", c.Categories.Get)
			operator.GET("/categories/:id", c.Categories.GetById)
			operator.POST("/categories", c.Categories.Create)
			operator.PUT("/categories/:id", c.Categories.Update)
			operator.DELETE("/categories/:id", c.Categories.Delete)

			// Media
			operator.GET("/media", c.Media.Get)
			operator.GET("/media/:id", c.Media.GetById)
			operator.POST("/media", c.Media.Upload)
			operator.PUT("/media/:id", c.Media.Update)
			operator.DELETE("/media/:id", c.Media.Delete)

			// Options
			operator.GET("/options", c.Options.Get)
			operator.GET("/options/:name", c.Options.GetByName)
			operator.POST("/options", c.Options.UpdateCreate)

			// Users
			operator.GET("/users", c.User.Get)
			operator.GET("/users/:id", c.User.GetById)
			operator.PUT("/users/:id", c.User.Update)
			operator.POST("/users/:id/reset-password", c.User.ResetPassword)

			// Roles
			operator.GET("/roles", c.User.GetRoles)

			// Cache
			operator.POST("/cache", c.Cache.Clear)
		}

		// Administrator
		admin := api.Group("")
		{
			admin.Use(middleware.AdminTokenCheck(m.User))
			operator.Use(middleware.SessionCheck(m.User))

			// Users
			admin.POST("/users", c.User.Create)
			admin.DELETE("/users/:id", c.User.Delete)
		}
	}
}
