// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/common/mime"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// AssetsChan is the channel for serving assets for the
// frontend.
var AssetsChan = make(chan int, api.AssetsChannel)

// Asset
//
// It then obtains the assets path from the site model, and then checks
// if the file exists, by opening the file, if it doesnt it will
// return a 404.
// It then sets cache headers using the cacher interface & checks if a webp
// image is available with the path of .jpg.webp. The minify is the used
// to see if the file can be minified.
func (r *publish) Asset(ctx *gin.Context, webp bool) (*[]byte, domain.Mime, error) {
	const op = "publish.GetAsset"

	AssetsChan <- 1
	defer func() {
		<-AssetsChan
	}()

	url := ctx.Request.URL.Path

	// Get the relevant paths
	assetsPath := r.ThemePath() + string(os.PathSeparator) + r.Config.AssetsPath
	fileName := strings.Replace(url, "/assets", "", 1)
	mimeType := mime.TypeByExtension(strings.Replace(filepath.Ext(fileName), ".", "", 1))

	file, err := ioutil.ReadFile(assetsPath + fileName)
	if err != nil {
		return nil, "", &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Error reading the file with the path: %s", assetsPath+fileName), Operation: op, Err: err}
	}

	// Set cache headers
	r.cacher.Cache(ctx)

	// Check if the serving of webp's is allowed & get the
	// webp images and assign if not nil
	if r.Options.MediaServeWebP && r.WebP.Accepts(ctx) && webp {
		webpFile, err := r.WebP.File(ctx, assetsPath+fileName, domain.Mime(mimeType))
		if err == nil {
			return &webpFile, "image/webp", nil
		}
	}

	// If the minified file is nil or the err is not empty, serve the original data
	//minifiedFile, err := r.minify.MinifyBytes(bytes.NewBuffer(file), mimeType)
	//if err != nil {
	//	return mimeType, &file, nil
	//}

	return &file, domain.Mime(mimeType), nil
}
