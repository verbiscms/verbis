// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package attributes

import (
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/tpl/funcs/auth"
	"github.com/verbiscms/verbis/api/tpl/internal"
	"testing"
)

func TestNamespace_Init(t *testing.T) {
	d := &deps.Deps{}
	p := &domain.PostDatum{}
	td := &internal.TemplateDeps{Post: p}

	ns := Init(d, td)
	assert.Equal(t, ns.Name, name)
	assert.Equal(t, &Namespace{deps: d, tpld: td, auth: auth.New(d, td)}, ns.Context())
}
