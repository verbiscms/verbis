// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// Theme
//
// Returns 200 if theme config was obtained successfully.
// Returns 500 if there was an error getting the theme config.
func (s *Site) Theme(ctx *gin.Context) {
	api.Respond(ctx, 200, "Successfully obtained theme config", s.Store.Site.GetThemeConfig())
}