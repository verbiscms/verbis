package controllers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// OptionsHandler defines methods for Options to interact with the server
type OptionsHandler interface {
	Get(g *gin.Context)
	GetByName(g *gin.Context)
	UpdateCreate(g *gin.Context)
}

// OptionsController defines the handler for Options
type OptionsController struct {
	controller Controller
	model      models.OptionsRepository
	server     *server.Server
}

// newOptions - Construct
func newOptions(m models.OptionsRepository) *OptionsController {
	return &OptionsController{
		model: m,
	}
}

// Get All
func (c *OptionsController) Get(g *gin.Context) {
	const op = "OptionsHandler.Delete"

	options, err := c.model.Get()
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	}
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained options", options)
}

// Get By name
// Returns errors.INVALID if the name was not passed.
func (c *OptionsController) GetByName(g *gin.Context) {
	const op = "OptionsHandler.GetByName"

	name := g.Param("name")
	if name == "" {
		Respond(g, 400,  "A name is required to obtain the option by name", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("no name passed"), Operation: op})
		return
	}

	option, err := c.model.GetByName(name)
	if err != nil {
		Respond(g, 400, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained option with name: " + name, option)
}

// Update & Create options
// Returns errors.INVALID if validation failed.
func (c *OptionsController) UpdateCreate(g *gin.Context) {
	const op = "OptionsHandler.UpdateCreate"

	var vOptions domain.Options
	if err := g.ShouldBindBodyWith(&vOptions, binding.JSON); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	var options domain.OptionsDB
	if err := g.ShouldBindBodyWith(&options, binding.JSON); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if err := c.model.UpdateCreate(options); err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully created/updated options", nil)
}


