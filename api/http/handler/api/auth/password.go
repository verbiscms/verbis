// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// ResetPassword defines the data to be validated when a
// user resets a password.
type ResetPassword struct {
	NewPassword     string `json:"new_password" binding:"required,min=8,max=60"`
	ConfirmPassword string `json:"confirm_password" binding:"eqfield=NewPassword,required"`
	Token           string `json:"token" binding:"required"`
}

// ResetPassword
//
// Returns 200 if successful.
// Returns 400 if the ID wasn't passed or failed to convert.
func (a *Auth) ResetPassword(ctx *gin.Context) {
	const op = "AuthHandler.ResetPassword"

	var rp ResetPassword
	err := ctx.ShouldBindJSON(&rp)
	if err != nil {
		api.Respond(ctx, 400, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = a.Store.Auth.ResetPassword(rp.Token, rp.NewPassword)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 400, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully reset password", nil)
}
