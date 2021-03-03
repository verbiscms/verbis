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

// Update
//
// Returns http.StatusOK if the redirect was updated.
// Returns http.StatusInternalServerError if there was an error updating the redirect.
// Returns http.StatusBadRequest if the the validation failed or the redirect wasn't found.
func (r *Redirects) Update(ctx *gin.Context) {
	const op = "RedirectHandler.Update"

	var redirect domain.Redirect
	err := ctx.ShouldBindJSON(&redirect)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "A valid ID is required to update the redirect", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	redirect.Id = id

	updatedForm, err := r.Store.Redirects.Update(&redirect)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully updated redirect with ID: "+strconv.FormatInt(redirect.Id, 10), updatedForm)
}
