// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/webp"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	f, err := os.Open("source.webp")
	assert.Nil(t, err)
	imgSource, err := Decode(f)
	assert.Nil(t, err)
	f.Seek(0, 0)
	imgTarget, err := webp.Decode(f)
	assert.Nil(t, err)
	assert.Equal(t, imgSource.Bounds(), imgTarget.Bounds())
}
