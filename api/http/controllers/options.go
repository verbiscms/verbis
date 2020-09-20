package controllers

import (
	"cms/api/domain"
	"cms/api/models"
	"cms/api/server"
	"github.com/gin-gonic/gin"
)

type OptionsController struct {
	controller Controller
	model      models.OptionsRepository
	server     *server.Server
}

type OptionsHandler interface {
	Get(g *gin.Context)
	GetByName(g *gin.Context)
	UpdateCreate(g *gin.Context)
}

// Construct
func newOptions(m models.OptionsRepository) *OptionsController {
	return &OptionsController{
		model: m,
	}
}

// Get All
func (c *OptionsController) Get(g *gin.Context) {
	options, err := c.model.GetAll()
	if err != nil {
		Respond(g, 500, err.Error(), nil)
		return
	}

	successMsg := "Successfully obtained options"
	if len(options) == 0 {
		successMsg = "No options available"
	}
	Respond(g, 200, successMsg, options)
}

// Get By name
func (c *OptionsController) GetByName(g *gin.Context) {
	name := g.Param("name")
	if name == "" {
		Respond(g, 500, "A name is required to obtain the option by name", nil)
		return
	}

	option, err := c.model.GetByName(name)
	if err != nil {
		Respond(g, 400, err.Error(), nil)
		return
	}

	Respond(g, 200, "Successfully obtained option with name: " + name, option)
}

// Update & Create options
func (c *OptionsController) UpdateCreate(g *gin.Context) {
	var options domain.Options
	if err := g.ShouldBindJSON(&options); err != nil {
		Respond(g, 400, "Validation failed", err)
		return
	}

	if err := c.model.UpdateCreate(options); err != nil {
		Respond(g, 500, err.Error(), nil)
	}

	Respond(g, 200, "Successfully created/updated options", nil)
}


