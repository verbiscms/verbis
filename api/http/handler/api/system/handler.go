// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
)

// Handler defines methods for the system to interact with the server.
type Handler interface {
	Preflight(ctx *gin.Context)
	Install(ctx *gin.Context)
	Update(ctx *gin.Context)
}

// System defines the handler for all system routes.
type System struct {
	*deps.Deps
}

// New
//
// Creates a new site handler.
func New(d *deps.Deps) *System {
	return &System{
		Deps: d,
	}
}
