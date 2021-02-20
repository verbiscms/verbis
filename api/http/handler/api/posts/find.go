// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Find
//
// Returns 200 if the posts were obtained.
// Returns 400 if the ID wasn't passed or failed to convert.
// Returns 500 if there as an error obtaining or formatting the post.
func (c *Posts) Find(ctx *gin.Context) {
	const op = "PostHandler.Find"

	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		api.Respond(ctx, 400, "Pass a valid number to obtain the post by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	post, err := c.Store.Posts.GetById(id, true)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 200, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully obtained post with ID: "+paramId, post)
}
