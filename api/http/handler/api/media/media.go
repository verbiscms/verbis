// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/services/media"
	"github.com/gin-gonic/gin"
)

// Handler defines methods for media items to interact with the server.
type Handler interface {
	List(g *gin.Context)
	Find(g *gin.Context)
	Upload(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

// Media defines the handler for all media item routes.
type Media struct {
	*deps.Deps
	Service media.Service
}
