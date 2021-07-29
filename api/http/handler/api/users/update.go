// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

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
// Returns http.StatusOK if the user was updated.
// Returns http.StatusInternalServerError if there was an error updating the user.
// Returns http.StatusBadRequest if the the validation failed or the user wasn't found.
func (u *Users) Update(ctx *gin.Context) {
	const op = "UserHandler.Update"

	var user domain.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "A valid ID is required to update the user", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	user.Id = id

	updatedUser, err := u.Store.User.Update(user)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully updated user with ID: "+strconv.Itoa(user.Id), updatedUser)
}
