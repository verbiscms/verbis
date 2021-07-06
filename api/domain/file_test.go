// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestFile_UploadPath(t *testing.T) {
	uniq := uuid.New()
	prefix := filepath.Join("Users", "verbis")

	tt := map[string]struct {
		prefix string
		input  File
		want   string
	}{
		"Local": {
			prefix,
			File{Name: "file.jpg", Path: "uploads/2020/01", UUID: uniq, Provider: StorageLocal},
			filepath.Join("Users", "verbis", "uploads", "2020", "01", uniq.String()+".jpg"),
		},
		"Remote": {
			prefix,
			File{Name: "file.jpg", Path: "https:/s3-eu-west-2.amazonaws.com/verbis/uploads/2021/07", UUID: uniq, Provider: StorageAWS},
			"https:/s3-eu-west-2.amazonaws.com/verbis/uploads/2021/07/" + uniq.String() + ".jpg",
		},
		"Local Leading Trailing": {
			prefix,
			File{Name: "file.jpg", Path: "/uploads/2020/01/", UUID: uniq, Provider: StorageLocal},
			filepath.Join("Users", "verbis", "uploads", "2020", "01", uniq.String()+".jpg"),
		},
		"Remote Trailing": {
			prefix,
			File{Name: "file.jpg", Path: "https:/s3-eu-west-2.amazonaws.com/verbis/uploads/2021/07/", UUID: uniq, Provider: StorageAWS},
			"https:/s3-eu-west-2.amazonaws.com/verbis/uploads/2021/07/" + uniq.String() + ".jpg",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.PrivatePath(test.prefix)
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
