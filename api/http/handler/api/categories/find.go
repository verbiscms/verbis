// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Find
//
// Returns http.StatusOK if the category was obtained.
// Returns http.StatusInternalServerError if there as an error obtaining the category.
// Returns http.StatusBadRequest if the ID wasn't passed or failed to convert.
func (c *Categories) Find(ctx *gin.Context) {
	const op = "CategoryHandler.Find"

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Pass a valid number to obtain the category by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	category, err := c.Store.Categories.GetByID(id)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained category with ID: "+strconv.Itoa(id), category)
}
