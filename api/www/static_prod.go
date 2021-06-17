// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package www

import "embed"

var (
	//go:embed layouts/* mail/* public/* sitemaps/* templates/* test/*
	Web embed.FS
)
