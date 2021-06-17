// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbisfs

import (
	"github.com/ainsleyclark/verbis/api/www"
	"testing"
)

func TestEmbedFS(t *testing.T) {
	fs := &embedFS{fs: www.Web, prefix: "test"}
	Open(fs, t)
	ReadFile(fs, t)
	ReadDir(fs, t)
}
