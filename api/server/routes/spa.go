// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/server"
)

// spaRoutes
//
// Vue (SPA) routes.
func spaRoutes(d *deps.Deps, s *server.Server) {
	h := handler.NewSPA(d)

	//fsys, err := fs.Sub(admin.SPA, verbisfs.SpaDistFolder)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//system := http.FS(fsys)
	//
	//s.Use(sta)
	//
	//s.StaticFS("/admin/*", system)
	//
	////spa := s.Group("/admin")
	////spa.GET("/*any", test)
	////spa.GET("", test)

	spa := s.Group("/admin")
	spa.GET("/*any", h.Serve)
	spa.GET("", h.Serve)
}
