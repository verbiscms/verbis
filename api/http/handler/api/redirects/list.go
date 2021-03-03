// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/http/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

// List
//
// Returns 200 if there are no redirects or success.
// Returns 500 if there was an error getting the redirects.
// Returns 400 if there was conflict or the request was invalid.
func (r *Redirects) List(ctx *gin.Context) {
	const op = "RedirectHandler.List"

	p := api.Params(ctx).Get()

	redirects, total, err := r.Store.Redirects.Get(p)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained redirects", redirects, pagination.Get(p, total))
}
