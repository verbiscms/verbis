// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArgs_ToOptions(t *testing.T) {
	tt := map[string]struct {
		input Args
		want  bool
	}{
		"Marshal Error": {
			Args{"menu": make(chan bool)},
			true,
		},
		"Unmarshal Error": {
			Args{"menu": 2},
			true,
		},
		"Success": {
			Args{},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			_, err := test.input.ToOptions()
			assert.Equal(t, test.want, err != nil)
		})
	}
}
