package routes

import (
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
)

// Load all of the routes groups specified in the package
// And any global middleware to be used on the server.
func Load(s *server.Server, c *controllers.Controller, m *models.Store) {

	// Global middleware
	s.Use(middleware.Log())

	// Load routes
	api(s, c, m)
	frontend(s, c)
	spa(s, c)
}
