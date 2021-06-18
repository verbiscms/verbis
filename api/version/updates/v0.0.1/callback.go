// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v0_0_1 //nolint

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("im in init")
	fmt.Println(os.Executable())
}
