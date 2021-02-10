// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

func (t *RecoverTestSuite) TestStack_Append() {

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
		t.Run(name, func() {
			test.stack.Append(test.input)
			t.Equal(test.want, test.stack)
		})
	}
}

func (t *RecoverTestSuite) TestStack_Prepend() {

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
		t.Run(name, func() {
			test.stack.Prepend(test.input)
			t.Equal(test.want, test.stack)
		})
	}
}

func (t *RecoverTestSuite) Test_GetStack() {

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
		t.Run(name, func() {
			got := GetStack(test.depth, test.traverse)
			t.Equal(test.want, len(got))
		})
	}
}

func (t *RecoverTestSuite) TestFile_Lines() {

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
		t.Run(name, func() {
			if len(test.input) == 0 {
				t.Fail("Wrong args for input")
			}
			got := test.input[0].Lines()
			for _, v := range got {
				for _, line := range test.want {
					t.Equal(*line, *v)
				}
			}
		})
	}
}
