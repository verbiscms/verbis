// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
)

// Handler defines methods for Redirect routes to interact with the server.
type Handler interface {
	List(ctx *gin.Context)
	Find(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// Redirects defines the handler for all redirect routes.
type Redirects struct {
	*deps.Deps
}

// New
//
// Creates a new redirects handler.
func New(d *deps.Deps) *Redirects {
	return &Redirects{
		Deps: d,
	}
}
