package redirects

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Find
//
// Returns 200 if the redirect was obtained.
// Returns 500 if there as an error obtaining the redirect.
// Returns 400 if the ID wasn't passed or failed to convert.
func (r *Redirects) Find(g *gin.Context) {
	const op = "RedirectHandler.GetById"

	id, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		api.Respond(g, 400, "Pass a valid number to obtain the redirect by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	redirect, err := r.Store.Redirects.GetById(id)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(g, 500, errors.Message(err), err)
		return
	}

	api.Respond(g, 200, "Successfully obtained redirect with ID: "+strconv.FormatInt(redirect.Id, 10), redirect)
}