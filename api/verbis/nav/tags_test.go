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
		input Relative
		want  string
	}{
		"Success": {
			Relative{"noopener", "nofollow"},
			"noopener nofollow",
		},
		"Empty": {
			Relative{},
			"",
		},
		"With Empty Field": {
			Relative{"nofollow", ""},
			"nofollow",
		},
		"With Space": {
			Relative{"nofollow", " "},
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
		input Relative
		want  bool
	}{
		"True": {
			Relative{"test"},
			true,
		},
		"False": {
			Relative{},
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
