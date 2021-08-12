// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

type Nav struct {
	Items []Item
}

type Item struct {
	Href        string
	Text        string
	IsActive    bool
	HasChildren bool
	Target      string
	Rel         []string
	Download string
	Nav      Nav
}

type Args map[string]interface{}
