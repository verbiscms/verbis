package controllers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/teamwork/reload"
	"time"
)

// OptionsHandler defines methods for Options to interact with the server
type OptionsHandler interface {
	Get(g *gin.Context)
	GetByName(g *gin.Context)
	UpdateCreate(g *gin.Context)
}

// OptionsController defines the handler for Options
type OptionsController struct {
	store  *models.Store
	config config.Configuration
}

// newOptions - Construct
func newOptions(m *models.Store, config config.Configuration) *OptionsController {
	return &OptionsController{
		store:  m,
		config: config,
	}
}

// Get All
//
// Returns 200 if there are no options or success.
// Returns 500 if there was an error getting the options.
func (c *OptionsController) Get(g *gin.Context) {
	const op = "OptionsHandler.Delete"

	options, err := c.store.Options.Get()
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
//
// Returns 200 if there are no options or success.
// Returns 400 if there was name param was missing.
// Returns 500 if there was an error getting the options.
func (c *OptionsController) GetByName(g *gin.Context) {
	const op = "OptionsHandler.GetByName"

	name := g.Param("name")
	if name == "" {
		Respond(g, 400, "A name is required to obtain the option by name", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("no name passed"), Operation: op})
		return
	}

	option, err := c.store.Options.GetByName(name)
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained option with name: "+name, option)
}

// UpdateCreate - Restarts the server at the end of the
// request to flush options.
//
// Returns 200 if the options was created/updated.
// Returns 400 if the validation failed on both structs.
// Returns 500 if there was an error updating/creating the options.
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

	if err := c.store.Options.UpdateCreate(options); err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully created/updated options", nil)

	go func() {
		time.Sleep(time.Second * 2)
		reload.Exec()
	}()
}
