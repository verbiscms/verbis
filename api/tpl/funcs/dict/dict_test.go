// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dict

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ns = New(&deps.Deps{})
)

func TestNamespace_Dict(t *testing.T) {
	tt := map[string]struct {
		input []interface{}
		want  interface{}
	}{
		"Valid": {
			[]interface{}{"test", 123},
			map[string]interface{}{"test": 123},
		},
		"Odd Value": {
			[]interface{}{"test"},
			"dict values are not divisable by two",
		},
		"Not a String": {
			[]interface{}{123, 123},
			"dict keys passed are not strings",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.Dict(test.input...)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
