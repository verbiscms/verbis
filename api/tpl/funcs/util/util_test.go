// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ns = New(&deps.Deps{})
)

type noStringer struct{}

func TestNamespace_Len(t *testing.T) {
	ptr := "hello"

	tt := map[string]struct {
		input interface{}
		want  int64
	}{
		"Slice": {
			[]int{1, 2, 3, 4, 5},
			5,
		},
		"Slice Pointer": {
			&[]int{1, 2, 3, 4, 5},
			5,
		},
		"Array": {
			[5]int{1, 2, 3, 4, 5},
			5,
		},
		"Array Pointer": {
			&[5]int{1, 2, 3, 4, 5},
			5,
		},
		"String": {
			"hello",
			5,
		},
		"String Pointer": {
			&ptr,
			5,
		},
		"Map": {
			map[string]string{
				"hello": "hello", "hello!": "hello",
			}, 2,
		},
		"Error": {
			123,
			0,
		},
		"Nil": {
			nil,
			0,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Len(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Explode(t *testing.T) {

	var str []string

	tt := map[string]struct {
		delim interface{}
		text  interface{}
		want  []string
	}{
		"Spaces": {
			delim: " ",
			text:  "hello world !",
			want:  []string{"hello", "world", "!"},
		},
		"Commas": {
			delim: ",",
			text:  "hello,world,!",
			want:  []string{"hello", "world", "!"},
		},
		"Int": {
			delim: "",
			text:  123,
			want:  []string{"1", "2", "3"},
		},
		"No Stringer Delim": {
			delim: noStringer{},
			text:  "hello,world,!",
			want:  str,
		},
		"No Stringer Text": {
			delim: " ",
			text:  noStringer{},
			want:  str,
		},
		"Length": {
			delim: "hello",
			text:  ",",
			want:  []string{"hello"},
		},
		"Nil Delim": {
			delim: nil,
			text:  ",",
			want:  []string{","},
		},
		"Nil Text": {
			delim: "hello",
			text:  nil,
			want:  []string{"h", "e", "l", "l", "o"},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Explode(test.delim, test.text)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Implode(t *testing.T) {
	tt := map[string]struct {
		glue  interface{}
		slice interface{}
		want  string
	}{
		"Spaces": {
			" ",
			[]string{"a", "b", "c"},
			"a b c",
		},
		"Commas": {
			",",
			[]string{"a", "b", "c"},
			"a,b,c",
		},
		"Int": {
			",",
			[]int{1, 2, 3},
			"1,2,3",
		},
		"No Stringer Glue": {
			noStringer{},
			[]string{"a", "b", "c"},
			"",
		},
		"No Stringer Slice": {
			" ",
			noStringer{},
			"",
		},
		"Nil Glue": {
			nil,
			",",
			"",
		},
		"Nil Text": {
			"hello",
			nil,
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Implode(test.glue, test.slice)
			assert.Equal(t, test.want, got)
		})
	}
}
