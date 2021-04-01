// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Create
//
// Returns http.StatusOK if the redirect was created.
// Returns http.StatusInternalServerError if there was an error creating the redirect.
// Returns http.StatusBadRequest if the the validation failed or there was a conflict.
func (r *Redirects) Create(ctx *gin.Context) {
	const op = "RedirectHandler.Create"

	var redirect domain.Redirect
	err := ctx.ShouldBindJSON(&redirect)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newRedirect, err := r.Store.Redirects.Create(redirect)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully created redirect with ID: "+strconv.Itoa(newRedirect.Id), newRedirect)
}
