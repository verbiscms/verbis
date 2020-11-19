package controllers

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/frontend"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// SEOHandler defines methods for SEO routes to interact with the server
type SEOHandler interface {
	Robots(g *gin.Context)
	SiteMapIndex(g *gin.Context)
	SiteMapResource(g *gin.Context)
	SiteMapXSL(g *gin.Context, index bool)
}

// SEOController defines the handler for all SEO Routes (sitemaps & robots)
type SEOController struct {
	models  *models.Store
	config  config.Configuration
	sitemap frontend.SiteMapper
	options domain.Options
}

// newSEO - Construct
func newSEO(m *models.Store, config config.Configuration) *SEOController {
	const op = "SEOHandler.newSEO"

	options, err := m.Options.GetStruct()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: err},
		}).Fatal()
	}

	sitemap := frontend.NewSitemap(m)

	return &SEOController{
		models:  m,
		config: config,
		options: options,
		sitemap: sitemap,
	}
}

// Robots - Obtains the Seo Robots field from the Options struct
// which is set in the settings, and returns the robots.txt
// file.
func (c *SEOController) Robots(g *gin.Context) {
	const op = "FrontendHandler.Robots"

	if c.options.SeoRobotsServe {
		frontend.Error(g, c.config)
		return
	}

	g.Data(200, "text/plain", []byte(c.options.SeoRobots))
}

// SiteMap - Creates a new frontend.Sitemap instance and passes the
// store. GetPages obtains the []bytes to send back as xml
// when /sitemap.xml is visited.
func (c *SEOController) SiteMapIndex(g *gin.Context) {
	const op = "FrontendHandler.SiteMapIndex"

	sitemap, err := c.sitemap.GetIndex()
	if err != nil {
		frontend.Error(g, c.config)
	}

	g.Data(200, "application/xml; charset=utf-8", sitemap)
}

// SiteMap - Creates a new frontend.Sitemap instance and passes the
// store. GetPages obtains the []bytes to send back as xml
// when /sitemap.xml is visited.
func (c *SEOController) SiteMapResource(g *gin.Context) {
	const op = "FrontendHandler.SiteMap"

	sitemap, err := c.sitemap.GetPages(g.Param("resource"))
	if err != nil {
		frontend.Error(g, c.config)
	}

	g.Data(200, "application/xml; charset=utf-8", sitemap)
}

// SiteMap - Creates a new frontend.Sitemap instance and passes the
// store. GetPages obtains the []bytes to send back as xml
// when /sitemap.xml is visited.
func (c *SEOController) SiteMapXSL(g *gin.Context, index bool) {
	const op = "FrontendHandler.SiteMapIndexXSL"

	sitemap, err := c.sitemap.GetXSL(index)
	if err != nil {
		frontend.Error(g, c.config)
	}

	g.Data(200, "application/xml; charset=utf-8", sitemap)
}
