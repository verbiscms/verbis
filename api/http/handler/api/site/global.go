// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// Global
//
// Returns 200 if site config was obtained successfully.
func (s *Site) Global(ctx *gin.Context) {
	api.Respond(ctx, 200, "Successfully obtained site config", s.Store.Site.GetGlobalConfig())
}
