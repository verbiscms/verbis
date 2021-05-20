// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package breadcrumbs

import (
	"github.com/ainsleyclark/verbis/api/verbis"
)

// Get
//
// Obtains the breadcrumbs for the post.
//
// Example: {{ $crumbs := breadcrumbs }}
func (ns *Namespace) Get() verbis.Breadcrumbs {
	return ns.crumbs
}
