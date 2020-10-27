package routes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/server"
)

// Vue (SPA) routes
func spa(s *server.Server, c *controllers.Controller, config config.Configuration) {
	spa := s.Group(config.Admin.Path)
	{
		spa.GET("/*any",  c.Spa.Serve)
		spa.GET("", c.Spa.Serve)
	}
}