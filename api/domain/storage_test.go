// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageProvider_String(t *testing.T) {
	got := StorageLocal.String()
	want := string(StorageLocal)
	assert.Equal(t, got, want)
}
