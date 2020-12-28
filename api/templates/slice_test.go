package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"testing"
)

func Test_Slice(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		tmpl  string
		want  interface{}
	}{
		"String": {
			input: nil,
			tmpl:  `{{ $s := slice "a" "b" "c" }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}`,
			want:  "abc",
		},
		"Int": {
			input: nil,
			tmpl:  `{{ $s := slice 1 2 3 }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}`,
			want:  "123",
		},
		"Float": {
			input: nil,
			tmpl:  `{{ $s := slice 1.1 2.2 3.3 }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}`,
			want:  "1.12.23.3",
		},
		"Mixed": {
			input: nil,
			tmpl:  `{{ $s := slice 1 1.1 "a" }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}`,
			want:  "11.1a",
		},
		"Posts": {
			input: domain.Post{},
			tmpl:  `{{ slice . . . }}`,
			want:  make([]domain.Post, 3),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), test.tmpl, test.want, test.input)
		})
	}
}

func Test_Append(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		tmpl  string
		want  interface{}
	}{
		"String": {
			input: []interface{}{"a", "b", "c"},
			tmpl:  `{{ $s := append . "d" }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}{{ index $s 3 }}`,
			want:  "abcd",
		},
		"Int": {
			input: []interface{}{1, 2, 3},
			tmpl:  `{{ $s := append . "4" }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}{{ index $s 3 }}`,
			want:  "1234",
		},
		"Float": {
			input: []interface{}{1.1, 2.2, 3.3},
			tmpl:  `{{ $s := append . "4.4" }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}{{ index $s 3 }}`,
			want:  "1.12.23.34.4",
		},
		"Mixed": {
			input: []interface{}{1, 1.1, "a"},
			tmpl:  `{{ $s := append . "hello" }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}{{ index $s 3 }}`,
			want:  "11.1ahello",
		},
		"Posts": {
			input: []interface{}{[]interface{}{domain.Post{}}, domain.Post{}},
			tmpl:  `{{ $arr := index . 0 }}{{ $post := index . 1 }}{{ append $arr $post }}`,
			want:  make([]domain.Post, 2),
		},
		"Error": {
			input: "wrongval",
			tmpl:  `{{ append . "hello" }}`,
			want:  "unable to append to slice with type: string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), test.tmpl, test.want, test.input)
		})
	}
}

func Test_Prepend(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		tmpl  string
		want  interface{}
	}{
		"String": {
			input: []interface{}{"a", "b", "c"},
			tmpl:  `{{ $s := prepend . "d" }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}{{ index $s 3 }}`,
			want:  "dabc",
		},
		"Int": {
			input: []interface{}{1, 2, 3},
			tmpl:  `{{ $s := prepend . "4" }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}{{ index $s 3 }}`,
			want:  "4123",
		},
		"Float": {
			input: []interface{}{1.1, 2.2, 3.3},
			tmpl:  `{{ $s := prepend . "4.4" }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}{{ index $s 3 }}`,
			want:  "4.41.12.23.3",
		},
		"Mixed": {
			input: []interface{}{1, 1.1, "a"},
			tmpl:  `{{ $s := prepend . "hello" }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}{{ index $s 3 }}`,
			want:  "hello11.1a",
		},
		"Posts": {
			input: []interface{}{[]interface{}{domain.Post{}}, domain.Post{}},
			tmpl:  `{{ $arr := index . 0 }}{{ $post := index . 1 }}{{ prepend $arr $post }}`,
			want:  make([]domain.Post, 2),
		},
		"Error": {
			input: "wrongval",
			tmpl:  `{{ prepend . "hello" }}`,
			want:  "unable to prepend to slice with type: string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), test.tmpl, test.want, test.input)
		})
	}
}

func Test_First(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"String": {
			input: []interface{}{"a", "b", "c"},
			want:  "a",
		},
		"Int": {
			input: []interface{}{1, 2, 3},
			want:  "1",
		},
		"Float": {
			input: []interface{}{1.1, 2.2, 3.3},
			want:  "1.1",
		},
		"Mixed": {
			input: []interface{}{1, 1.1, "a"},
			want:  "1",
		},
		"Empty": {
			input: []interface{}{},
			want:  "",
		},
		"Error": {
			input: "wrongval",
			want:  "unable to get first element of slice with type: string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), `{{ first . }}`, test.want, test.input)
		})
	}
}

func Test_Last(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"String": {
			input: []interface{}{"a", "b", "c"},
			want:  "c",
		},
		"Int": {
			input: []interface{}{1, 2, 3},
			want:  "3",
		},
		"Float": {
			input: []interface{}{1.1, 2.2, 3.3},
			want:  "3.3",
		},
		"Mixed": {
			input: []interface{}{1, 1.1, "a"},
			want:  "a",
		},
		"Empty": {
			input: []interface{}{},
			want:  "",
		},
		"Error": {
			input: "wrongval",
			want:  "unable to get last element of slice with type: string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), `{{ last . }}`, test.want, test.input)
		})
	}
}

func Test_Reverse(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"String": {
			input: []interface{}{"a", "b", "c"},
			want:  "cba",
		},
		"Int": {
			input: []interface{}{1, 2, 3},
			want:  "321",
		},
		"Float": {
			input: []interface{}{1.1, 2.2, 3.3},
			want:  "3.32.21.1",
		},
		"Mixed": {
			input: []interface{}{1, 1.1, "a"},
			want:  "a1.11",
		},
		"Error": {
			input: "wrongval",
			want:  "unable to get reverse slice of type: string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), `{{ $s := reverse . }}{{ index $s 0 }}{{ index $s 1 }}{{ index $s 2 }}`, test.want, test.input)
		})
	}
}
