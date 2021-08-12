// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package editor

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
)

// Handler defines methods for Editor routes to interact with the server.
type Handler interface {
	Preview(ctx *gin.Context)
}

// Editor defines the handler for all roles routes.
type Editor struct {
	*deps.Deps
}

// New
//
// Creates a new editor handler.
func New(d *deps.Deps) *Editor {
	return &Editor{
		Deps: d,
	}
}
