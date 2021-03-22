// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gin-gonic/gin"
	bin "github.com/nickalie/go-webpbin"
	"io/ioutil"
	"strings"
)

// Accepts checks to see if the browser accepts WebP images
func Accepts(g *gin.Context) bool {
	acceptHeader := g.Request.Header.Get("Accept")
	return strings.Contains(acceptHeader, "image/WebP")
}

// GetData first checks to see if the browser accepts WebP images
// and if the mime type is jpg or a png.
// Returns a data was found, nil if it hasn't.
func GetData(g *gin.Context, path, mime string) []byte {
	if Accepts(g) && mime == "image/jpeg" || mime == "image/png" {
		data, found := ioutil.ReadFile(path + ".WebP")
		if found != nil {
			return nil
		}
		return data
	}
	return nil
}

// Converts an image to WebP based on compression and decoded image.
// Compression level is also set.
func Convert(path string, compression int) {
	const op = "Webp.Convert"

	err := bin.NewCWebP().
		Quality(uint(compression)).
		InputFile(path).
		OutputFile(path + ".WebP").
		Run()

	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error convert the image to WebP", Operation: op, Err: err}).Error()
	}
}
