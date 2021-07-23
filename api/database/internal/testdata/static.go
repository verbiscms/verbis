// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testdata

import "embed"

var (
	//go:embed v0/*.sql v1/*.sql v2/*.sql
	Static embed.FS
)
