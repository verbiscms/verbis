// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/events"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
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

	user, token, err := a.Store.Auth.SendResetPassword(srp.Email)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	err = a.resetPassword.Dispatch(events.ResetPassword{
		User: user,
		URL:  a.Deps.Options.SiteUrl + "/admin/password/reset/" + token,
	}, []string{user.Email}, nil)

	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "A fresh verification link has been sent to your email", nil)
}
