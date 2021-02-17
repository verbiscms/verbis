package redirects

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Delete
//
// Returns 200 if the redirect was deleted.
// Returns 500 if there was an error deleting the redirect.
// Returns 400 if the the redirect wasn't found or no ID was passed.
func (r *Redirects) Delete(g *gin.Context) {
	const op = "RedirectHandler.Delete"

	id, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		api.Respond(g, 400, "A valid ID is required to delete a redirect", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = r.Store.Redirects.Delete(id)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		api.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(g, 500, errors.Message(err), err)
		return
	}

	api.Respond(g, 200, "Successfully deleted redirect with ID: "+strconv.FormatInt(id, 10), nil)
}