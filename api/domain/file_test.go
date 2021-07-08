// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFile_ID(t *testing.T) {
	tt := map[string]struct {
		input  File
		prefix string
		want   interface{}
	}{
		"Local": {
			File{BucketId: "uploads/2020/01/test.jpg", Provider: StorageLocal},
			"prefix",
			"prefix/uploads/2020/01/test.jpg",
		},
		"Remote": {
			File{BucketId: "uploads/2020/01/test.jpg"},
			"prefix",
			"uploads/2020/01/test.jpg",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.ID(test.prefix)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestFile_Extension(t *testing.T) {
	tt := map[string]struct {
		input File
		want  interface{}
	}{
		"JPG": {
			File{Name: "file.jpg"},
			".jpg",
		},
		"PNG": {
			File{Name: "file.png"},
			".png",
		},
		"SVG": {
			File{Name: "file.svg"},
			".svg",
		},
		"PDF": {
			File{Name: "file.pdf"},
			".pdf",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Extension()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestUpload_Validate(t *testing.T) {
	tt := map[string]struct {
		input Upload
		want  interface{}
	}{
		"Success": {
			Upload{Path: "path", Size: 1, Contents: strings.NewReader("test"), SourceType: MediaSourceType},
			nil,
		},
		"No Path": {
			Upload{},
			"no path attached to upload",
		},
		"No Size": {
			Upload{Path: "path"},
			"no size attached to upload",
		},
		"Nil Contents": {
			Upload{Path: "path", Size: 1, Contents: nil},
			"upload contents is nil",
		},
		"No Source Type": {
			Upload{Path: "path", Size: 1, Contents: strings.NewReader("test")},
			"no source type attached to upload",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Validate()
			if got != nil {
				assert.Contains(t, got.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestFile_IsLocal(t *testing.T) {
	tt := map[string]struct {
		input File
		want  interface{}
	}{
		"Local": {
			File{Provider: StorageLocal},
			true,
		},
		"Remote": {
			File{Provider: StorageAWS},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.IsLocal()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMime_CanResize(t *testing.T) {
	tt := map[string]struct {
		input string
		want  bool
	}{
		"Jpeg": {
			"image/jpeg",
			true,
		},
		"SVG": {
			"image/svg+xml",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Mime(test.input).CanResize()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMime_IsJPG(t *testing.T) {
	tt := map[string]struct {
		input string
		want  bool
	}{
		"Is Jpeg": {
			"image/jpeg",
			true,
		},
		"Isn't Jpeg": {
			"image/svg+xml",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Mime(test.input).IsJPG()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMime_IsPNG(t *testing.T) {
	tt := map[string]struct {
		input string
		want  bool
	}{
		"Is PNG": {
			"image/png",
			true,
		},
		"Isn't PNG": {
			"image/svg+xml",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Mime(test.input).IsPNG()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMime_String(t *testing.T) {
	got := Mime("image/png").String()
	assert.IsType(t, string("test"), got)
}
