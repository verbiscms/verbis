package templates

import "testing"

func Test_Len(t *testing.T) {

	ptr := "hello"

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Slice": {
			input: []int{1, 2, 3, 4, 5},
			want:  5,
		},
		"Slice Pointer": {
			input: &[]int{1, 2, 3, 4, 5},
			want:  5,
		},
		"Array": {
			input: [5]int{1, 2, 3, 4, 5},
			want:  5,
		},
		"Array Pointer": {
			input: &[5]int{1, 2, 3, 4, 5},
			want:  5,
		},
		"String": {
			input: "hello",
			want:  5,
		},
		"String Pointer": {
			input: &ptr,
			want:  5,
		},
		"Map": {
			input: map[string]string{
				"hello": "hello", "hello!": "hello",
			},
			want: 2,
		},
		"Error": {
			input: 123,
			want:  0,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), `{{ len . }}`, test.want, test.input)
		})
	}
}

func Test_Explode(t *testing.T) {

	tt := map[string]struct {
		delim interface{}
		text  interface{}
		want  interface{}
	}{
		"Spaces": {
			delim: " ",
			text:  "hello world !",
			want:  "[hello world !]",
		},
		"Commas": {
			delim: ",",
			text:  "hello,world,!",
			want:  "[hello world !]",
		},
		"Int": {
			delim: "",
			text:  123,
			want:  "[1 2 3]",
		},
		"No Stringer Delim": {
			delim: noStringer{},
			text:  "hello,world,!",
			want:  "[]",
		},
		"No Stringer Text": {
			delim: " ",
			text:  noStringer{},
			want:  "[]",
		},
		"Length": {
			delim: "hello",
			text:  ",",
			want:  "[hello]",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), `{{ explode .Delim .Text }}`, test.want, map[string]interface{}{
				"Delim": test.delim,
				"Text":  test.text,
			})
		})
	}
}

func Test_Implode(t *testing.T) {

	tt := map[string]struct {
		glue  interface{}
		slice interface{}
		want  interface{}
	}{
		"Spaces": {
			glue:  " ",
			slice: []string{"a", "b", "c"},
			want:  "a b c",
		},
		"Commas": {
			glue:  ",",
			slice: []string{"a", "b", "c"},
			want:  "a,b,c",
		},
		"Int": {
			glue:  ",",
			slice: []int{1, 2, 3},
			want:  "1,2,3",
		},
		"No Stringer Glue": {
			glue:  noStringer{},
			slice: []string{"a", "b", "c"},
			want:  "",
		},
		"No Stringer Slice": {
			glue:  " ",
			slice: noStringer{},
			want:  "",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), `{{ implode .Glue .Slice }}`, test.want, map[string]interface{}{
				"Glue":  test.glue,
				"Slice": test.slice,
			})
		})
	}
}
