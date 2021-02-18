// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Find
//
// Returns 200 if the user was obtained.
// Returns 500 if there as an error obtaining the user.
// Returns 400 if the ID wasn't passed or failed to convert.
func (u *Users) Find(ctx *gin.Context) {
	const op = "UserHandler.Find"

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		api.Respond(ctx, 400, "Pass a valid number to obtain the user by ID", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user, err := u.Store.User.GetById(id)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 200, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully obtained user with ID: "+strconv.Itoa(id), user.HideCredentials())
}