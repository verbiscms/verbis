// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/store/posts"
	"github.com/gin-gonic/gin"
)

// Handler defines methods for categories to interact with the server.
type Handler interface {
	List(ctx *gin.Context)
	Find(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// Categories defines the handler for all category routes.
type Categories struct {
	*deps.Deps
}

// clearCache
//
// TODO: This needs to be in a model, or the cache store.
//
// Clear the post cache that have the given category ID
// attached to it.
func (c *Categories) clearCache(id int) {
	go func() {
		p, _, err := c.Store.Posts.List(params.Params{LimitAll: true}, false, posts.ListConfig{})
		if err != nil {
			logger.WithError(err).Error()
		}
		cache.ClearCategoryCache(id, p)
	}()
}
