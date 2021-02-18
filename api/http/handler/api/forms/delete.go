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

// Delete
//
// Returns 200 if the form was deleted.
// Returns 500 if there was an error deleting the form.
// Returns 400 if the the form wasn't found or no ID was passed.
func (f *Forms) Delete(ctx *gin.Context) {
	const op = "FormHandler.Delete"

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, 400, "A valid ID is required to delete a form", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = f.Store.Forms.Delete(id)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully deleted form with ID: "+strconv.Itoa(id), nil)
}