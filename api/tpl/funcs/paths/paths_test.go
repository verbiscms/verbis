// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ns = New(&deps.Deps{
		Paths: paths.Paths{
			Base:    "/base",
			Admin:   "/admin/",
			API:     "/api/",
			Uploads: "/uploads/",
			Storage: "/storage/",
			Web:     "/web/",
		},
		Config: &domain.ThemeConfig{
			TemplateDir: "templates",
			LayoutDir:   "layouts",
			AssetsPath:  "/assets/",
		},
	})
)

func TestNamespace_Base(t *testing.T) {
	got := ns.Base()
	assert.Equal(t, "/base", got)
}

func TestNamespace_Admin(t *testing.T) {
	got := ns.Admin()
	assert.Equal(t, "/admin/", got)
}

func TestNamespace_API(t *testing.T) {
	got := ns.API()
	assert.Equal(t, "/api/", got)
}

func TestNamespace_Theme(t *testing.T) {
	got := ns.Theme()
	assert.Equal(t, "/base/theme", got)
}

func TestNamespace_Uploads(t *testing.T) {
	got := ns.Uploads()
	assert.Equal(t, "/uploads/", got)
}

func TestNamespace_Storage(t *testing.T) {
	got := ns.Storage()
	assert.Equal(t, "/storage/", got)
}

func TestNamespace_Assets(t *testing.T) {
	got := ns.Assets()
	assert.Equal(t, "/assets/", got)
}

func TestNamespace_Templates(t *testing.T) {
	got := ns.Templates()
	assert.Equal(t, "/base/theme/templates", got)
}

func TestNameSpace_Layouts(t *testing.T) {
	got := ns.Layouts()
	assert.Equal(t, "/base/theme/layouts", got)
}
