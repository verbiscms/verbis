package api

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	params2 "github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// FormHandler defines methods for Form routes to interact with the server
type FormHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	Create(g *gin.Context)
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

// Get all forms
//
// Returns 200 if there are no forms or success.
// Returns 500 if there was an error getting the forms.
// Returns 400 if there was conflict or the request was invalid.
func (c *Forms) Get(g *gin.Context) {
	const op = "FormHandler.Get"

	params := params2.ApiParams(g, DefaultParams).Get()

	forms, total, err := c.store.Forms.Get(params)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	pagination := http.NewPagination().Get(params, total)

	Respond(g, 200, "Successfully obtained forms", forms, pagination)
}

// GetById
//
// Returns 200 if the form was obtained.
// Returns 500 if there as an error obtaining the form.
// Returns 400 if the ID wasn't passed or failed to convert.
func (c *Forms) GetById(g *gin.Context) {
	const op = "FormHandler.GetById"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "Pass a valid number to obtain the form by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	form, err := c.store.Forms.GetById(id)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained form with ID: "+strconv.Itoa(id), form)
}

// Create
//
// Returns 200 if the form was created.
// Returns 500 if there was an error creating the form.
// Returns 400 if the the validation failed or there was a conflict.
func (c *Forms) Create(g *gin.Context) {
	const op = "FormHandler.Create"

	var form domain.Form
	if err := g.ShouldBindJSON(&form); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newForm, err := c.store.Forms.Create(&form)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully created form with ID: "+strconv.Itoa(form.Id), newForm)
}

// Update
//
// Returns 200 if the form was updated.
// Returns 500 if there was an error updating the form.
// Returns 400 if the the validation failed or the form wasn't found.
func (c *Forms) Update(g *gin.Context) {
	const op = "FormHandler.Update"

	var form domain.Form
	if err := g.ShouldBindJSON(&form); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "A valid ID is required to update the form", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	form.Id = id

	updatedForm, err := c.store.Forms.Update(&form)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully updated form with ID: "+strconv.Itoa(form.Id), updatedForm)
}

// Delete
//
// Returns 200 if the form was deleted.
// Returns 500 if there was an error deleting the form.
// Returns 400 if the the form wasn't found or no ID was passed.
func (c *Forms) Delete(g *gin.Context) {
	const op = "FormHandler.Delete"

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Respond(g, 400, "A valid ID is required to delete a form", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = c.store.Forms.Delete(id)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully deleted form with ID: "+strconv.Itoa(id), nil)
}

//

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

	err = g.Bind(form.Body)
	if err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	//	color.Red.Printf("%+v\n", form.Body)

	err = c.store.Forms.Send(&form, g.ClientIP(), g.Request.UserAgent())
	if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Form submitted & sent successfully", nil)
}
