// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
	"strconv"
)

// Find
//
// Returns http.StatusOK if the posts were obtained.
// Returns http.StatusBadRequest if the ID wasn't passed or failed to convert.
// Returns http.StatusInternalServerError if there as an error obtaining or formatting the post.
func (c *Posts) Find(ctx *gin.Context) {
	const op = "PostHandler.Find"

	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Pass a valid number to obtain the post by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	post, err := c.Store.Posts.Find(id, true)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained post with ID: "+paramID, post)
}
