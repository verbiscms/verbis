// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
)

// Load all of the routes groups specified in the package
// And any global middleware to be used on the server.
func Load(s *server.Server, c *handler.Handler, m *models.Store, cf config.Configuration) {

	// Global middleware
	s.Use(middleware.Log())
	s.Use(middleware.CORS())

	// Load routes
	api(s, c, m)
	frontend(s, c, m, cf)
	spa(s, c, cf)
}
