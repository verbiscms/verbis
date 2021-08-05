// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
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

	var item domain.Media
	err := ctx.ShouldBindJSON(&item)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "A valid ID is required to update the media item", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	item.ID = id

	updated, err := m.service.Update(item)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully updated media item with ID: "+strconv.Itoa(id), updated.Public())
}
