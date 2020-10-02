package controllers

import (
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
	model models.SiteRepository
}

// newSite - Construct
func newSite(m models.SiteRepository) *SiteController {
	return &SiteController{
		model: m,
	}
}

// GetSite gets site's general config
func (c *SiteController) GetSite(g *gin.Context) {
	Respond(g, 200, "Successfully obtained site config", c.model.GetGlobalConfig())
}

// GetTheme gets the theme's config from the theme path
func (c *SiteController) GetTheme(g *gin.Context) {
	const op = "SiteHandler.GetTheme"
	config, err := c.model.GetThemeConfig()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}
	Respond(g, 200,"Successfully obtained theme config", config)
}

// GetTemplates gets all templates
func (c *SiteController) GetTemplates(g *gin.Context) {
	const op = "SiteHandler.GetTemplates"
	templates, err := c.model.GetTemplates()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}
	Respond(g, 200,"Successfully obtained templates", templates)
}

// GetTemplates gets al layouts
func (c *SiteController) GetLayouts(g *gin.Context) {
	const op = "SiteHandler.GetLayouts"
	templates, err := c.model.GetLayouts()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}
	Respond(g, 200,"Successfully obtained layouts", templates)
}