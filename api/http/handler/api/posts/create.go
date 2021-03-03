// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Create
//
// Returns http.StatusOK if the post was created.
// Returns http.StatusInternalServerError if there was an error creating or formatting the post.
// Returns http.StatusBadRequest if the the validation failed or there was a conflict with the post.
func (c *Posts) Create(ctx *gin.Context) {
	const op = "PostHandler.Create"

	var post domain.PostCreate
	err := ctx.ShouldBindJSON(&post)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newPost, err := c.Store.Posts.Create(&post)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully created post with ID: "+strconv.Itoa(newPost.Id), newPost)
}
