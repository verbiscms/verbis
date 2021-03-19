// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gin-gonic/gin"
	"github.com/nickalie/go-webpbin"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"strings"
)

// Execer defines methods for WebP images to Convert and
// Obtain files from the file system.
type Execer interface {
	Install() error
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

// Install
//
// Installs the WebP executables to the path provided New.
// Returns errors.INTERNAL if an error occurred
// downloading.
func (w *WebP) Install() error {
	const op = "WebP.Install"

	width := 200
	height := 100

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: width, Y: height},
	})

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width && y < height:
				img.Set(x, y, color.RGBA{R: 100, G: 200, B: 200, A: 0xff})
			}
		}
	}

	webpbin.Dest(w.binPath)

	b := &bytes.Buffer{}
	err := webpbin.Encode(b, img)

	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error downloading WebP executables", Operation: op, Err: err}
	}

	return nil
}

// Accepts
//
// Determines if the browser can serve webp images by
// looking up the 'image/WebP' header.
func (w *WebP) Accepts(ctx *gin.Context) bool {
	acceptHeader := ctx.Request.Header.Get("Accept")
	return strings.Contains(acceptHeader, "image/WebP")
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
		Quality(uint(100 - compression)). //nolint
		InputFile(path).
		OutputFile(path + ".WebP").
		Run()

	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error convert the image to WebP", Operation: op, Err: err}).Error()
	}
}
