// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gin-gonic/gin"
	"github.com/nickalie/go-webpbin"
	"io/ioutil"
	"os"
	"strings"
)

// Execer defines methods for WebP images to Convert and
// Obtain files from the file system.
type Execer interface {
	Accepts(ctx *gin.Context) bool
	File(g *gin.Context, path string, mime domain.Mime) ([]byte, error)
	Convert(path string, compression int)
}

// Path defines the path where the executables reside.
const Path = string(os.PathSeparator) + "webp"

// WebP
//
// Defines the service for installing, converting, serving
// and determining if the browser can accept WebP
// images.
type WebP struct {
	binPath string
}

// New
//
// Creates a new WebP Execer.
func New(binPath string) *WebP {
	return &WebP{
		binPath: binPath,
	}
}

const (
	// Header is the WebP web header to look for.
	Header = "image/webp"
)

// Accepts
//
// Determines if the browser can serve webp images by
// looking up the 'image/WebP' header.
func (w *WebP) Accepts(ctx *gin.Context) bool {
	acceptHeader := ctx.Request.Header.Get("Accept")
	return strings.Contains(acceptHeader, "image/webp")
}

// File
//
// File first checks to see if the browser accepts WebP
// images and if the mime type is jpg or a png.
//
// Returns the data file if found.
// Returns errors.INTERNAL if no file was found.
func (w *WebP) File(ctx *gin.Context, path string, mime domain.Mime) ([]byte, error) {
	const op = "WebP.Install"

	if !w.Accepts(ctx) && !mime.CanResize() {
		return nil, nil
	}

	p := path + domain.WebPExtension
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error finding file with the path: " + p, Operation: op, Err: err}
	}

	return data, nil
}

// Convert
//
// Converts an image to WebP based on compression and
// decoded image. Compression level is also set.
func (w *WebP) Convert(path string, compression int) {
	const op = "Webp.Convert"

	webpbin.Dest(w.binPath)

	err := webpbin.NewCWebP().
		Quality(uint(compression)).
		InputFile(path).
		OutputFile(path + ".webp").
		Run()

	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error convert the image to WebP", Operation: op, Err: err}).Error()
	}

	logger.Debug("Saved WebP file with the path: " + path)
}
