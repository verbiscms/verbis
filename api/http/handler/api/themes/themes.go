// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/gin-gonic/gin"
)

// Handler defines methods for the themes to interact with the server.
type Handler interface {
	List(ctx *gin.Context)
	Find(ctx *gin.Context)
	Config(ctx *gin.Context)
	Templates(ctx *gin.Context)
	Layouts(ctx *gin.Context)
	Update(ctx *gin.Context)
}

// Themes defines the handler for all site routes.
type Themes struct {
	*deps.Deps
}
