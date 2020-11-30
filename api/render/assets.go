package render

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/helpers/webp"

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
func (r *Render) Asset(g *gin.Context) (*string, *[]byte, error) {
	const op = "Render.GetAsset"

	url := g.Request.URL.Path

	// Check if the file has been cached
	var cachedFile *[]byte
	var cachedMime *string
	if r.options.CacheServerAssets {
		cachedFile, cachedMime = r.getCachedAsset(url)
		if cachedFile != nil && cachedMime != nil {
			return cachedMime, cachedFile, nil
		}
	}

	// Get the relevant paths
	assetsPath := paths.Theme() + r.theme.AssetsPath
	fileName := strings.Replace(url, "/assets", "", 1)
	mimeType := mime.TypeByExtension(strings.Replace(filepath.Ext(fileName), ".", "", 1))

	file, err := ioutil.ReadFile(assetsPath + fileName)
	if err != nil {
		return nil, nil, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Unable to read the file with the path: %s", assetsPath+fileName), Operation: op, Err: err}
	}

	// Set the cache if options allow
	defer func() {
		go func() {
			if r.options.CacheServerAssets && cachedFile == nil {
				cache.Store.Set(url, &file, cache.RememberForever)
				cache.Store.Set(url+"mimetype", &mimeType, cache.RememberForever)
			}
		}()
	}()

	// Set cache headers
	r.cacher.Cache(g)

	// Check if the serving of webp's is allowed & get the
	// webp images and assign if not nil
	if r.options.MediaServeWebP && webp.Accepts(g) {
		webpFile := webp.GetData(g, assetsPath+fileName, mimeType)
		if webpFile != nil {
			mimeType = "image/webp"
			file = webpFile
		}
	}

	// If the minified file is nil or the err is not empty, serve the original data
	minifiedFile, err := r.minify.MinifyBytes(bytes.NewBuffer(file), mimeType)
	if err != nil || minifiedFile != nil {
		file = minifiedFile
	}

	return &mimeType, &file, nil
}
