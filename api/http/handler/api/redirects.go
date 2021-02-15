package api

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

// RedirectHandler defines methods for Redirect routes to interact with the server.
type RedirectHandler interface {
	Get(g *gin.Context)
	GetById(g *gin.Context)
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// Redirects defines the handler for all Redirect Routes
type Redirects struct {
	*deps.Deps
}

// NewRedirects - Construct
func NewRedirects(d *deps.Deps) *Redirects {
	return &Redirects{d}
}

// Get all redirects
//
// Returns 200 if there are no redirects or success.
// Returns 500 if there was an error getting the redirects.
// Returns 400 if there was conflict or the request was invalid.
func (c *Redirects) Get(g *gin.Context) {
	const op = "RedirectHandler.Get"

	p := params.ApiParams(g, DefaultParams).Get()

	redirects, total, err := c.Store.Redirects.Get(p)
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

	pagination := http.NewPagination().Get(p, total)

	Respond(g, 200, "Successfully obtained redirects", redirects, pagination)
}

// GetById
//
// Returns 200 if the redirect was obtained.
// Returns 500 if there as an error obtaining the redirect.
// Returns 400 if the ID wasn't passed or failed to convert.
func (c *Redirects) GetById(g *gin.Context) {
	const op = "RedirectHandler.GetById"

	id, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		Respond(g, 400, "Pass a valid number to obtain the redirect by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	redirect, err := c.Store.Redirects.GetById(id)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained redirect with ID: "+strconv.FormatInt(redirect.Id, 10), redirect)
}

// Create
//
// Returns 200 if the redirect was created.
// Returns 500 if there was an error creating the redirect.
// Returns 400 if the the validation failed or there was a conflict.
func (c *Redirects) Create(g *gin.Context) {
	const op = "RedirectHandler.Create"

	var redirect domain.Redirect
	if err := g.ShouldBindJSON(&redirect); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newForm, err := c.Store.Redirects.Create(&redirect)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully created redirect with ID: "+strconv.FormatInt(redirect.Id, 10), newForm)
}

// Update
//
// Returns 200 if the redirect was updated.
// Returns 500 if there was an error updating the redirect.
// Returns 400 if the the validation failed or the redirect wasn't found.
func (c *Redirects) Update(g *gin.Context) {
	const op = "RedirectHandler.Update"

	var redirect domain.Redirect
	if err := g.ShouldBindJSON(&redirect); err != nil {
		Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		Respond(g, 400, "A valid ID is required to update the redirect", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	redirect.Id = id

	updatedForm, err := c.Store.Redirects.Update(&redirect)
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}


	Respond(g, 200, "Successfully updated redirect with ID: "+strconv.FormatInt(redirect.Id, 10), updatedForm)
}

// Delete
//
// Returns 200 if the redirect was deleted.
// Returns 500 if there was an error deleting the redirect.
// Returns 400 if the the redirect wasn't found or no ID was passed.
func (c *Redirects) Delete(g *gin.Context) {
	const op = "RedirectHandler.Delete"

	id, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		Respond(g, 400, "A valid ID is required to delete a redirect", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = c.Store.Redirects.Delete(id)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully deleted redirect with ID: "+strconv.FormatInt(id, 10), nil)
}
