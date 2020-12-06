package api

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
)

// FormHandler defines methods for Form routes to interact with the server
type FormHandler interface {
	Send(g *gin.Context)
}

// Forms defines the handler for all Form Routes
type Forms struct {
	store  *models.Store
	config config.Configuration
}

// NewForms - Construct
func NewForms(m *models.Store, config config.Configuration) *Forms {
	return &Forms{
		store:  m,
		config: config,
	}
}

func (c *Forms) Send(g *gin.Context) {
	const op = "FormHandler.Send"

	form, err := c.store.Forms.GetByUUID(g.Param("uuid"))
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	if err := g.ShouldBindJSON(form.Body); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if err := c.store.Forms.Send(&form, g.ClientIP(), g.Request.UserAgent()); err != nil {
		Respond(g, 500, errors.Message(err), err)
	}

	Respond(g, 200, "Form submitted & sent successfully", nil)
}
