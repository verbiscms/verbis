package routes

import (
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
)

func api(s *server.Server, c *controllers.Controller, m *models.Store) {

	// Auth routes outside of admin
	s.GET("/email/verify/:token", c.Auth.VerifyEmail)
	s.GET("/password/verify/:token", c.Auth.VerifyPasswordToken)

	// API Routes
	api := s.Group("/api/v1")
	{
		// Site
		api.GET("/site", c.Site.GetSite)

		// Auth
		api.POST("/login", c.Auth.Login)
		api.POST("/logout", c.Auth.Logout)
		api.POST("/password/reset", c.Auth.ResetPassword)
		api.POST("/password/email", c.Auth.SendResetPassword)
		// TODO: Use gin and not vue for reset password
		api.GET("/password/email", c.Auth.SendResetPassword)

		// Operator
		operator := api.Group("")
		{
			operator.Use(middleware.SessionCheck(m.Session))
			operator.Use(middleware.OperatorTokenCheck(m.User, m.Session))
			operator.Use(middleware.EmptyBody())

			// Resources
			operator.GET("/resources", c.Site.GetResources)

			// Templates
			operator.GET("/templates", c.Site.GetTemplates)

			// Posts
			operator.GET("/posts", c.Posts.Get)
			operator.GET("/posts/:id", c.Posts.GetById)
			operator.POST("/posts", c.Posts.Create)
			operator.PUT("/posts/:id", c.Posts.Update)
			operator.DELETE("/posts/:id", c.Posts.Delete)

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
		}

		// Administrator
		admin := api.Group("")
		{
			admin.Use(middleware.AdminTokenCheck(m.User, m.Session))

			// Users
			admin.GET("/users", c.User.Get)
			admin.GET("/users/:id", c.User.GetById)
			admin.POST("/users", c.User.Create)
			admin.PUT("/users/:id", c.User.Update)
			admin.DELETE("/users/:id", c.User.Delete)
		}
	}
}
