// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

// FieldHandler defines methods for fields to interact with the server
type FieldHandler interface {
	Get(g *gin.Context)
}

// Fields defines the handler for Fields
type Fields struct {
	*deps.Deps
}

// newFields - Construct
func NewFields(d *deps.Deps) *Fields {
	return &Fields{d}
}

// Get - Filter fields and get layouts based on query params.
//
// Returns 200 if login was successful.
// Returns 500 if the layouts failed to be obtained.
func (c *Fields) Get(g *gin.Context) {
	const op = "FieldHandler.Get"

	resource := g.Query("resource")

	userId, err := strconv.Atoi(g.Query("user_id"))
	if err != nil || userId == 0 {
		owner, err := c.Store.User.GetOwner()
		if err != nil {
			Respond(g, 500, errors.Message(err), err)
		}
		userId = owner.Id
	}

	categoryId, err := strconv.Atoi(g.Query("category_id"))
	if err != nil {
		categoryId = 0
		//Respond(g, 400, "Field search failed, wrong type passed to category id", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})x
	}

	post := domain.PostData{
		Post: domain.Post{
			Id:                0,
			Slug:              "",
			Title:             "",
			Status:            "",
			Resource:          &resource,
			PageTemplate:      g.Query("page_template"),
			PageLayout:        g.Query("layout"),
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

	Respond(g, 200, "Successfully obtained fields", fields)
}
