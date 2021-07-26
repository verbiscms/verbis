// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/events"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
	"time"
)

// SendResetPassword defines the data to be validated when a
// user sends the reset email.
type SendResetPassword struct {
	Email string `json:"email" binding:"required,email"`
}

const (
	PasswordExpiry = time.Minute * 15
)

// SendResetPassword
//
// Reset password email & generate token.
//
// Returns http.StatusOK if successful.
// Returns http.StatusBadRequest if validation failed the user wasn't found.
func (a *Auth) SendResetPassword(ctx *gin.Context) {
	const op = "AuthHandler.SendResetPassword"

	var srp SendResetPassword
	err := ctx.ShouldBindJSON(&srp)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user, err := a.Store.User.FindByEmail(srp.Email)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "No user found with email: "+srp.Email, &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	token, err := a.generateTokenFunc(user.Email)
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, "Error generating user token", &errors.Error{Code: errors.INTERNAL, Err: err, Operation: op})
		return
	}

	err = cache.Set(context.Background(), token, user, cache.Options{Expiration: PasswordExpiry})
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, "Error sending password reset", &errors.Error{Code: errors.INTERNAL, Err: err, Operation: op})
		return
	}

	err = a.resetPassword.Dispatch(events.ResetPassword{
		User: user.UserPart,
		URL:  a.Deps.Options.SiteUrl + "/admin/password/reset/" + token,
	}, []string{user.Email}, nil)

	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "A fresh verification link has been sent to your email", nil)
}
