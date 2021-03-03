// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/http/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

// List
//
// Get all posts, obtain resource param to pass to the get
// function.
//
// Returns http.StatusOK if there are no posts or success.
// Returns http.StatusBadRequest if there was conflict or the request was invalid.
// Returns http.StatusInternalServerError if there was an error getting or formatting the posts.
func (c *Posts) List(ctx *gin.Context) {
	const op = "PostHandler.List"

	p := api.Params(ctx).Get()

	posts, total, err := c.Store.Posts.Get(p, true, ctx.Query("resource"), ctx.Query("status"))
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained posts", posts, pagination.Get(p, total))
}
