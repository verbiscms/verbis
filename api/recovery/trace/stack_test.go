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
			&File{File: "test", Line: 1, Name: "name", Contents: "contents"},
			Stack{
				&File{File: "test", Line: 1, Name: "name", Contents: "contents"},
			},
		},
		"Multiple": {
			Stack{
				&File{File: "test", Line: 1, Name: "name", Contents: "contents"},
			},
			&File{File: "test", Line: 2, Name: "name", Contents: "contents"},
			Stack{
				&File{File: "test", Line: 1, Name: "name", Contents: "contents"},
				&File{File: "test", Line: 2, Name: "name", Contents: "contents"},
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
			&File{File: "test", Line: 1, Name: "name", Contents: "contents"},
			Stack{
				&File{File: "test", Line: 1, Name: "name", Contents: "contents"},
			},
		},
		"Multiple": {
			Stack{
				&File{File: "test", Line: 1, Name: "name", Contents: "contents"},
			},
			&File{File: "test", Line: 2, Name: "name", Contents: "contents"},
			Stack{
				&File{File: "test", Line: 2, Name: "name", Contents: "contents"},
				&File{File: "test", Line: 1, Name: "name", Contents: "contents"},
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

	f := &File{File: "test", Line: 1, Name: "name", Contents: "contents"}

	tt := map[string]struct {
		stack Stack
		input string
		want  *File
	}{
		"Found": {
			Stack{f},
			"name",
			f,
		},
		"Not Found": {
			nil,
			"name",
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

func TestFile_Lines(t *testing.T) {

	tt := map[string]struct {
		input Stack
		want  []*FileLine
	}{
		"Single": {
			Stack{
				{Line: 1, Contents: "test\ntest"},
			},
			[]*FileLine{
				{Line: 2, Content: "test"},
			},
		},
		"Multiple": {
			Stack{
				{Line: 1, Contents: "test"},
				{Line: 2, Contents: "test\ntest"},
			},
			[]*FileLine{
				{Line: 1, Content: "test"},
				{Line: 2, Content: "test\ntest"},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			if len(test.input) == 0 {
				t.Error("Wrong args for input")
			}
			got := test.input[0].Lines()
			for _, v := range got {
				for _, line := range test.want {
					assert.Equal(t, *line, *v)
				}
			}
		})
	}
}
