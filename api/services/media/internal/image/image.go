// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"bytes"
	"image"
)

// Imager describes the methods for encoding and decoding
// different image types such as PNGs and JPGs.
type Imager interface {
	Encode(image image.Image, comp int) (*bytes.Buffer, error)
	Decode() (image.Image, error)
}

const (
	// ErrEncodeMessage is returned by the Imager when there
	// was an error encoding the image.
	ErrEncodeMessage = "Error encoding image"
	// ErrSeekFileMessage is returned by the Imager when there
	// was an error seeking the file when encoding.
	ErrSeekFileMessage = "Error seeking file"
	// ErrDecodeMessage is returned by the Imager when there
	// was an error encoding the image.
	ErrDecodeMessage = "Error decoding image"
)
