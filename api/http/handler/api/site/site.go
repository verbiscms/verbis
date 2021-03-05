// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler defines methods for the site to interact with the server.
type Handler interface {
	Global(ctx *gin.Context)
}

// Site defines the handler for all site routes.
type Site struct {
	*deps.Deps
}

// Global
//
// Returns http.StatusOK if site config was obtained successfully.
func (s *Site) Global(ctx *gin.Context) {
	api.Respond(ctx, http.StatusOK, "Successfully obtained site config", s.Site.Global())
}
