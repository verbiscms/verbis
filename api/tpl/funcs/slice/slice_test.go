// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slice

import (
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"testing"
)

var (
	ns = New(&deps.Deps{})
)

func TestNamespace_Slice(t *testing.T) {
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
			got := ns.Slice(test.input...)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Append(t *testing.T) {
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
			got, err := ns.Append(test.slice, test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Prepend(t *testing.T) {
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
			got, err := ns.Prepend(test.slice, test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_First(t *testing.T) {
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
			got, err := ns.First(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Last(t *testing.T) {
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
			got, err := ns.Last(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Reverse(t *testing.T) {
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
			got, err := ns.Reverse(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
