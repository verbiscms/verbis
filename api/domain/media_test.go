// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedia_PossibleFiles(t *testing.T) {
	uniq := uuid.New()
	uniq2 := uuid.New()

	tt := map[string]struct {
		input Media
		want  []string
	}{
		"No Sizes": {
			Media{File: File{UUID: uniq, Path: "2020/01", Name: "file.jpg"}},
			[]string{
				"2020/01/" + uniq.String() + ".jpg",
				"2020/01/" + uniq.String() + ".jpg.webp",
			},
		},
		"With Sizes": {
			Media{File: File{Name: "file.jpg", UUID: uniq, Path: "2020/01"}, Sizes: MediaSizes{
				"test": MediaSize{File: File{Name: "size.jpg", UUID: uniq2, Path: "2020/01"}},
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
			got := test.input.PossibleFiles("")
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMediaSizes_Scan(t *testing.T) {
	UtilTestScanner(MediaSizes{}, t)
}

func TestMediaSizes_Value(t *testing.T) {
	UtilTestValue(MediaSizes{
		"test": MediaSize{Name: "name"},
	}, t)
	UtilTestValueNil(MediaSizes{}, t)
}
