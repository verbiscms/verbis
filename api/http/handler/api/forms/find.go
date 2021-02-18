// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Find
//
// Returns 200 if the form was obtained.
// Returns 500 if there as an error obtaining the form.
// Returns 400 if the ID wasn't passed or failed to convert.
func (f *Forms) Find(ctx *gin.Context) {
	const op = "FormHandler.GetById"

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, 400, "Pass a valid number to obtain the form by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	form, err := f.Store.Forms.GetById(id)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 200, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully obtained form with ID: "+strconv.Itoa(id), form)
}