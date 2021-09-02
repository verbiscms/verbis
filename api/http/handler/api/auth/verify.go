// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

// VerifyPasswordToken
//
// Checks to see if the token is valid for resetting.
//
// Returns http.StatusOK if successful.
// Returns http.StatusNotFound if the token does not exist.
func (a *Auth) VerifyPasswordToken(ctx *gin.Context) {
	const op = "AuthHandler.VerifyPasswordToken"

	token := ctx.Param("token")

	_, err := a.Cache.Get(ctx, token, nil)
	if err != nil {
		api.Respond(ctx, http.StatusNotFound, "No user exists with the token: "+token, &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully verified token", nil)
}
