// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelative_Tags(t *testing.T) {
	tt := map[string]struct {
		input TagSlice
		want  string
	}{
		"Success": {
			TagSlice{"noopener", "nofollow"},
			"noopener nofollow",
		},
		"Empty": {
			TagSlice{},
			"",
		},
		"With Empty Field": {
			TagSlice{"nofollow", ""},
			"nofollow",
		},
		"With Space": {
			TagSlice{"nofollow", " "},
			"nofollow",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Tags()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestRelative_HasTags(t *testing.T) {
	tt := map[string]struct {
		input TagSlice
		want  bool
	}{
		"True": {
			TagSlice{"test"},
			true,
		},
		"False": {
			TagSlice{},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.HasTags()
			assert.Equal(t, test.want, got)
		})
	}
}
