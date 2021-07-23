// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"testing"
)

var (
	ns = New(&deps.Deps{})
)

func TestNamespace_Replace(t *testing.T) {
	tt := map[string]struct {
		old  string
		new  string
		src  string
		want interface{}
	}{
		"Valid": {
			"-",
			" ",
			"verbis-cms-is-amazing",
			"verbis cms is amazing",
		},
		"Valid 2": {
			"v",
			"",
			"verbis",
			"erbis",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Replace(test.old, test.new, test.src)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Substr(t *testing.T) {
	tt := map[string]struct {
		str   string
		start interface{}
		end   interface{}
		want  interface{}
	}{
		"Valid": {
			"verbiscms",
			0,
			2,
			"ve",
		},
		"Valid 2": {
			"hello world",
			0,
			5,
			"hello",
		},
		"Strings as Params": {
			"hello world",
			"0",
			"5",
			"hello",
		},
		"Negative Start": {
			"hello world",
			"-1",
			"5",
			"hello",
		},
		"Negative End": {
			"hello world",
			"5",
			"-1",
			" world",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Substr(test.str, test.start, test.end)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Trunc(t *testing.T) {
	tt := map[string]struct {
		str   string
		trunc interface{}
		want  interface{}
	}{
		"Positive": {
			"hello world",
			5,
			"hello",
		},
		"Negative": {
			"hello world",
			-5,
			"world",
		},
		"Strings as Params": {
			"hello world",
			"-5",
			"world",
		},
		"Original": {
			"hello world",
			-1000,
			"hello world",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Trunc(test.str, test.trunc)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Ellipsis(t *testing.T) {
	tt := map[string]struct {
		str  string
		len  interface{}
		want interface{}
	}{
		"Valid": {
			"hello world",
			5,
			"hello...",
		},
		"Valid 2": {
			"hello world this is Verbis CMS",
			11,
			"hello world...",
		},
		"Strings as Params": {
			"hello world",
			"5",
			"hello...",
		},
		"Short String": {
			"cms",
			3,
			"cms",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Ellipsis(test.str, test.len)
			assert.Equal(t, test.want, got)
		})
	}
}
