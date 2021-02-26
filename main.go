// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/ainsleyclark/verbis/api/cmd"
	"runtime"
)

func main() {
	// Execute Verbis
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}