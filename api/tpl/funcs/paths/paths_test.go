// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ns = New(&deps.Deps{
		Paths: deps.Paths{
			Theme: "/test/",
		},
		Theme: domain.ThemeConfig{
			TemplateDir: "templates",
			LayoutDir:   "layouts",
		},
	})
)

func TestNamespace_Templates(t *testing.T) {
	got := ns.Templates()
	assert.Equal(t, "/test/templates", got)
}

func TestNameSpace_Layouts(t *testing.T) {
	got := ns.Layouts()
	assert.Equal(t, "/test/layouts", got)
}
