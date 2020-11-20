package controllers

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
)

// SiteHandler defines methods for the Site to interact with the server
type SiteHandler interface {
	GetSite(g *gin.Context)
	GetTheme(g *gin.Context)
	GetTemplates(g *gin.Context)
	GetLayouts(g *gin.Context)
}

// SiteController defines the handler for Posts
type SiteController struct {
	store *models.Store
	config    config.Configuration
}

// newSite - Construct
func newSite(m *models.Store, config config.Configuration) *SiteController {
	return &SiteController{
		store: m,
		config:    config,
	}
}

// GetSite gets site's general config
//
// Returns 200 if site config was obtained successfully.
func (c *SiteController) GetSite(g *gin.Context) {
	Respond(g, 200, "Successfully obtained site config", c.store.Site.GetGlobalConfig())
}

// GetTheme gets the theme's config from the theme path
//
// Returns 200 if theme config was obtained successfully.
// Returns 500 if there was an error getting the theme config.
func (c *SiteController) GetTheme(g *gin.Context) {
	const op = "SiteHandler.GetTheme"
	config, err := c.store.Site.GetThemeConfig()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}
	Respond(g, 200, "Successfully obtained theme config", config)
}

// GetTemplates gets all page templates
//
// Returns 200 if the templates were obtained successfully.
// Returns 500 if there was an error getting the templates.
func (c *SiteController) GetTemplates(g *gin.Context) {
	const op = "SiteHandler.GetTemplates"
	templates, err := c.store.Site.GetTemplates()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}
	Respond(g, 200, "Successfully obtained templates", templates)
}

// GetTemplates gets all layouts
//
// Returns 200 if the layouts were obtained successfully.
// Returns 500 if there was an error getting the layouts.
func (c *SiteController) GetLayouts(g *gin.Context) {
	const op = "SiteHandler.GetLayouts"
	templates, err := c.store.Site.GetLayouts()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}
	Respond(g, 200, "Successfully obtained layouts", templates)
}
