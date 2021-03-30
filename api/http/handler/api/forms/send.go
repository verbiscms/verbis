// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package forms

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Send
//
// Returns http.StatusOK if the form was deleted.
// Returns http.StatusInternalServerError if there was an error deleting the form.
// Returns http.StatusBadRequest if the the form wasn't found or no ID was passed.
func (f *Forms) Send(ctx *gin.Context) {
	const op = "FormHandler.Send"

	form, err := f.Store.Forms.GetByUUID(ctx.Param("uuid"))
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	err = ctx.ShouldBind(form.Body)
	if err != nil {
		// If file has an empty value, no validation data is returned.
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = f.Store.Forms.Send(&form, ctx.ClientIP(), ctx.Request.UserAgent())
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully sent form with ID: "+strconv.Itoa(form.Id), nil)
}
