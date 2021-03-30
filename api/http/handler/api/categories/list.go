// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/http/pagination"
	store "github.com/ainsleyclark/verbis/api/store/categories"
	"github.com/gin-gonic/gin"
	"net/http"
)

// List
//
// Returns http.StatusOK if there are no categories or success.
// Returns http.StatusInternalServerError if there was an error getting the categories.
// Returns http.StatusBadRequest if there was conflict or the request was invalid.
func (c *Categories) List(ctx *gin.Context) {
	const op = "CategoryHandler.List"

	p := api.Params(ctx).Get()

	cfg := store.ListConfig{
		Resource: ctx.Query("resource"),
	}

	categories, total, err := c.Store.Categories.List(p, cfg)
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

	api.Respond(ctx, http.StatusOK, "Successfully obtained categories", categories, pagination.Get(p, total))
}
