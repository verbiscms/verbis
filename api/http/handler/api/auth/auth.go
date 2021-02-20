// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/gin-gonic/gin"
)

// AuthHandler defines methods for auth methods to interact with the server
type Handler interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
	VerifyPasswordToken(ctx *gin.Context)
	SendResetPassword(ctx *gin.Context)
}

// Auth defines the handler for Authentication methods
type Auth struct {
	*deps.Deps
}
