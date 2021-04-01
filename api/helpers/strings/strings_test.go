// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInSlice(t *testing.T) {
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

func TestBetween(t *testing.T) {
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

func TestAddSpace(t *testing.T) {
	got := AddSpace("HelloWorld")
	want := "Hello World"
	assert.Equal(t, want, got)
}

func TestRandom(t *testing.T) {
	tt := map[string]struct {
		len  int64
		want int
	}{
		"Valid": {
			5,
			5,
		},
		"Valid 2": {
			10,
			10,
		},
		"Valid 3": {
			100,
			100,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Random(test.len, false)
			assert.Equal(t, test.len, int64(len(got)))
		})
	}
}
