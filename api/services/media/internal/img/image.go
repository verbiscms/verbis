// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package img

import (
	"bytes"
	"image"
)

// Imager describes the methods for decoding and saving
// different image types such as PNG's and JPG's
type Imager interface {
	Encode(image image.Image, comp int) (*bytes.Buffer, error)
	Decode() (image.Image, error)
}
