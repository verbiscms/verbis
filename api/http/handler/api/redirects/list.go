package redirects

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// List all redirects
//
// Returns 200 if there are no redirects or success.
// Returns 500 if there was an error getting the redirects.
// Returns 400 if there was conflict or the request was invalid.
func (r *Redirects) List(g *gin.Context) {
	const op = "RedirectHandler.Get"

	p := params.ApiParams(g, api.DefaultParams).Get()

	redirects, total, err := r.Store.Redirects.Get(p)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(g, 200, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(g, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(g, 500, errors.Message(err), err)
		return
	}

	pagination := http.NewPagination().Get(p, total)

	api.Respond(g, 200, "Successfully obtained redirects", redirects, pagination)
}