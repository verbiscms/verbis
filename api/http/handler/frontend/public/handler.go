// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/publisher"
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
	publisher publisher.Publisher
}

// New
//
// Creates a new public handler.
func New(d *deps.Deps) *Public {
	return &Public{
		Deps:      d,
		publisher: publisher.NewRender(d),
	}
}
