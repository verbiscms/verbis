// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

// frontendRoutes
//
// Public facing routes.
func frontendRoutes(d *deps.Deps, s *server.Server) {
	h := handler.NewFrontend(d)

	// TODO: This check should be in config
	uploadPath := d.Config.Media.UploadPath
	if uploadPath == "" {
		uploadPath = "uploads"
	}

	frontend := s.Group("")
	frontend.Use(middleware.Redirects(d))

	// Serve assets
	s.GET("/assets/*any", h.Public.Assets)

	// Serve Verbis Assets
	//s.GET("/verbis/*any", d.FS.Web.HTTP("/verbis", "public"))
	fsys, err := fs.Sub(d.FS.Web, "public")
	if err != nil {
		logger.Fatal(&errors.Error{Code: errors.INTERNAL, Message: "Error creating sub FS", Operation: "Router.Frontend", Err: err})
	}

	s.StaticFS("/verbis", http.FS(fsys))

	// Serve uploads
	s.GET("/"+uploadPath+"/*any", h.Public.Uploads)

	s.GET(config.ScreenshotURL+":theme/:file", h.Public.Screenshot)

	// Robots
	s.GET("/robots.txt", h.SEO.Robots)

	// Sitemap
	s.GET("/sitemap.xml", h.SEO.SiteMapIndex)
	s.GET("/sitemaps/:resource/:map", h.SEO.SiteMapResource)
	s.GET("/resource-sitemap.xsl", func(g *gin.Context) {
		h.SEO.SiteMapXSL(g, true)
	})
	s.GET("/main-sitemap.xsl", func(g *gin.Context) {
		h.SEO.SiteMapXSL(g, false)
	})

	// Favicon
	s.StaticFile("/favicon.ico", d.ThemePath()+"/favicon.ico")

	// Serve the front end
	s.NoRoute(h.Public.Serve)
}
