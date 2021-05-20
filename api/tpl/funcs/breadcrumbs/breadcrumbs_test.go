// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package breadcrumbs

import (
	"github.com/ainsleyclark/verbis/api/verbis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNamespace_Get(t *testing.T) {
	c := verbis.Breadcrumbs{
		Enabled: true,
		Title:   "Items",
	}
	ns := Namespace{
		deps:   nil,
		crumbs: c,
	}
	got := ns.Get()
	assert.Equal(t, c, got)
}
