// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/publisher"
	"github.com/gin-gonic/gin"
)

// Handler defines methods for the frontend to
// interact with the server.
type Handler interface {
	Uploads(ctx *gin.Context)
	Assets(ctx *gin.Context)
	Serve(ctx *gin.Context)
	Screenshot(ctx *gin.Context)
}

// Public defines the handler for all public routes.
type Public struct {
	*deps.Deps
	Publisher publisher.Publisher
}
