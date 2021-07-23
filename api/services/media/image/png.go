// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"bytes"
	"github.com/verbiscms/verbis/api/errors"
	"image"
	"image/png"
	"mime/multipart"
)

// PNG is an implementation of an Imager to encode
// and decode an image.
type PNG struct {
	File multipart.File
}

// Encode transforms a PNG to a bytes.Buffer with the
// given compression ration.
// Returns errors.INTERNAL if the PNG could not be encoded.
func (p *PNG) Encode(img image.Image, comp int) (*bytes.Buffer, error) {
	const op = "PNG.Encode"

	enc := &png.Encoder{
		CompressionLevel: png.CompressionLevel(comp),
	}

	buf := new(bytes.Buffer)
	err := enc.Encode(buf, img)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: ErrEncodeMessage, Operation: op, Err: err}
	}

	return buf, nil
}

// Decode Seeks the file and decodes the file into a new
// PNG type.
// Returns errors.INTERNAL if there was an error decoding.
func (p *PNG) Decode() (image.Image, error) {
	const op = "PNG.Decode"

	_, err := p.File.Seek(0, 0)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: ErrSeekFileMessage, Operation: op, Err: err}
	}

	file, err := png.Decode(p.File)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: ErrDecodeMessage, Operation: op, Err: err}
	}

	return file, nil
}
