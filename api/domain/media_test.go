// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedia_Public(t *testing.T) {
	tt := map[string]struct {
		input Media
		want  interface{}
	}{
		"Converted": {
			Media{
				Title: "title",
				Sizes: MediaSizes{"test": MediaSize{Width: 100, SizeKey: "test"}},
			},
			MediaPublic{
				Title: "title",
				Sizes: MediaSizesPublic{"test": MediaSizePublic{Width: 100}},
			},
		},
		"Nil": {
			Media{
				Title: "title",
			},
			MediaPublic{
				Title: "title",
				Sizes: nil,
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Public()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMediaItems_Public(t *testing.T) {
	m := MediaItems{
		{Title: "title"},
	}
	got := m.Public()
	want := []MediaPublic{
		{Title: "title"},
	}
	assert.Equal(t, want, got)
}

func TestMediaSizes_Scan(t *testing.T) {
	UtilTestScanner(MediaSizes{}, t)
}

func TestMediaSizes_Value(t *testing.T) {
	UtilTestValue(MediaSizes{
		"test": MediaSize{SizeName: "name"},
	}, t)
	UtilTestValueNil(MediaSizes{}, t)
}
