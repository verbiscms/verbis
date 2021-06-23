// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMedia_UploadPath(t *testing.T) {
	uniq := uuid.New()

	tt := map[string]struct {
		input Media
		want  interface{}
	}{
		"Date": {
			Media{FilePath: "2020/01", UUID: uniq, Url: "file.jpg"},
			"2020/01" + string(os.PathSeparator) + uniq.String() + ".jpg",
		},
		"No Date": {
			Media{FilePath: "", UUID: uniq, Url: "file.jpg"},
			uniq.String() + ".jpg",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.UploadPath()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMedia_IsOrganiseYearMonth(t *testing.T) {
	tt := map[string]struct {
		input Media
		want  bool
	}{
		"With Date": {
			Media{FilePath: "2020/01"},
			true,
		},
		"Empty": {
			Media{FilePath: ""},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.IsOrganiseYearMonth()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMedia_Extension(t *testing.T) {
	tt := map[string]struct {
		input Media
		want  interface{}
	}{
		"JPG": {
			Media{Url: "file.jpg"},
			".jpg",
		},
		"PNG": {
			Media{Url: "file.png"},
			".png",
		},
		"SVG": {
			Media{Url: "file.svg"},
			".svg",
		},
		"PDF": {
			Media{Url: "file.pdf"},
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

func TestMedia_PossibleFiles(t *testing.T) {
	uniq := uuid.New()
	uniq2 := uuid.New()

	tt := map[string]struct {
		input Media
		want  []string
	}{
		"No Sizes": {
			Media{Url: "file.jpg", UUID: uniq, FilePath: "2020/01"},
			[]string{
				"2020/01/" + uniq.String() + ".jpg",
				"2020/01/" + uniq.String() + ".jpg.webp",
			},
		},
		"With Sizes": {
			Media{Url: "file.jpg", UUID: uniq, FilePath: "2020/01", Sizes: MediaSizes{
				"Test": MediaSize{
					UUID: uniq2,
				},
			}},
			[]string{
				"2020/01/" + uniq.String() + ".jpg",
				"2020/01/" + uniq.String() + ".jpg.webp",
				"2020/01/" + uniq2.String() + ".jpg",
				"2020/01/" + uniq2.String() + ".jpg.webp",
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.PossibleFiles()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMediaSizes_Scan(t *testing.T) {
	UtilTestScanner(MediaSizes{}, t)
}

func TestMediaSizes_Value(t *testing.T) {
	UtilTestValue(MediaSizes{
		"test": MediaSize{Url: "/test"},
	}, t)
	UtilTestValueNil(MediaSizes{}, t)
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
