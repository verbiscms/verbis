// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestGet(t *testing.T) {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	assert.NoError(t, err)
	got := Get()
	want := Paths{
		Base:    path,
		Admin:   path + Admin,
		API:     path + API,
		Uploads: path + Uploads,
		Storage: path + Storage,
		Themes:  path + Themes,
		Web:     path + Web,
		Forms:   path + Forms,
		Bin:     path + Bin,
	}
	assert.Equal(t, want, got)
}

func TestGetError(t *testing.T) {
	orig := abs
	defer func() {
		abs = orig
	}()
	abs = func(path string) (string, error) {
		return "", fmt.Errorf("error")
	}
	got := base()
	assert.Equal(t, "", got)
}
