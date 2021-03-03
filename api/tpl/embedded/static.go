// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package embedded

import "embed"

// Static file handler for embedded meta template
// files.
var (
	//go:embed *
	Static embed.FS
)
