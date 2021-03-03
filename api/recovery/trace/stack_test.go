// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Append(t *testing.T) {
	tt := map[string]struct {
		stack Stack
		input *File
		want  Stack
	}{
		"Nil Length": {
			nil,
			&File{File: "test", Line: 1, Function: "function", Contents: "contents"},
			Stack{
				&File{File: "test", Line: 1, Function: "function", Contents: "contents"},
			},
		},
		"Multiple": {
			Stack{
				&File{File: "test", Line: 1, Function: "function", Contents: "contents"},
			},
			&File{File: "test", Line: 2, Function: "function", Contents: "contents"},
			Stack{
				&File{File: "test", Line: 1, Function: "function", Contents: "contents"},
				&File{File: "test", Line: 2, Function: "function", Contents: "contents"},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			test.stack.Append(test.input)
			assert.Equal(t, test.want, test.stack)
		})
	}
}

func TestStack_Prepend(t *testing.T) {
	tt := map[string]struct {
		stack Stack
		input *File
		want  Stack
	}{
		"Nil Length": {
			nil,
			&File{File: "test", Line: 1, Function: "function", Contents: "contents"},
			Stack{
				&File{File: "test", Line: 1, Function: "function", Contents: "contents"},
			},
		},
		"Multiple": {
			Stack{
				&File{File: "test", Line: 1, Function: "function", Contents: "contents"},
			},
			&File{File: "test", Line: 2, Function: "function", Contents: "contents"},
			Stack{
				&File{File: "test", Line: 2, Function: "function", Contents: "contents"},
				&File{File: "test", Line: 1, Function: "function", Contents: "contents"},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			test.stack.Prepend(test.input)
			assert.Equal(t, test.want, test.stack)
		})
	}
}

func TestStack_Find(t *testing.T) {
	f := &File{File: "test", Line: 1, Function: "function", Contents: "contents"}

	tt := map[string]struct {
		stack Stack
		input string
		want  *File
	}{
		"Found": {
			Stack{f},
			"function",
			f,
		},
		"Not Found": {
			nil,
			"function",
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.stack.Find(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_Language(t *testing.T) {
	tt := map[string]struct {
		input string
		want  string
	}{
		"Default":  {"test", "handlebars"},
		"Go":       {".go", "go"},
		"HTML":     {".html", "handlebars"},
		"CMS":      {".cms", "handlebars"},
		"Assembly": {".s", "assembly"},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Language(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_GetStack(t *testing.T) {
	tt := map[string]struct {
		depth    int
		traverse int
		want     int
	}{
		"Single":   {1, 0, 1},
		"Multiple": {3, 0, 3},
		"Traverse": {3, 1, 2},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := New().Trace(test.depth, test.traverse)
			assert.Equal(t, test.want, len(got))
		})
	}
}

func TestFile_Vendor(t *testing.T) {
	tt := map[string]struct {
		input File
		want  bool
	}{
		"Handlebars": {
			File{Language: "handlebars"},
			false,
		},
		"Non Vendor": {
			File{Function: "verbis/recovery"},
			false,
		},
		"Vendor": {
			File{Function: "wrongval"},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Vendor()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestFile_Lines(t *testing.T) {
	tt := map[string]struct {
		input Stack
		want  []*FileLine
	}{
		"Single": {
			Stack{
				{Contents: "test"},
			},
			[]*FileLine{
				{Line: 1, Content: "test"},
			},
		},
		"Multiple": {
			Stack{
				{Contents: "test\ntest"},
			},
			[]*FileLine{
				{Line: 1, Content: "test"},
				{Line: 2, Content: "test"},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			if len(test.input) == 0 {
				t.Error("Wrong args for input")
			}
			got := test.input[0].Lines()
			for i := 0; i < len(got); i++ {
				assert.Equal(t, test.want[i], got[i])
			}
		})
	}
}
