// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package admin

import "embed"

var (
	//go:embed dist/*
	SPA embed.FS
)
