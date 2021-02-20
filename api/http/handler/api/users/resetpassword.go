// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ResetPassword
//
// Returns 200 if the reset password was successful.
// Returns 500 if there was an error resetting the user failed.
// Returns 400 if the the user wasn't found, no ID was passed or validation failed.
func (u *Users) ResetPassword(ctx *gin.Context) {
	const op = "UserHandler.ResetPassword"

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, 400, "A valid ID is required to update a user's password", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user, err := u.Store.User.GetById(id)
	if err != nil {
		api.Respond(ctx, 400, "No user has been found with the ID: "+strconv.Itoa(id), err)
		return
	}

	var reset domain.UserPasswordReset
	reset.DBPassword = user.Password
	if err := ctx.ShouldBindJSON(&reset); err != nil {
		api.Respond(ctx, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = u.Store.User.ResetPassword(id, reset)
	if errors.Code(err) == errors.INVALID {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully updated password for the user with ID: "+strconv.Itoa(id), nil)
}
