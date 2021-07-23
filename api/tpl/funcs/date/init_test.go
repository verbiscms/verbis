// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/tpl/internal"
	"testing"
)

func TestNamespace_Init(t *testing.T) {
	var found bool
	var ns *internal.FuncsNamespace

	for _, nsf := range internal.GenericNamespaceRegistry {
		ns = nsf(&deps.Deps{})
		if ns.Name == name {
			found = true
			break
		}
	}

	assert.True(t, found)
	assert.Equal(t, &Namespace{&deps.Deps{}}, ns.Context())
}
