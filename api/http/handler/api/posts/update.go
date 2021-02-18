// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Update
//
// Returns 200 if the post was updated.
// Returns 500 if there was an error updating or formatting the post.
// Returns 400 if the the validation failed, there was a conflict, or the post wasn't found.
func (c *Posts) Update(ctx *gin.Context) {
	const op = "PostHandler.Update"

	var post domain.PostCreate
	if err := ctx.ShouldBindJSON(&post); err != nil {
		api.Respond(ctx, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	// TODO: This needs to be in the store.
	cache.ClearPostCache(post.Id)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, 400, "A valid ID is required to update the post", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	post.Id = id

	updatedPost, err := c.Store.Posts.Update(&post)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully updated post with ID: "+strconv.Itoa(updatedPost.Id), updatedPost)
}