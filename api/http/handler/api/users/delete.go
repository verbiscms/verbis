// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
	"strconv"
)

// Delete
//
// Returns http.StatusOK if the user was deleted.
// Returns http.StatusInternalServerError if there was an error deleting the user.
// Returns http.StatusBadRequest if the the user wasn't found or no ID was passed.
func (u *Users) Delete(ctx *gin.Context) {
	const op = "UserHandler.Delete"

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "A valid ID is required to delete a user", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = u.Store.User.Delete(id)
	if errors.Code(err) == errors.NOTFOUND || errors.Code(err) == errors.CONFLICT {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully deleted user with ID: "+strconv.Itoa(id), nil)
}
