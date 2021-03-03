// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	mocks "github.com/ainsleyclark/verbis/api/mocks/site"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *SiteTestSuite) TestSite_Global() {
	t.RequestAndServe(http.MethodGet, "/site", "/site", nil, func(ctx *gin.Context) {
		t.Setup(func(m *mocks.Repository) {
			m.On("Global").Return(site)
		}).Global(ctx)
	})
	t.RunT(site, http.StatusOK, "Successfully obtained site config")
}
