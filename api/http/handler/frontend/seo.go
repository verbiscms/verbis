package frontend

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/render"
	"github.com/gin-gonic/gin"
)

// SEOHandler defines methods for SEO routes to interact with the server
type SEOHandler interface {
	Robots(g *gin.Context)
	SiteMapIndex(g *gin.Context)
	SiteMapResource(g *gin.Context)
	SiteMapXSL(g *gin.Context, index bool)
}

// SEO defines the handler for all SEO Routes (sitemaps & robots)
type SEO struct {
	models  *models.Store
	config  config.Configuration
	sitemap render.SiteMapper
	options domain.Options
	render.ErrorHandler
}

// newSEO - Construct
func NewSEO(m *models.Store, config config.Configuration) *SEO {
	return &SEO{
		models:       m,
		config:       config,
		options:      m.Options.GetStruct(),
		sitemap:      render.NewSitemap(m),
		ErrorHandler: &render.Errors{},
	}
}

// Robots - Obtains the Seo Robots field from the Options struct
// which is set in the settings, and returns the robots.txt
// file.
//
// Returns a 404 if the options don't allow serving of robots.txt
func (c *SEO) Robots(g *gin.Context) {
	const op = "FrontendHandler.Robots"

	if !c.options.SeoRobotsServe {
		c.NotFound(g, c.config)
		return
	}

	g.Data(200, "text/plain", []byte(c.options.SeoRobots))
}

// SiteMapIndex obtains the sitemap index file from the sitemap
// model Obtains the []bytes to send back as data when
// /sitemap.xml is visited.
//
// Returns a 404 if there was an error obtaining the XML file.
// or there was no resource items found.
func (c *SEO) SiteMapIndex(g *gin.Context) {
	const op = "FrontendHandler.SiteMapIndex"

	sitemap, err := c.sitemap.GetIndex()
	if err != nil {
		c.NotFound(g, c.config)
	}

	g.Data(200, "application/xml; charset=utf-8", sitemap)
}

// SiteMapResource obtains the sitemap pages from the sitemap model
// by using the resource in the URL. Obtains the []bytes to send
// back as data when /:resource/sitemap.xml is visited.
//
// Returns a 404 if there was an error obtaining the XML file.
// or there was no resource items found.
func (c *SEO) SiteMapResource(g *gin.Context) {
	const op = "FrontendHandler.SiteMap"

	sitemap, err := c.sitemap.GetPages(g.Param("resource"))
	if err != nil {
		c.NotFound(g, c.config)
	}

	g.Data(200, "application/xml; charset=utf-8", sitemap)
}

// SiteMapXSL - Serves the XSL files for use with any .xml file that
// is used to serve the sitemap.
//
// Returns a 404 if there was an error obtaining the XSL.
func (c *SEO) SiteMapXSL(g *gin.Context, index bool) {
	const op = "FrontendHandler.SiteMapIndexXSL"

	sitemap, err := c.sitemap.GetXSL(index)
	if err != nil {
		c.NotFound(g, c.config)
	}

	g.Data(200, "application/xml; charset=utf-8", sitemap)
}
