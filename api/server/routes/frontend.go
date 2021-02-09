// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/recovery"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func frontend(d *deps.Deps, s *server.Server, c *handler.Handler) {

	s.Use(recovery.New(d).New(http.StatusInternalServerError).HttpRecovery())
	s.Use(middleware.Redirects(d.Options))

	// TODO: This check should be in config
	uploadPath := d.Config.Media.UploadPath
	if uploadPath == "" {
		uploadPath = "uploads"
	}

	_ = s.Group("")
	{
		// Serve assets
		s.GET("/assets/*any", c.Frontend.GetAssets)

		// Serve Verbis Assets
		s.Static("/verbis", paths.Api()+"/web/public")

		// Serve uploads
		s.GET("/"+uploadPath+"/*any", c.Frontend.GetUploads)

		// Robots
		s.GET("/robots.txt", c.SEO.Robots)

		// Sitemap
		s.GET("/sitemap.xml", c.SEO.SiteMapIndex)
		s.GET("/sitemaps/:resource/:map", c.SEO.SiteMapResource)
		s.GET("/resource-sitemap.xsl", func(g *gin.Context) {
			c.SEO.SiteMapXSL(g, true)
		})
		s.GET("/main-sitemap.xsl", func(g *gin.Context) {
			c.SEO.SiteMapXSL(g, false)
		})

		// Favicon
		s.StaticFile("/favicon.ico", paths.Theme()+"/favicon.ico")

		// Serve the front end
		s.NoRoute(c.Frontend.Serve)
	}
}
