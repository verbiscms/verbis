// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/gin-gonic/gin"
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