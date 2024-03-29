// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
)

// Handler defines methods for storage to interact with the server.
type Handler interface {
	Config(ctx *gin.Context)
	Connect(ctx *gin.Context)
	Disconnect(ctx *gin.Context)
	Migrate(ctx *gin.Context)
	Download(ctx *gin.Context)
	ListBuckets(ctx *gin.Context)
	CreateBucket(ctx *gin.Context)
	DeleteBucket(ctx *gin.Context)
}

// Storage defines the handler for all storage routes.
type Storage struct {
	*deps.Deps
}

// New creates a new Storage handler.
func New(d *deps.Deps) *Storage {
	return &Storage{
		Deps: d,
	}
}
