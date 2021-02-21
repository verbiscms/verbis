// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/server"
)

// Load
//
// Loads all of the routes groups specified in the package
// And any global middleware to be used on the server.
func Load(d *deps.Deps, s *server.Server) {
	apiRoutes(d, s)
	frontendRoutes(d, s)
	spaRoutes(d, s)
}
