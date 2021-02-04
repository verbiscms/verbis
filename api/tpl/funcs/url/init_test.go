// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package url

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNamespace_Init(t *testing.T) {
	d := &deps.Deps{}
	ctx := &gin.Context{}
	td := &internal.TemplateDeps{Context: ctx}

	fns := Init(d, td)
	assert.Equal(t, fns.Name, name)
	assert.Equal(t, &Namespace{deps: d, ctx: ctx}, fns.Context())
}
