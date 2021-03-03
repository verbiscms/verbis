// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package media

import (
	"github.com/spf13/cast"
)

// Find
//
// Obtains the media by ID and returns a domain.Media type
// or nil if not found or the ID parameter failed to be
// parsed.
//
// Example:
// {{ $image := media 10 }}
// {{ $image.URL }}
func (ns *Namespace) Find(i interface{}) interface{} {
	id, err := cast.ToIntE(i)
	if err != nil || i == nil {
		return nil
	}

	m, err := ns.deps.Store.Media.GetByID(id)
	if err != nil {
		return nil
	}

	return m
}
