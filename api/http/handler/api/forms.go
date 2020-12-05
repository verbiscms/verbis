package api

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/ompluscator/dynamic-struct"
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

	instance := dynamicstruct.NewStruct()
	for k, v := range form.Fields {
		tag := fmt.Sprintf(`json:"%s" binding:"required"`, v.Key)
		instance.AddField(k, "", tag)
	}

	test := instance.Build().New()

	if err := g.ShouldBindJSON(test); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}


	fmt.Println(instance.GetField("FirstName"))


	Respond(g, 200, "Passed", nil)
}


