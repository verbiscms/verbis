// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Asset
//
// It then obtains the assets path from the site model, and then checks
// if the file exists, by opening the file, if it doesnt it will
// return a 404.
// It then sets cache headers using the cacher interface & checks if a webp
// image is available with the path of .jpg.webp. The minify is the used
// to see if the file can be minfied.
func (r *publish) Asset(g *gin.Context) (string, *[]byte, error) {
	const op = "publish.GetAsset"

	api.AssetsChan <- 1
	defer func() {
		<-api.AssetsChan
	}()

	url := g.Request.URL.Path

	// Get the relevant paths
	assetsPath := r.ThemePath() + r.Config.AssetsPath
	fileName := strings.Replace(url, "/assets", "", 1)
	mimeType := mime.TypeByExtension(strings.Replace(filepath.Ext(fileName), ".", "", 1))

	file, err := ioutil.ReadFile(assetsPath + fileName)
	if err != nil {
		return "", nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to read the file with the path: %s", assetsPath+fileName), Operation: op, Err: err}
	}

	// Set cache headers
	r.cacher.Cache(g)

	// Check if the serving of webp's is allowed & get the
	// webp images and assign if not nil
	if r.Options.MediaServeWebP && r.WebP.Accepts(g) {
		webpFile, err := r.WebP.File(g, assetsPath+fileName, domain.Mime(mimeType))
		if err == nil {
			return "image/webp", &webpFile, nil
		}
	}

	// If the minified file is nil or the err is not empty, serve the original data
	//minifiedFile, err := r.minify.MinifyBytes(bytes.NewBuffer(file), mimeType)
	//if err != nil {
	//	return mimeType, &file, nil
	//}

	return mimeType, &file, nil
}
