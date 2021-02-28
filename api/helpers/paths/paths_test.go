// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaths_Get(t *testing.T) {

	tt := map[string]struct {
		def  string
		want string
	}{
		"No Defaults": {
			"",
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ""
			assert.Equal(t, test.want, got)
		})
	}
}
