// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/gin-gonic/gin"
)

// CacheHandler defines methods for fields to interact with the server
type CacheHandler interface {
	Clear(g *gin.Context)
}

// CacheController defines the handler for Cache
type Cache struct{}

// newCache - Construct
func NewCache() *Cache {
	return &Cache{}
}

// Clear server cache
func (c *Cache) Clear(g *gin.Context) {
	const op = "CacheHandler.Clear"
	cache.Store.Flush()
	Respond(g, 200, "Successfully cleared server cache", nil)
}
