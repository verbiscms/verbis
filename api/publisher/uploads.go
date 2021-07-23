// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/domain"
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

	// Set cache headers
	r.cacher.Cache(g)

	var (
		bytes []byte
		file  domain.File
		err   error
	)

	path := g.Request.URL.Path
	if webp && r.Options.MediaServeWebP && r.WebP.Accepts(g) {
		bytes, file, err = r.Storage.Find(path + domain.WebPExtension)
		if err == nil {
			return &bytes, file.Mime, nil
		}
	}

	bytes, file, err = r.Storage.Find(path)
	if err != nil {
		return nil, "", err
	}

	// If the minified file is nil or the err is not empty, serve the original data
	//minifiedFile, err := r.minify.MinifyBytes(bytes.NewBuffer(file), mimeType)
	//if err != nil {
	//	return mimeType, &file, nil
	//}

	return &bytes, file.Mime, nil
}
