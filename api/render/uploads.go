// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/helpers/webp"
	"github.com/gin-gonic/gin"
)

func (r *Render) Upload(g *gin.Context) (*string, *[]byte, error) {
	const op = "Render.Upload"

	url := g.Request.URL.Path

	api.UploadChan <- 1
	defer func() {
		<-api.UploadChan
	}()

	// Check if the file has been cached
	var cachedFile *[]byte
	var cachedMime *string
	if r.options.CacheServerUploads {
		cachedFile, cachedMime = r.getCachedAsset(url)
		if cachedFile != nil && cachedMime != nil {
			return cachedMime, cachedFile, nil
		}
	}

	// Set cache headers
	r.cacher.Cache(g)

	// Get the data & mime type from the media store
	file, mimeType, err := r.store.Media.Serve(url, webp.Accepts(g))
	if err != nil {
		return nil, nil, err
	}

	// Set the cache if the app is in production
	defer func() {
		go func() {
			if r.options.CacheServerUploads && cachedFile == nil {
				cache.Store.Set(url, &file, cache.RememberForever)
				cache.Store.Set(url+"mimetype", &mimeType, cache.RememberForever)
			}
		}()
	}()

	// If the minified file is nil or the err is not empty, serve the original data
	minifiedFile, err := r.minify.MinifyBytes(bytes.NewBuffer(file), mimeType)
	if err != nil || minifiedFile != nil {
		file = minifiedFile
	}

	return &mimeType, &file, nil
}
