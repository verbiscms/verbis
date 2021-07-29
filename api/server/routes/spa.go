// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	app "github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/http/handler"
	"github.com/verbiscms/verbis/api/server"
)

// spaRoutes
//
// Vue (SPA) routes.
func spaRoutes(d *deps.Deps, s *server.Server) {
	h := handler.NewSPA(d)
	spa := s.Group(app.AdminPath)
	spa.GET("/*any", h.Serve)
	spa.GET("", h.Serve)
}
