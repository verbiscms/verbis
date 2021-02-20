// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// SendResetPassword defines the data to be validated when a
// user sends the reset email.
type SendResetPassword struct {
	Email string `json:"email" binding:"required,email"`
}

// SendResetPassword
//
// Reset password email & generate token.
//
// Returns 200 if successful.
// Returns 400 if validation failed the user wasn't found.
func (a *Auth) SendResetPassword(ctx *gin.Context) {
	const op = "AuthHandler.SendResetPassword"

	var srp SendResetPassword
	err := ctx.ShouldBindJSON(&srp)
	if err != nil {
		api.Respond(ctx, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = a.Store.Auth.SendResetPassword(srp.Email)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "A fresh verification link has been sent to your email", nil)
}
