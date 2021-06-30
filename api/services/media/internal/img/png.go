// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package img

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
)

// PNG implementation of an Imager.
type PNG struct {
	File multipart.File
}

func (p *PNG) Encode(image image.Image) ([]byte, error) {
	const op = "PNG.encode"

	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, image, nil)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error encoding PNG file", Operation: op, Err: err}
	}

	return buf.Bytes(), nil
}

// Decode
//
// Seeks the file and decodes the file into a new PNG type
// Returns errors.INTERNAL if there was an error decoding.
func (p *PNG) Decode() (image.Image, error) {
	const op = "PNG.decode"

	_, err := p.File.Seek(0, 0)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error seeking PNG file", Operation: op, Err: err}
	}

	file, err := png.Decode(p.File)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error decoding PNG file", Operation: op, Err: err}
	}

	return file, nil
}

// Compression - TODO
func (p *PNG) Compression(comp int) imaging.EncodeOption {
	return imaging.PNGCompressionLevel(png.CompressionLevel(comp))
}
