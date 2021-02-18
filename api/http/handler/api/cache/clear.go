// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// Clear
//
// Returns 200 upon cache clearing.
func (c *Cache) Clear(ctx *gin.Context) {
	const op = "CacheHandler.Clear"

	cache.Store.Flush()

	api.Respond(ctx, 200, "Successfully cleared server cache", nil)
}
