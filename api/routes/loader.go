package routes

import (
	"cms/api/http/controllers"
	"cms/api/http/middleware"
	"cms/api/models"
	"cms/api/server"
)

// Load all of the routes groups specified in the package
// And any global middleware to be used on the server.
func Load(s *server.Server, c *controllers.Controller, m *models.Store) {

	// Global middleware
	s.Use(middleware.CORSMiddleware())

	// Load routes
	api(s, c, m)
	frontend(s, c)
	spa(s, c)
}
