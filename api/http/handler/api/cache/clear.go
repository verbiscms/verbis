// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
)

// Clear
//
// Returns http.StatusOK upon cache clearing.
func (c *Cache) Clear(ctx *gin.Context) {
	const op = "CacheHandler.Clear"

	cache.Store.Flush()

	api.Respond(ctx, http.StatusOK, "Successfully cleared server cache", nil)
}
