// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

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
// Returns http.StatusOK if the media item was updated successfully.
// Returns http.StatusBadRequest if the ID wasn't passed or failed to convert.
// Returns http.StatusInternalServerError if there was an error updating the media item.
func (m *Media) Update(ctx *gin.Context) {
	const op = "MediaHandler.Update"

	var media domain.Media
	err := ctx.ShouldBindJSON(&m)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "A valid ID is required to update the media item", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	media.Id = id

	err = m.Store.Media.Update(&media)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully updated media item with ID: "+strconv.Itoa(id), media)
}
