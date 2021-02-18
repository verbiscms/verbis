// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redirects

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/gin-gonic/gin"
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