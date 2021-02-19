// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Create
//
// Returns 200 if the form was created.
// Returns 500 if there was an error creating the form.
// Returns 400 if the the validation failed or there was a conflict.
func (f *Forms) Create(ctx *gin.Context) {
	const op = "FormHandler.Create"

	var form domain.Form
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		api.Respond(ctx, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	newForm, err := f.Store.Forms.Create(&form)
	if errors.Code(err) == errors.INVALID || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully created form with ID: "+strconv.Itoa(form.Id), newForm)
}