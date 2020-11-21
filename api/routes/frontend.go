package routes

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/http/controllers"
	"github.com/ainsleyclark/verbis/api/http/middleware"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/gin-gonic/gin"
)

func frontend(s *server.Server, c *controllers.Controller, m *models.Store, config config.Configuration) {

	// Set Frontend Middleware
	s.Use(middleware.Recovery(server.Recover))
	s.Use(middleware.Redirects(m.Options))
	s.Use(middleware.FrontEndCache(m.Options))

	_ = s.Group("")
	{
		// Serve assets
		s.GET("/assets/*any", c.Frontend.GetAssets)

		// Serve Verbis Assets
		s.Static("/verbis", paths.Api()+"/web/public")

		// Serve uploads
		s.GET("/uploads/*any", c.Frontend.GetUploads)

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
