// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Create
//
// Returns 200 if the category was created.
// Returns 500 if there was an error creating the category.
// Returns 400 if the the validation failed or there was a conflict.
func (c *Categories) Create(ctx *gin.Context) {
	const op = "CategoryHandler.Create"

	var category domain.Category
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		api.Respond(ctx, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newCategory, err := c.Store.Categories.Create(&category)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully created category with ID: "+strconv.Itoa(category.Id), newCategory)
}
