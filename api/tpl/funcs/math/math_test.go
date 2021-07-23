// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"testing"
)

var (
	ns = New(&deps.Deps{})
)

func TestNamespace_Add(t *testing.T) {
	tt := map[string]struct {
		input []interface{}
		want  int64
	}{
		"Valid": {
			[]interface{}{1, 2, 3},
			int64(6),
		},
		"Valid 2": {
			[]interface{}{10, 10},
			int64(20),
		},
		"Strings": {
			[]interface{}{"10", "10"},
			int64(20),
		},
		"Nil": {
			[]interface{}{},
			int64(0),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Add(test.input...)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Subtract(t *testing.T) {
	tt := map[string]struct {
		a    interface{}
		b    interface{}
		want int64
	}{
		"Valid": {
			10,
			1,
			int64(9),
		},
		"Valid 2": {
			100,
			50, int64(50),
		},
		"Strings": {
			"10",
			"5", int64(5),
		},
		"Nil": {
			nil,
			nil, int64(0),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Subtract(test.a, test.b)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Divide(t *testing.T) {
	tt := map[string]struct {
		a    interface{}
		b    interface{}
		want interface{}
	}{
		"Valid": {
			16,
			2,
			int64(8),
		},
		"Valid 2": {
			100,
			50,
			int64(2),
		},
		"Strings": {
			"10",
			"5",
			int64(2),
		},
		"Nil": {
			nil,
			nil,
			"integer divide by zero",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.Divide(test.a, test.b)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Multiply(t *testing.T) {
	tt := map[string]struct {
		a    interface{}
		b    []interface{}
		want int64
	}{
		"Valid": {
			10,
			[]interface{}{10},
			int64(100)},

		"Valid 2": {
			2,
			[]interface{}{4, 4},
			int64(32),
		},
		"Strings": {
			10,
			[]interface{}{5, 2},
			int64(100),
		},
		"Nil": {
			nil,
			nil,
			int64(0),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Multiply(test.a, test.b...)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Modulus(t *testing.T) {
	tt := map[string]struct {
		a    interface{}
		b    interface{}
		want interface{}
	}{
		"Valid": {
			10,
			2,
			int64(0),
		},
		"Valid 2": {
			16,
			3,
			int64(1),
		},
		"Strings": {
			100,
			4,
			int64(0),
		},
		"Nil": {
			nil,
			nil,
			"integer divide by zero",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.Modulus(test.a, test.b)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Round(t *testing.T) {
	tt := map[string]struct {
		input interface{}
		want  float64
	}{
		"Valid": {
			10.1234,
			float64(10),
		},
		"Valid 2": {
			16,
			float64(16),
		},
		"Strings": {
			100.564988,
			float64(101),
		},
		"Nil": {
			0,
			float64(0),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Round(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Ceil(t *testing.T) {
	tt := map[string]struct {
		input interface{}
		want  float64
	}{
		"Valid": {
			10.1234,
			float64(11),
		},
		"Valid 2": {
			16,
			float64(16),
		},
		"Strings": {
			100.564988,
			float64(101),
		},
		"Nil": {
			0,
			float64(0),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Ceil(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Floor(t *testing.T) {
	tt := map[string]struct {
		input interface{}
		want  float64
	}{
		"Valid": {
			10.1234,
			float64(10),
		},
		"Valid 2": {
			16,
			float64(16),
		},
		"Strings": {
			100.564988,
			float64(100),
		},
		"Nil": {
			0,
			float64(0),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Floor(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Min(t *testing.T) {
	tt := map[string]struct {
		a    interface{}
		b    []interface{}
		want int64
	}{
		"Valid": {
			1,
			[]interface{}{2, 3, 4, 5, 6, 7, 8, 9, 10},
			int64(1),
		},
		"Valid 2": {
			102,
			[]interface{}{3004, 323, 2848},
			int64(102),
		},
		"Smaller Comparison": {
			102,
			[]interface{}{2, 40, 2949},
			int64(2),
		},
		"Strings": {
			"1",
			[]interface{}{"2", "3"},
			int64(1),
		},
		"Nil": {
			nil,
			nil,
			int64(0),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Min(test.a, test.b...)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Max(t *testing.T) {
	tt := map[string]struct {
		a    interface{}
		b    []interface{}
		want int64
	}{
		"Valid": {
			1,
			[]interface{}{2, 3, 4, 5, 6, 7, 8, 9, 10},
			int64(10),
		},
		"Valid 2": {
			102,
			[]interface{}{3004, 323, 2848},
			int64(3004),
		},
		"Smaller Comparison": {
			102,
			[]interface{}{2, 40, 2949},
			int64(2949),
		},
		"Strings": {
			"1",
			[]interface{}{"2", "3"},
			int64(3),
		},
		"Nil": {
			nil,
			nil,
			int64(0),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Max(test.a, test.b...)
			assert.Equal(t, test.want, got)
		})
	}
}
