// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"io"
)

// Decode reads a WebP image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	return NewDWebP().Input(r).Run()
}
