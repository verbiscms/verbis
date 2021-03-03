// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ns = New(&deps.Deps{})
)

func TestNamespace_RandInt(t *testing.T) {
	tt := map[string]struct {
		a    interface{}
		b    interface{}
		want []int
	}{
		"Valid": {
			1,
			100,
			[]int{1, 100},
		},
		"Valid 2": {
			1,
			5,
			[]int{1, 5},
		},
		"Large": {
			1,
			99999999,
			[]int{1, 99999999},
		},
		"Nil A": {
			nil,
			1,
			[]int{0, 0},
		},
		"Nil B": {
			1,
			nil,
			[]int{0, 0},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Int(test.a, test.b)
			if test.want[0] > got || test.want[1] < got {
				t.Error(fmt.Errorf("got %v expecting in between %v", got, test.want))
			}
		})
	}
}

func TestNamespace_RandFloat(t *testing.T) {
	tt := map[string]struct {
		a    interface{}
		b    interface{}
		want []float64
	}{
		"Valid": {
			1,
			100,
			[]float64{1, 100},
		},
		"Valid2": {
			1.555555,
			10.44444,
			[]float64{1.555555, 10.44444},
		},
		"Large": {
			1,
			99999999.99999,
			[]float64{1, 99999999.99999},
		},
		"Nil A": {
			nil,
			1,
			[]float64{0.0, 0.0},
		},
		"Nil B": {
			1,
			nil,
			[]float64{0.0, 0.0},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Float(test.a, test.b)
			if test.want[0] > got || test.want[1] < got {
				t.Error(fmt.Errorf("got %v expecting in between %v", got, test.want))
			}
		})
	}
}

func TestNamespace_RandAlpha(t *testing.T) {
	tt := map[string]struct {
		len  int64
		want int
	}{
		"Valid": {
			5,
			5,
		},
		"Valid 2": {
			10,
			10,
		},
		"Valid 3": {
			100,
			100,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Alpha(test.len)
			assert.Equal(t, test.len, int64(len(got)))
		})
	}
}

func TestNamespace_RandAlphaNum(t *testing.T) {

	tt := map[string]struct {
		len  int64
		want int
	}{
		"Valid": {
			5,
			5,
		},
		"Valid 2": {
			10,
			10,
		},
		"Valid 3": {
			100,
			100,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.AlphaNum(test.len)
			assert.Equal(t, test.len, int64(len(got)))
		})
	}
}
