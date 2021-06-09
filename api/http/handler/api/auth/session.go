// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckSession
//
// Returns http.StatusOK if the user is authenticated (from middleware).
func (a *Auth) CheckSession(ctx *gin.Context) {
	const op = "AuthHandler.CheckSession"
	api.Respond(ctx, http.StatusOK, "Session valid", nil)
}
