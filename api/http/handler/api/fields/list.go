// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Get - Filter fields and get layouts based on query params.
//
// Returns 200 if login was successful.
// Returns 500 if the layouts failed to be obtained.
func (c *Fields) Get(ctx *gin.Context) {
	const op = "FieldHandler.Get"

	resource := ctx.Query("resource")

	userId, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil || userId == 0 {
		owner, err := c.Store.User.GetOwner()
		if err != nil {
			api.Respond(ctx, 500, errors.Message(err), err)
		}
		userId = owner.Id
	}

	categoryId, err := strconv.Atoi(ctx.Query("category_id"))
	if err != nil {
		categoryId = 0
	}

	post := domain.PostData{
		Post: domain.Post{
			Id:                0,
			Slug:              "",
			Title:             "",
			Status:            "",
			Resource:          &resource,
			PageTemplate:      ctx.Query("page_template"),
			PageLayout:        ctx.Query("layout"),
			CodeInjectionHead: nil,
			CodeInjectionFoot: nil,
			UserId:            userId,
		},
	}

	// Get the author associated with the post
	author, err := c.Store.User.GetById(post.UserId)
	if err != nil {
		post.Author = author.HideCredentials()
	}

	// Get the categories associated with the post
	category, err := c.Store.Categories.GetById(categoryId)
	if err != nil {
		post.Category = &category
	}

	fields := c.Store.Fields.GetLayout(post)

	api.Respond(ctx, 200, "Successfully obtained fields", fields)
}