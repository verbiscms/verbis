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

// Update
//
// Returns 200 if the user was updated.
// Returns 500 if there was an error updating the user.
// Returns 400 if the the validation failed or the user wasn't found.
func (u *Users) Update(ctx *gin.Context) {
	const op = "UserHandler.Update"

	var user domain.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		api.Respond(ctx, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, 400, "A valid ID is required to update the user", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}
	user.Id = id

	updatedUser, err := u.Store.User.Update(&user)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	defer u.clearCache(updatedUser.Id)

	api.Respond(ctx, 200, "Successfully updated user with ID: "+strconv.Itoa(user.Id), updatedUser)
}
