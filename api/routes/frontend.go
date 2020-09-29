package routes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/server"
)

func frontend(s *server.Server, c *controllers.Controller) {

	s.Use(middleware.Recovery(server.Recover))

	_ = s.Group("")
	{
		// Serve assets
		s.Static("/assets", paths.Theme() + config.Theme.AssetsPath)

		// Serve Verbis Assets
		s.Static("/verbis", paths.Api() + "/web/public")

		s.GET("/", c.Frontend.Home)
		s.GET("/test", c.Frontend.Test)
		s.GET("/uploads/*any", c.Frontend.GetUploads)

		s.NoRoute(c.Frontend.Serve)
	}
}