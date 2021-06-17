// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-gonic/gin"
)

// UploadChan is the channel for serving uploads for the
// frontend.
var UploadChan = make(chan int, api.UploadChannel)

func (r *publish) Upload(g *gin.Context, webp bool) (*[]byte, domain.Mime, error) {
	const op = "publish.Upload"

	UploadChan <- 1
	defer func() {
		<-UploadChan
	}()

	url := g.Request.URL.Path

	// Set cache headers
	r.cacher.Cache(g)

	media, path, err := r.Store.Media.FindByURL(url)
	if err != nil {
		return nil, "", err
	}

	acceptsWebP := r.WebP.Accepts(g)
	if !webp {
		acceptsWebP = false
	}

	// Get the data & mime type from the media store
	file, mimeType, err := r.media.Serve(media, path, acceptsWebP)
	if err != nil {
		return nil, "", err
	}

	// If the minified file is nil or the err is not empty, serve the original data
	//minifiedFile, err := r.minify.MinifyBytes(bytes.NewBuffer(file), mimeType)
	//if err != nil {
	//	return mimeType, &file, nil
	//}

	return &file, mimeType, nil
}
