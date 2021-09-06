// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/nickalie/go-webpbin"
	"github.com/verbiscms/verbis/api/common/files"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Execer defines methods for WebP images to Convert and
// Obtain files from the file system.
type Execer interface {
	Accepts(ctx *gin.Context) bool
	File(g *gin.Context, path string, mime domain.Mime) ([]byte, error)
	Convert(in io.Reader, compression int) (*bytes.Reader, error)
}

// Path defines the path where the executables reside.
const Path = "webp"

// WebP defines the service for installing, converting, serving
// and determining if the browser can accept WebP
// images.
type WebP struct {
	binPath string
}

// New creates a new WebP Execer.
func New(binPath string) *WebP {
	const op = "WebP.New"
	path := filepath.Join(binPath, Path)
	if !files.DirectoryExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error creating webp path", Operation: op, Err: err})
		}
	}
	return &WebP{
		binPath: path,
	}
}

const (
	// Header is the WebP web header to look for.
	Header = "image/webp"
)

// Accepts determines if the browser can serve webp images
// by looking up the 'image/WebP' header.
func (w *WebP) Accepts(ctx *gin.Context) bool {
	acceptHeader := ctx.Request.Header.Get("Accept")
	return strings.Contains(acceptHeader, Header)
}

// File first checks to see if the browser accepts WebP
// images and if the mime type is jpg or a png.
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

// Convert converts an image to WebP based on compression and
// decoded image. Compression level is also set.
func (w *WebP) Convert(in io.Reader, compression int) (*bytes.Reader, error) {
	const op = "Webp.Convert"

	webpbin.Dest(w.binPath)

	var buf = &bytes.Buffer{}

	err := webpbin.NewCWebP().
		Quality(uint(compression)).
		Input(in).
		Output(buf).
		Run()

	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error converting image to WebP", Operation: op, Err: err}
	}

	return bytes.NewReader(buf.Bytes()), nil
}
