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

// Find
//
// Returns 200 if the media items were obtained.
// Returns 400 if the ID wasn't passed or failed to convert.
// Returns 500 if there as an error obtaining the media items.
func (m *Media) Find(ctx *gin.Context) {
	const op = "MediaHandler.Find"

	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		api.Respond(ctx, 400, "Pass a valid number to obtain the media item by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	media, err := m.Store.Media.GetByID(id)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 200, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully obtained media item with ID: "+paramID, media)
}
