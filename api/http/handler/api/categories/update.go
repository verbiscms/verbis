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

// Update
//
// Returns 200 if the category was updated.
// Returns 500 if there was an error updating the category.
// Returns 400 if the the validation failed or the category wasn't found.
func (c *Categories) Update(ctx *gin.Context) {
	const op = "CategoryHandler.Update"

	var category domain.Category
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		api.Respond(ctx, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, 400, "A valid ID is required to update the category", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	category.Id = id

	updatedCategory, err := c.Store.Categories.Update(&category)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	defer c.clearCache(updatedCategory.Id)

	api.Respond(ctx, 200, "Successfully updated category with ID: "+strconv.Itoa(category.Id), updatedCategory)
}