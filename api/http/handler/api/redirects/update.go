package redirects

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Update
//
// Returns 200 if the redirect was updated.
// Returns 500 if there was an error updating the redirect.
// Returns 400 if the the validation failed or the redirect wasn't found.
func (r *Redirects) Update(g *gin.Context) {
	const op = "RedirectHandler.Update"

	var redirect domain.Redirect
	if err := g.ShouldBindJSON(&redirect); err != nil {
		api.Respond(g, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		api.Respond(g, 400, "A valid ID is required to update the redirect", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	redirect.Id = id

	updatedForm, err := r.Store.Redirects.Update(&redirect)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(g, 500, errors.Message(err), err)
		return
	}

	api.Respond(g, 200, "Successfully updated redirect with ID: "+strconv.FormatInt(redirect.Id, 10), updatedForm)
}