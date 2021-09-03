// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
	"strconv"
)

// Update
//
// Returns http.StatusOK if the post was updated.
// Returns http.StatusInternalServerError if there was an error updating or formatting the post.
// Returns http.StatusBadRequest if the the validation failed, there was a conflict, or the post wasn't found.
func (c *Posts) Update(ctx *gin.Context) {
	const op = "PostHandler.Update"

	var post domain.PostCreate
	err := ctx.ShouldBindJSON(&post)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "A valid ID is required to update the post", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	post.ID = id

	updatedPost, err := c.Store.Posts.Update(post)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	// Clear the field cache
	err = c.Cache.Invalidate(ctx, cache.InvalidateOptions{
		Tags: []string{"posts"},
	})
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully updated post with ID: "+strconv.Itoa(updatedPost.ID), updatedPost)
}
