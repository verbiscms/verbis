// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
	"time"
)

// UpdateCreate
//
// Restarts the server at the end of the request
// to flush options.
//
// Returns http.StatusOK if the options was created/updated.
// Returns http.StatusBadRequest if the validation failed on both structs.
// Returns http.StatusInternalServerError if there was an error updating/creating the options.
func (o *Options) UpdateCreate(ctx *gin.Context) {
	const op = "OptionsHandler.UpdateCreate"

	var options domain.OptionsDBMap
	err := ctx.ShouldBindBodyWith(&options, binding.JSON)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	var vOptions domain.Options
	err = ctx.ShouldBindBodyWith(&vOptions, binding.JSON)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = o.Store.Options.Insert(options)
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	err = o.Cache.Invalidate(ctx, cache.InvalidateOptions{Tags: []string{"options"}})
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	o.Cache.Set(ctx, cache.OptionsKey, vOptions, cache.Options{Expiration: time.Minute * 15})
	o.SetOptions(&vOptions)

	api.Respond(ctx, http.StatusOK, "Successfully created/updated options", nil)
}
