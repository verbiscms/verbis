// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsValidMime(t *testing.T) {
	tt := map[string]struct {
		input   string
		allowed []string
		want    interface{}
	}{
		"Allowed": {
			"image/jpeg",
			[]string{"image/jpeg"},
			true,
		},
		"Invalid": {
			"image/jpeg",
			[]string{""},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := IsValidMime(test.allowed, test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_TypeByExtension(t *testing.T) {
	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"JPG": {
			".jpg",
			"image/jpeg",
		},
		"PNG": {
			".png",
			"image/png",
		},
		"JPG Trimmed": {
			"jpg",
			"image/jpeg",
		},
		"PNG Trimmed": {
			"png",
			"image/png",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := TypeByExtension(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
