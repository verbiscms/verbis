// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"bytes"
	"github.com/verbiscms/verbis/api/errors"
	"image"
	"image/jpeg"
	"mime/multipart"
)

// JPG is an implementation of an Imager to encode
// and decode an image.
type JPG struct {
	File multipart.File
}

// Encode transforms a JPG to a bytes.Buffer with the
// given compression ration.
// Returns errors.INTERNAL if the JPG could not be encoded.
func (j *JPG) Encode(img image.Image, comp int) (*bytes.Buffer, error) {
	const op = "JPG.Encode"

	opts := &jpeg.Options{
		Quality: comp,
	}

	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, opts)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: ErrEncodeMessage, Operation: op, Err: err}
	}

	return buf, nil
}

// Decode Seeks the file and decodes the file into a new
// JPG type.
// Returns errors.INTERNAL if there was an error decoding.
func (j *JPG) Decode() (image.Image, error) {
	const op = "JPG.Decode"

	_, err := j.File.Seek(0, 0)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: ErrSeekFileMessage, Operation: op, Err: err}
	}

	file, err := jpeg.Decode(j.File)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: ErrDecodeMessage, Operation: op, Err: err}
	}

	return file, nil
}
