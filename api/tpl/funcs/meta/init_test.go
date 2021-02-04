// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNamespace_Init(t *testing.T) {
	d := &deps.Deps{}
	p := &domain.PostData{}
	td := &internal.TemplateDeps{Post: p}

	ns := Init(d, td)
	assert.Equal(t, ns.Name, name)
	assert.Equal(t, &Namespace{deps: d, post: p}, ns.Context())
}
