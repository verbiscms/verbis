package api

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler"
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

// Options defines the handler for Options
type Options struct {
	store  *models.Store
	config config.Configuration
}

// newOptions - Construct
func NewwOptions(m *models.Store, config config.Configuration) *Options {
	return &Options{
		store:  m,
		config: config,
	}
}

// Get All
//
// Returns 200 if there are no options or success.
// Returns 500 if there was an error getting the options.
func (c *Options) Get(g *gin.Context) {
	const op = "OptionsHandler.Delete"

	options, err := c.store.Options.Get()
	if errors.Code(err) == errors.NOTFOUND {
		handler.Respond(g, 200, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		handler.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	handler.Respond(g, 200, "Successfully obtained options", options)
}

// Get By name
//
// Returns 200 if there are no options or success.
// Returns 400 if there was name param was missing.
// Returns 500 if there was an error getting the options.
func (c *Options) GetByName(g *gin.Context) {
	const op = "OptionsHandler.GetByName"

	name := g.Param("name")
	option, err := c.store.Options.GetByName(name)
	if errors.Code(err) == errors.NOTFOUND {
		handler.Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	handler.Respond(g, 200, "Successfully obtained option with name: "+name, option)
}

// UpdateCreate - Restarts the server at the end of the
// request to flush options.
//
// Returns 200 if the options was created/updated.
// Returns 400 if the validation failed on both structs.
// Returns 500 if there was an error updating/creating the options.
func (c *Options) UpdateCreate(g *gin.Context) {
	const op = "OptionsHandler.UpdateCreate"

	var options domain.OptionsDB
	if err := g.ShouldBindBodyWith(&options, binding.JSON); err != nil {
		handler.Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	var vOptions domain.Options
	if err := g.ShouldBindBodyWith(&vOptions, binding.JSON); err != nil {
		handler.Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if err := c.store.Options.UpdateCreate(&options); err != nil {
		handler.Respond(g, 500, errors.Message(err), err)
		return
	}

	handler.Respond(g, 200, "Successfully created/updated options", nil)

	go func() {
		time.Sleep(time.Second * 2)
		reload.Exec()
	}()
}
