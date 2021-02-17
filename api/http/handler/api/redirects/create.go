package redirects

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Create
//
// Returns 200 if the redirect was created.
// Returns 500 if there was an error creating the redirect.
// Returns 400 if the the validation failed or there was a conflict.
func (r *Redirects) Create(g *gin.Context) {
	const op = "RedirectHandler.Create"

	var redirect domain.Redirect
	if err := g.ShouldBindJSON(&redirect); err != nil {
		api.Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newForm, err := r.Store.Redirects.Create(&redirect)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(g, 500, errors.Message(err), err)
		return
	}

	api.Respond(g, 200, "Successfully created redirect with ID: "+strconv.FormatInt(redirect.Id, 10), newForm)
}