// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
)

type Imager interface {
	Decode() (image.Image, error)
	Save(image image.Image, path string, comp int) error
}

// JPG
type PNG struct {
	File multipart.File
}

// Decode
//
//
func (p *PNG) Decode() (image.Image, error) {
	const op = "PNG.Decode"
	file, err := png.Decode(p.File)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error decoding PNG", Operation: op, Err: err}
	}
	return file, nil
}

// Save
//
//
func (p *PNG) Save(img image.Image, path string, comp int) error {
	const op = "PNG.Save"
	err := imaging.Save(img, path, imaging.PNGCompressionLevel(png.CompressionLevel(comp)))
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error saving the PNG file with the path: " + path, Operation: op, Err: err}
	}
	return nil
}

// JPG
type JPG struct {
	File        multipart.File
	Compression int
}

// Decode
//
//
func (j *JPG) Decode() (image.Image, error) {
	const op = "JPG.Decode"
	file, err := jpeg.Decode(j.File)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error decoding JPG", Operation: op, Err: err}
	}
	return file, nil
}

// Save
//
//
func (j *JPG) Save(img image.Image, path string, comp int) error {
	const op = "JPG.Save"
	err := imaging.Save(img, path, imaging.JPEGQuality(comp))
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error saving the PNG file with the path: " + path, Operation: op, Err: err}
	}
	return nil
}
