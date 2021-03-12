// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
	"github.com/spf13/cast"
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
	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Success": {
			[]byte(`{"large": {"url": "/test"}}`),
			nil,
		},
		"Bad Unmarshal": {
			[]byte(`{"large": wrong}`),
			"Error unmarshalling into MediaSize",
		},
		"Nil": {
			nil,
			MediaSizes{},
		},
		"Unsupported Scan": {
			"wrong",
			"Scan unsupported for MediaSize",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := MediaSizes{}
			err := m.Scan(test.input)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Nil(t, err)
		})
	}
}

func TestMediaSizes_Value(t *testing.T) {
	tt := map[string]struct {
		input MediaSizes
		want  interface{}
	}{
		"Success": {
			MediaSizes{
				"test": MediaSize{Url: "/test"},
			},
			MediaSizes{
				"test": MediaSize{Url: "/test"},
			},
		},
		"Nil Length": {
			nil,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			value, _ := test.input.Value()

			if test.input == nil {
				assert.Nil(t, value)
				return
			}

			got, err := cast.ToStringE(value)
			assert.NoError(t, err)

			want, err := json.Marshal(test.input)
			assert.NoError(t, err)

			assert.Equal(t, string(want), got)
		})
	}
}
