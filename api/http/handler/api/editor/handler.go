// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
)

// Handler defines methods for Roles routes to interact with the server.
type Handler interface {
	List(ctx *gin.Context)
}

// Redirects defines the handler for all roles routes.
type Roles struct {
	*deps.Deps
}

// New
//
// Creates a new roles handler.
func New(d *deps.Deps) *Roles {
	return &Roles{
		Deps: d,
	}
}
