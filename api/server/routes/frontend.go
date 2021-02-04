// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/http/handler"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/gin-gonic/gin"
)

func frontend(s *server.Server, c *handler.Handler, m *models.Store, config config.Configuration) {

	// Set Frontend Middleware
	s.Use(middleware.Recovery(server.Recover))
	s.Use(middleware.Redirects(m.Options))

	// TODO: This check should be in config
	uploadPath := config.Media.UploadPath
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
