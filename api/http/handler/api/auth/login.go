// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login defines the data to be validated when a
// user logins into the SPA.
type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login the user.
//
// Returns http.StatusOK if login was successful.
// Returns http.StatusBadRequest if the validation failed.
// Returns http.StatusUnauthorized if the credentials didn't match.
func (a *Auth) Login(ctx *gin.Context) {
	const op = "AuthHandler.Login"

	var l Login
	err := ctx.ShouldBindJSON(&l)
	if err != nil {
		api.Respond(ctx, http.StatusBadRequest, "Validation failed", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	user, err := a.Store.Auth.Login(l.Email, l.Password)
	if err != nil {
		api.Respond(ctx, http.StatusUnauthorized, errors.Message(err), err)
		return
	}
	user.HidePassword()

	ctx.SetCookie("verbis-session", user.Token, 172800, "/", "", false, true)

	api.Respond(ctx, http.StatusOK, "Successfully logged in & session started", user)
}
