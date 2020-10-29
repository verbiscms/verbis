package routes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
)

func frontend(s *server.Server, c *controllers.Controller, m *models.Store, config config.Configuration) {

	// Set Frontend Middleware
	s.Use(middleware.Recovery(server.Recover))
	s.Use(middleware.FrontEndCache(m.Options))

	_ = s.Group("")
	{
		// Serve assets
		s.GET("/assets/*any", c.Frontend.GetAssets)

		// Serve Verbis Assets
		s.Static("/verbis", paths.Api() + "/web/public")

		// Serve uploads
		s.GET("/uploads/*any", c.Frontend.GetUploads)

		// Serve the front end
		s.NoRoute(c.Frontend.Serve)
	}
}