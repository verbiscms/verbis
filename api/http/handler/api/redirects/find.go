// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Find
//
// Returns http.StatusOK if the redirect was obtained.
// Returns http.StatusInternalServerError if there as an error obtaining the redirect.
// Returns http.StatusBadRequest if the ID wasn't passed or failed to convert.
func (r *Redirects) Find(ctx *gin.Context) {
	const op = "RedirectHandler.Find"

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Pass a valid number to obtain the redirect by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	redirect, err := r.Store.Redirects.GetByID(id)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained redirect with ID: "+strconv.FormatInt(redirect.Id, 10), redirect)
}
