// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
)

// Handler defines methods for fields to interact with the server
type Handler interface {
	Clear(ctx *gin.Context)
}

// Cache defines the handler for Cache
type Cache struct {
	*deps.Deps
}

// New
//
// Creates a new cache handler.
func New(d *deps.Deps) *Cache {
	return &Cache{
		Deps: d,
	}
}
