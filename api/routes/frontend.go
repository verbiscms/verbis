package routes

import (
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
)

func frontend(s *server.Server, c *controllers.Controller) {

	s.Use(middleware.Recovery(server.Recover))

	_ = s.Group("")
	{
		// Serve assets
		s.Static("/assets", paths.Theme() +	models.ThemeConfig.AssetsPath)

		// Serve Verbis Assets
		s.Static("/verbis", paths.Api() + "/web/public")

		s.GET("/uploads/*any", c.Frontend.GetUploads)

		s.NoRoute(c.Frontend.Serve)
	}
}