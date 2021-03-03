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

// Logout
//
// Returns http.StatusOK if logout was successful.
// Returns http.StatusBadRequest if the user wasn't found.
// Returns http.StatusInternalServerError if there was an error logging out.
func (a *Auth) Logout(ctx *gin.Context) {
	const op = "AuthHandler.Logout"

	token := ctx.Request.Header.Get("token")
	_, err := a.Store.Auth.Logout(token)
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusBadRequest, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	ctx.SetCookie("verbis-session", "", -1, "/", "", false, true)

	api.Respond(ctx, http.StatusOK, "Successfully logged out", nil)
}
