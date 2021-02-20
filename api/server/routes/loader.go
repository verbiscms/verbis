// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/server"
)

// Load all of the routes groups specified in the package
// And any global middleware to be used on the server.
func Load(d *deps.Deps, s *server.Server, c *handler.Handler) {

	// Global middleware
	s.Use(middleware.Log())
	s.Use(middleware.CORS())

	// Load routes
	api(d, s)
	frontend(d, s)
	spa(d, s, c)
}
