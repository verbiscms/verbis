// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/teamwork/reload"
	"time"
)

// UpdateCreate
//
// Restarts the server at the end of the request
// to flush options.
//
// Returns 200 if the options was created/updated.
// Returns 400 if the validation failed on both structs.
// Returns 500 if there was an error updating/creating the options.
func (o *Options) UpdateCreate(ctx *gin.Context) {
	const op = "OptionsHandler.UpdateCreate"

	var options domain.OptionsDBMap
	err := ctx.ShouldBindBodyWith(&options, binding.JSON)
	if err != nil {
		api.Respond(ctx, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	var vOptions domain.Options
	err = ctx.ShouldBindBodyWith(&vOptions, binding.JSON)
	if err != nil {
		api.Respond(ctx, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = o.Store.Options.UpdateCreate(&options)
	if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully created/updated options", nil)

	go func() {
		// Set the deps options, TODO, were restarting the server here.
		o.SetOptions(&vOptions)
		time.Sleep(time.Second * 2)
		reload.Exec()
	}()
}
