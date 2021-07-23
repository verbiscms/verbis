// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package themes

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/config"
	"net/http"
)

func (t *ThemesTestSuite) TestThemes_Config() {
	t.RequestAndServe(http.MethodGet, "/theme", "/theme", nil, func(ctx *gin.Context) {
		t.Setup(nil).Config(ctx)
	})
	t.RunT(config.DefaultTheme, http.StatusOK, "Successfully obtained theme config")
}
