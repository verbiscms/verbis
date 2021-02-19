// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *SiteTestSuite) TestSite_Theme() {
	t.RequestAndServe(http.MethodGet, "/theme", "/theme", nil, func(ctx *gin.Context) {
		t.Setup(func(m *mocks.SiteRepository) {
			m.On("GetThemeConfig").Return(theme)
		}).Theme(ctx)
	})
	t.RunT(theme, 200, "Successfully obtained theme config")
}