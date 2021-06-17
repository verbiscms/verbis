// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbisfs

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestOsFS(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)
	fs := &osFS{path: filepath.Join(wd, "..", "www", "test")}
	Open(fs, t)
	ReadFile(fs, t)
	ReadDir(fs, t)
}
