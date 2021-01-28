package slice

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ns = New(&deps.Deps{})
)

type noStringer struct{}

func Test_Slice(t *testing.T) {

	tt := map[string]struct {
		input []interface{}
		want  interface{}
	}{
		"String": {
			[]interface{}{"a", "b", "c"},
			[]interface{}{"a", "b", "c"},
		},
		"Int": {
			[]interface{}{1, 2, 3},
			[]interface{}{1, 2, 3},
		},
		"Float": {
			[]interface{}{1.1, 2.2, 3.3},
			[]interface{}{1.1, 2.2, 3.3},
		},
		"Mixed": {
			[]interface{}{"a", 1, 1.1},
			[]interface{}{"a", 1, 1.1},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.slice(test.input...)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_Append(t *testing.T) {

	slice := []string{"a", "b", "c"}

	tt := map[string]struct {
		input interface{}
		slice interface{}
		want  interface{}
	}{
		"String": {
			"d",
			slice,
			[]interface{}{"a", "b", "c", "d"},
		},
		"Int": {
			1,
			slice,
			[]interface{}{"a", "b", "c", 1},
		},
		"Float": {
			1.1,
			slice,
			[]interface{}{"a", "b", "c", 1.1},
		},
		"Error": {
			"a",
			"wrongval",
			"unable to append to slice with type: string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.append(test.slice, test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_Prepend(t *testing.T) {

	slice := []string{"a", "b", "c"}

	tt := map[string]struct {
		input interface{}
		slice interface{}
		want  interface{}
	}{
		"String": {
			"d",
			slice,
			[]interface{}{"d", "a", "b", "c"},
		},
		"Int": {
			1,
			slice,
			[]interface{}{1, "a", "b", "c"},
		},
		"Float": {
			1.1,
			slice,
			[]interface{}{1.1, "a", "b", "c"},
		},
		"Error": {
			"a",
			"wrongval",
			"unable to prepend to slice with type: string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.prepend(test.slice, test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_First(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"String": {
			[]interface{}{"a", "b", "c"},
			"a",
		},
		"Int": {
			[]interface{}{1, 2, 3},
			1,
		},
		"Float": {
			[]interface{}{1.1, 2.2, 3.3},
			1.1,
		},
		"Mixed": {
			[]interface{}{1, 1.1, "a"},
			1,
		},
		"Nil": {
			[]interface{}{},
			nil,
		},
		"Error": {
			"wrongval",
			"unable to get first element of slice with type: string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.first(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_Last(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"String": {
			[]interface{}{"a", "b", "c"},
			"c",
		},
		"Int": {
			[]interface{}{1, 2, 3},
			3,
		},
		"Float": {
			[]interface{}{1.1, 2.2, 3.3},
			3.3,
		},
		"Mixed": {
			[]interface{}{1, 1.1, "a"},
			"a",
		},
		"Nil": {
			[]interface{}{},
			nil,
		},
		"Error": {
			"wrongval",
			"unable to get last element of slice with type: string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.last(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_Reverse(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"String": {
			[]interface{}{"a", "b", "c"},
			[]interface{}{"c", "b", "a"},
		},
		"Int": {
			[]interface{}{1, 2, 3},
			[]interface{}{3, 2, 1},
		},
		"Float": {
			[]interface{}{1.1, 2.2, 3.3},
			[]interface{}{3.3, 2.2, 1.1},
		},
		"Mixed": {
			[]interface{}{1, 1.1, "a"},
			[]interface{}{"a", 1.1, 1},
		},
		"Error": {
			input: "wrongval",
			want:  "unable to get reverse slice of type: string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.reverse(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
