// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// Login the user.
//
// Returns 200 if login was successful.
// Returns 400 if the validation failed.
// Returns 401 if the credentials didn't match.
func (a *Auth) Login(ctx *gin.Context) {
	const op = "AuthHandler.Login"

	var l domain.Login
	if err := ctx.ShouldBindJSON(&l); err != nil {
		api.Respond(ctx, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user, err := a.Store.Auth.Authenticate(l.Email, l.Password)
	if err != nil {
		api.Respond(ctx, 401, errors.Message(err), err)
		return
	}
	user.HidePassword()

	ctx.SetCookie("verbis-session", user.Token, 172800, "/", "", false, true)

	api.Respond(ctx, 200, "Successfully logged in & session started", user)
}
