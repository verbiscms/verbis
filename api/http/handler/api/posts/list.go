// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/http"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// List
//
// Get all posts, obtain resource param to pass to the get
// function.
//
// Returns 200 if there are no posts or success.
// Returns 400 if there was conflict or the request was invalid.
// Returns 500 if there was an error getting or formatting the posts.
func (c *Posts) List(ctx *gin.Context) {
	const op = "PostHandler.List"

	p := params.ApiParams(ctx, api.DefaultParams).Get()

	posts, total, err := c.Store.Posts.Get(p, true, ctx.Query("resource"), ctx.Query("status"))
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 200, errors.Message(err), err)
		return
	} else if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	pagination := http.NewPagination().Get(p, total)

	api.Respond(ctx, 200, "Successfully obtained posts", posts, pagination)
}