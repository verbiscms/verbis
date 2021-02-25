// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/helpers/webp"
	"github.com/gin-gonic/gin"
)

func (r *publish) Upload(g *gin.Context) (string, *[]byte, error) {
	const op = "publish.Upload"

	api.UploadChan <- 1
	defer func() {
		<-api.UploadChan
	}()

	url := g.Request.URL.Path

	// Set cache headers
	r.cacher.Cache(g)

	// Get the data & mime type from the media store
	file, mimeType, err := r.Store.Media.Serve(url, webp.Accepts(g))
	if err != nil {
		return "", nil, err
	}

	// If the minified file is nil or the err is not empty, serve the original data
	//minifiedFile, err := r.minify.MinifyBytes(bytes.NewBuffer(file), mimeType)
	//if err != nil {
	//	return mimeType, &file, nil
	//}

	return mimeType, &file, nil
}
