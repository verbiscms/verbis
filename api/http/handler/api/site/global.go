// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Global
//
// Returns http.StatusOK if site config was obtained successfully.
func (s *Site) Global(ctx *gin.Context) {
	api.Respond(ctx, http.StatusOK, "Successfully obtained site config", s.Site.Global())
}