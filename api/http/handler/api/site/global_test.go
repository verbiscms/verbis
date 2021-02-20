// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *SiteTestSuite) TestSite_Global() {
	t.RequestAndServe(http.MethodGet, "/site", "/site", nil, func(ctx *gin.Context) {
		t.Setup(func(m *mocks.SiteRepository) {
			m.On("GetGlobalConfig").Return(site)
		}).Global(ctx)
	})
	t.RunT(site, 200, "Successfully obtained site config")
}
