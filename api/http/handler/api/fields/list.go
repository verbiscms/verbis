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

// List
//
// Filter fields and get layouts based on query params.
//
// Returns 200 if login was successful.
// Returns 500 if the layouts failed to be obtained.
func (c *Fields) List(ctx *gin.Context) {
	const op = "FieldHandler.List"

	resource := ctx.Query("resource")

	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil || userID == 0 {
		owner, err := c.Store.User.GetOwner()
		if err != nil {
			api.Respond(ctx, 500, errors.Message(err), err)
		}
		userID = owner.Id
	}

	categoryID, err := strconv.Atoi(ctx.Query("category_id"))
	if err != nil {
		categoryID = 0
	}

	post := domain.PostDatum{
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
			UserId:            userID,
		},
	}

	// Get the author associated with the post
	author, err := c.Store.User.GetByID(post.UserId)
	if err != nil {
		post.Author = author.HideCredentials()
	}

	// Get the categories associated with the post
	category, err := c.Store.Categories.GetByID(categoryID)
	if err != nil {
		post.Category = &category
	}

	fields := c.Store.Fields.GetLayout(post)

	api.Respond(ctx, 200, "Successfully obtained fields", fields)
}
