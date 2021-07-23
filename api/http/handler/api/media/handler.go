// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/services/media"
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
	service media.Library
}

// New
//
// Creates a new media handler.
func New(d *deps.Deps) *Media {
	return &Media{
		Deps:    d,
		service: media.New(d.Options, d.Storage, d.Store.Media),
	}
}
