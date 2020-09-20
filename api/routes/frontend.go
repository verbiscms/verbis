package routes

import (
	"cms/api/config"
	"cms/api/helpers/paths"
	"cms/api/http/controllers"
	"cms/api/server"
)

func frontend(s *server.Server, c *controllers.Controller) {

	// Serve assets
	s.Static("/assets", paths.Theme() + config.Theme.AssetsPath)

	//s.GET("/verbis/images/verbis-logo.svg", c.Site.GetLogo)
	s.Static("/verbis", paths.Api() + "/static")

	s.GET("/", c.Frontend.Home)
	s.GET("/style-guide", c.Frontend.StyleGuide)
	s.POST("/ajax/subscribe", c.Frontend.Subscribe)
	s.GET("/uploads/*any", c.Frontend.GetUploads)

	s.NoRoute(c.Frontend.Serve)
}