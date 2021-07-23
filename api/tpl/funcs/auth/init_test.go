// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/tpl/internal"
	"testing"
)

func TestNamespace_Init(t *testing.T) {
	d := &deps.Deps{}
	ctx := &gin.Context{}
	td := &internal.TemplateDeps{Context: ctx}

	ns := Init(d, td)
	assert.Equal(t, ns.Name, name)
	assert.Equal(t, &Namespace{deps: d, ctx: ctx}, ns.Context())
}
