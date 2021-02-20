// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// VerifyPasswordToken
//
// Checks to see if the token is valid for resetting.
//
// Returns 200 if successful.
// Returns 404 if the token does not exist.
func (a *Auth) VerifyPasswordToken(ctx *gin.Context) {
	const op = "AuthHandler.VerifyPasswordToken"

	err := a.Store.Auth.VerifyPasswordToken(ctx.Param("token"))
	if err != nil {
		api.Respond(ctx, 404, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully verified token", nil)
}
