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

// Delete
//
// Returns http.StatusOK if the redirect was deleted.
// Returns http.StatusInternalServerError if there was an error deleting the redirect.
// Returns http.StatusBadRequest if the the redirect wasn't found or no ID was passed.
func (r *Redirects) Delete(ctx *gin.Context) {
	const op = "RedirectHandler.Delete"

	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "A valid ID is required to delete a redirect", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = r.Store.Redirects.Delete(id)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully deleted redirect with ID: "+paramID, nil)
}
