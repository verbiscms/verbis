// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/gin-gonic/gin"
)

// Handler defines methods for storage to interact with the server.
type Handler interface {
	Update(ctx *gin.Context)
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
