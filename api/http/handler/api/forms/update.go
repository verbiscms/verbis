// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
	"strconv"
)

// Update
//
// Returns http.StatusOK if the form was updated.
// Returns http.StatusInternalServerError if there was an error updating the form.
// Returns http.StatusBadRequest if the the validation failed or the form wasn't found.
func (f *Forms) Update(ctx *gin.Context) {
	const op = "FormHandler.Update"

	var form domain.Form
	if err := ctx.ShouldBindJSON(&form); err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "A valid ID is required to update the form", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	form.ID = id

	updatedForm, err := f.Store.Forms.Update(form)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully updated form with ID: "+strconv.Itoa(form.ID), updatedForm)
}
