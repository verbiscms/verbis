// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
)

// Handler defines methods for categories to interact with the server.
type Handler interface {
	List(ctx *gin.Context)
	Find(ctx *gin.Context)
	UpdateCreate(ctx *gin.Context)
}

// Options defines the handler for all options routes.
type Options struct {
	*deps.Deps
}

// New
//
// Creates a new options handler.
func New(d *deps.Deps) *Options {
	return &Options{
		Deps: d,
	}
}
