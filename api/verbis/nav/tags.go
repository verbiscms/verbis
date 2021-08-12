// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import "strings"

// Relative defines the tag for 'rel' appearing on
// link items.
type Relative []string

// Tags joins the tag on a space, for use with
// templates.
func (r Relative) Tags() string {
	var parts []string
	for _, s := range r {
		if strings.TrimSpace(s) != "" {
			parts = append(parts, s)
		}
	}
	return strings.Join(parts, " ")
}

// HasTags determines if the rel tags in the slice
// to output.
func (r Relative) HasTags() bool {
	return len(r) != 0
}
