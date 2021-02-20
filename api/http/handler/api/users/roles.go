// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// Roles
//
// Returns 200 if the user roles were obtained.
// Returns 500 if there as an error obtaining the user roles.
func (u *Users) Roles(ctx *gin.Context) {
	const op = "UserHandler.Roles"

	roles, err := u.Store.User.GetRoles()
	if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully obtained user roles", roles)
}
