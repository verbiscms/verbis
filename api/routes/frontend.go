package routes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/server"
)

func frontend(s *server.Server, c *controllers.Controller) {

	s.Use(middleware.Recovery(c.Frontend.Recovery))

	_ = s.Group("")
	{
		// Serve assets
		s.Static("/assets", paths.Theme() + config.Theme.AssetsPath)

		//s.GET("/verbis/images/verbis-logo.svg", c.Site.GetLogo)
		s.Static("/verbis", paths.Api() + "/web")

		s.GET("/", c.Frontend.Home)
		s.GET("/style-guide", c.Frontend.StyleGuide)
		s.POST("/ajax/subscribe", c.Frontend.Subscribe)
		s.GET("/uploads/*any", c.Frontend.GetUploads)

		s.NoRoute(c.Frontend.Serve)
	}
}