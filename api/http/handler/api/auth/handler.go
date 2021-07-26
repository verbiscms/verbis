// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/common/encryption"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/events"
)

// AuthHandler defines methods for auth methods to interact with the server
type Handler interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
	VerifyPasswordToken(ctx *gin.Context)
	SendResetPassword(ctx *gin.Context)
	CheckSession(ctx *gin.Context)
}

// Auth defines the handler for Authentication methods
type Auth struct {
	*deps.Deps
	// Reset password email event.
	resetPassword events.Dispatcher
	// The function used for hashing passwords.
	hashPasswordFunc func(password string) (string, error)
	// The function used for generating tokens.
	generateTokenFunc func(email string) (string, error)
}

// New
//
// Creates a new auth handler.
func New(d *deps.Deps) *Auth {
	return &Auth{
		Deps:              d,
		resetPassword:     events.NewResetPassword(d),
		hashPasswordFunc:  encryption.HashPassword,
		generateTokenFunc: encryption.GenerateEmailToken,
	}
}
