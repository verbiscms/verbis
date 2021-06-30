// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"mime/multipart"
	"testing"
)

const TestPNG = "gopher.png"

func TestPNG_Encode(t *testing.T) {
	UtilTestEncode(&PNG{}, t)
}

func TestPNG_Decode(t *testing.T) {
	UtilTestDecode(func(file multipart.File) Imager {
		return &PNG{File: file}
	}, TestPNG, t)
}
