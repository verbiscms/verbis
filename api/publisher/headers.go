// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
)

// Cacher represents the cache function to set cache headers.
type headerWriter interface {
	Cache(g *gin.Context)
}

// Headers represents the the header struct for setting gin headers
// for frontend caching.
type headers struct {
	options domain.Options
}

// NewCache - Construct
func newHeaders(o domain.Options) *headers {
	return &headers{
		options: o,
	}
}

// Cache
//
// Returns if the asset is with path of admin or the caching
// is disabled in the options.
// Sets the gin headers if extensions are allowed.
func (c *headers) Cache(g *gin.Context) {
	const op = "Cacher.Cache"

	// Bail if the cache frontend is disabled
	if !c.options.CacheFrontend {
		return
	}

	path := g.Request.URL.Path

	// Get the expiration
	expiration := c.options.CacheFrontendSeconds

	// Get the request type
	request := c.options.CacheFrontendRequest
	allowedRequest := []string{"max-age", "max-stale", "min-fresh", "no-cache", "no-store", "no-transform", "only-if-cached"}
	if request == "" || !helpers.StringInSlice(request, allowedRequest) {
		request = "max-age"
	}

	// Get the extensions to be cached
	extensionsAllowed := c.options.CacheFrontendExtension
	extension := filepath.Ext(path)

	// Check if the extensions
	if len(extensionsAllowed) > 0 {
		for _, v := range extensionsAllowed {
			if extension == "."+v {
				cache := ""
				if request == "max-age" || request == "min-fresh" || request == "max-stale" {
					cache = fmt.Sprintf("%s=%s, %s", request, strconv.FormatInt(expiration, 10), "public")
				} else {
					cache = request
				}
				g.Header("Cache-Control", cache)
			}
		}
	}
}
