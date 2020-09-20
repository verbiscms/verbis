package routes

import (
	"cms/api/config"
	"cms/api/http/controllers"
	"cms/api/server"
)

// Vue (SPA) routes
func spa(s *server.Server, c *controllers.Controller) {
	spa := s.Group(config.Admin.Path)
	{
		spa.GET("/*any",  c.Spa.Serve)
		spa.GET("", c.Spa.Serve)
	}
}