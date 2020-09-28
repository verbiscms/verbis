package controllers

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
)

// SiteHandler defines methods for the Site to interact with the server
type SiteHandler interface {
	GetSite(g *gin.Context)
	GetResources(g *gin.Context)
	GetTemplates(g *gin.Context)
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

// GetResources gets all resources
func (c *SiteController) GetResources(g *gin.Context) {
	const op = "SiteHandler.GetResources"
	resources, err := c.model.GetAllResources()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}
	Respond(g, 200,"Successfully obtained resources", resources)
}

// GetTemplates gets all templates
func (c *SiteController) GetTemplates(g *gin.Context) {
	const op = "SiteHandler.GetAllTemplates"
	templates, err := c.model.GetAllTemplates()
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}
	Respond(g, 200,"Successfully obtained templates", templates)
}