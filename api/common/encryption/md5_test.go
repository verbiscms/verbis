// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encryption

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandom(t *testing.T) {
	got := MD5Hash("hello")
	assert.Equal(t, 32, len(got))
}
