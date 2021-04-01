// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InSlice(t *testing.T) {
	tt := map[string]struct {
		input string
		list  []string
		want  bool
	}{
		"Truthy": {
			"test",
			[]string{"test"},
			true,
		},
		"Falsey": {
			"test",
			[]string{""},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := InSlice(test.input, test.list)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_Between(t *testing.T) {
	tt := map[string]struct {
		input string
		a     string
		b     string
		want  string
	}{
		"Simple": {
			"helloverbis",
			"ello",
			"erbis",
			"v",
		},
		"A Not Found": {
			"helloverbis",
			"wrong",
			"",
			"",
		},
		"B Not Found": {
			"helloverbis",
			"",
			"wrong",
			"",
		},
		"Empty": {
			"",
			"",
			"",
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Between(test.input, test.a, test.b)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_AddSpace(t *testing.T) {
	got := AddSpace("HelloWorld")
	want := "Hello World"
	assert.Equal(t, want, got)
}
