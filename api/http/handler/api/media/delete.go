// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Delete
//
// Returns 200 if the media item was deleted.
// Returns 500 if there was an error updating the media item.
// Returns 400 if the the media item wasn't found or no ID was passed.
func (m *Media) Delete(ctx *gin.Context) {
	const op = "MediaHandler.Delete"

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, 400, "A valid ID is required to delete a media item", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = m.Store.Media.Delete(id)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully deleted media item with ID: "+strconv.Itoa(id), nil)
}