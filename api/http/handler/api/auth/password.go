// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
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
// Returns http.StatusOK if successful.
// Returns http.StatusBadRequest if the ID wasn't passed or failed to convert.
func (a *Auth) ResetPassword(ctx *gin.Context) {
	const op = "AuthHandler.ResetPassword"

	var rp ResetPassword
	err := ctx.ShouldBindJSON(&rp)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user := domain.User{}
	err = a.Cache.Get(ctx, rp.Token, &user)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "No user exists with the token: "+rp.Token, &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	err = a.Store.Auth.ResetPassword(user.Email, rp.NewPassword)
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully reset password", nil)
}
