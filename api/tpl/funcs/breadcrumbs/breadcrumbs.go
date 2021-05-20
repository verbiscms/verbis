// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/ainsleyclark/verbis/api/verbis"
)

// Find
//
// Obtains the post by ID and returns a domain.PostDatum type
// or nil if not found.
//
// Example: {{ post 123 }}
func (ns *Namespace) Get() verbis.Breadcrumbs {
	return ns.deps
}
