package controllers

import (
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
)

type SiteController struct {
	model models.SiteRepository
}

type SiteHandler interface {
	GetSite(g *gin.Context)
	GetResources(g *gin.Context)
	GetTemplates(g *gin.Context)
}

// Construct
func newSite(m models.SiteRepository) *SiteController {
	r := &SiteController{
		model: m,
	}
	return r
}

// Get site's general config
func (c *SiteController) GetSite(g *gin.Context) {
	Respond(g, 200, "Successfully obtained site config", c.model.GetGlobalConfig())
}

// Get all resources
func (c *SiteController) GetResources(g *gin.Context) {
	resources, err := c.model.GetAllResources()
	if err != nil {
		Respond(g, 500, err.Error(), nil)
		return
	}
	Respond(g, 200,"Successfully obtained resources", resources)
}

// Get all resources
func (c *SiteController) GetTemplates(g *gin.Context) {
	templates, err := c.model.GetAllTemplates()
	if err != nil {
		Respond(g, 500, err.Error(), nil)
		return
	}
	Respond(g, 200,"Successfully obtained templates", templates)
}