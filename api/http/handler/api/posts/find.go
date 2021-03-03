// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Find
//
// Returns 200 if the posts were obtained.
// Returns 400 if the ID wasn't passed or failed to convert.
// Returns 500 if there as an error obtaining or formatting the post.
func (c *Posts) Find(ctx *gin.Context) {
	const op = "PostHandler.Find"

	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		api.Respond(ctx, 400, "Pass a valid number to obtain the post by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	post, err := c.Store.Posts.GetByID(id, true)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained post with ID: "+paramID, post)
}
