// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
)

// Resizer describes the method for resizing images for
// the library.
type Resizer interface {
	Resize(imager Imager, dest string, media domain.MediaSize, comp int) error
}

// Imager describes the methods for decoding and saving
// different image types such as PNG's and JPG's
type Imager interface {
	Decode() (image.Image, error)
	Save(image image.Image, path string, comp int) error
}

// Resize implements the Resizer interface.
type Resize struct{}

// Resize satisfies the Resizer by decoding, cropping and
// resizing and finally saving the resized image.
func (i *Resize) Resize(imager Imager, dest string, media domain.MediaSize, comp int) error {
	img, err := imager.Decode()
	if err != nil {
		return err
	}

	var resized *image.NRGBA
	if media.Crop {
		resized = imaging.Fill(img, media.Width, media.Height, imaging.Center, imaging.Lanczos)
	} else {
		resized = imaging.Resize(img, media.Width, media.Height, imaging.Lanczos)
	}

	err = imager.Save(resized, dest, comp)
	if err != nil {
		return err
	}

	return nil
}

// PNG implementation of an Imager.
type PNG struct {
	File multipart.File
}

// Decode
//
// Seeks the file and decodes the file into a new PNG type
// Returns errors.INTERNAL if there was an error decoding.
func (p *PNG) Decode() (image.Image, error) {
	const op = "PNG.Decode"
	_, err := p.File.Seek(0, 0)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error seeking PNG file", Operation: op, Err: err}
	}
	file, err := png.Decode(p.File)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error decoding PNG", Operation: op, Err: err}
	}
	return file, nil
}

// Save
//
// Accepts an image.Image, a path for where to save the
// image and a compression level.
// Returns errors.INTERNAL if the file could not be saved.
func (p *PNG) Save(img image.Image, path string, comp int) error {
	const op = "PNG.Save"
	err := imaging.Save(img, path, imaging.PNGCompressionLevel(png.CompressionLevel(comp)))
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error saving the PNG file with the path: " + path, Operation: op, Err: err}
	}
	return nil
}

// JPG implementation of an Imager.
type JPG struct {
	File        multipart.File
	Compression int
}

// Decode
//
// Seeks the file and decodes the file into a new PNG type
// Returns errors.INTERNAL if there was an error decoding.
func (j *JPG) Decode() (image.Image, error) {
	const op = "JPG.Decode"
	_, err := j.File.Seek(0, 0)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error seeking JPG file", Operation: op, Err: err}
	}
	file, err := jpeg.Decode(j.File)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error decoding JPG", Operation: op, Err: err}
	}
	return file, nil
}

// Save
//
// Accepts an image.Image, a path for where to save the
// image and a compression level.
// Returns errors.INTERNAL if the file could not be saved.
func (j *JPG) Save(img image.Image, path string, comp int) error {
	const op = "JPG.Save"
	err := imaging.Save(img, path, imaging.JPEGQuality(comp))
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error saving the JPG file with the path: " + path, Operation: op, Err: err}
	}
	return nil
}
