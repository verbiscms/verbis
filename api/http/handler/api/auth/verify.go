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

// VerifyPasswordToken
//
// Checks to see if the token is valid for resetting.
//
// Returns http.StatusOK if successful.
// Returns http.StatusNotFound if the token does not exist.
func (a *Auth) VerifyPasswordToken(ctx *gin.Context) {
	const op = "AuthHandler.VerifyPasswordToken"

	_, err := a.Store.Auth.VerifyPasswordToken(ctx.Param("token"))
	if err != nil {
		api.Respond(ctx, http.StatusNotFound, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully verified token", nil)
}
