// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"github.com/verbiscms/verbis/api/http/pagination"
	"net/http"
)

// List
//
// Returns http.StatusOK if there are no redirects or success.
// Returns http.StatusInternalServerError if there was an error getting the redirects.
// Returns http.StatusBadRequest if there was conflict or the request was invalid.
func (r *Redirects) List(ctx *gin.Context) {
	const op = "RedirectHandler.List"

	p := api.Params(ctx).Get()

	redirects, total, err := r.Store.Redirects.List(p)
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
