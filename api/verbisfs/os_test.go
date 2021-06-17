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

func SetupOS() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, "..", "test", "testdata", "fs"), nil
}

func TestOsFS(t *testing.T) {
	path, err := SetupOS()
	assert.NoError(t, err)
	fs := &osFS{path: path}
	Open(fs, t)
	ReadFile(fs, t)
	ReadDir(fs, t)
}
