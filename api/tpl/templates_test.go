// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tpl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_GetRoot(t *testing.T) {
	c := Config{Root: "test"}
	got := c.GetRoot()
	assert.Equal(t, "test", got)
}

func TestConfig_GetExtension(t *testing.T) {
	c := Config{Extension: "test"}
	got := c.GetExtension()
	assert.Equal(t, "test", got)
}

func TestConfig_GetMaster(t *testing.T) {
	c := Config{Master: "test"}
	got := c.GetMaster()
	assert.Equal(t, "test", got)
}
