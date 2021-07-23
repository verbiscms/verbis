// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
	"strconv"
)

// Create
//
// Returns http.StatusOK if the category was created.
// Returns http.StatusInternalServerError if there was an error creating the category.
// Returns http.StatusBadRequest if the the validation failed or there was a conflict.
func (c *Categories) Create(ctx *gin.Context) {
	const op = "CategoryHandler.Create"

	var category domain.Category
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newCategory, err := c.Store.Categories.Create(category)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully created category with ID: "+strconv.Itoa(category.Id), newCategory)
}
