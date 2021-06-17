// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package verbisfs

import (
	"fmt"
	"testing"
)

func SetupEmbed() {
	fs := New(false)

	fmt.Println(fs.SPA.ReadDir("css"))
}

func TestEmbedFS(t *testing.T) {
	SetupEmbed()

	/*fs := New(false)
	Open(fs.SPA, t)*/
	//ReadFile(fs, t)
	//ReadDir(fs, t)
}
