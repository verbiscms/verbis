// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/http/handler/api"
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

// New
//
// Creates a new site handler.
func New(d *deps.Deps) *Site {
	return &Site{
		Deps: d,
	}
}

// Global
//
// Returns http.StatusOK if site config was obtained successfully.
func (s *Site) Global(ctx *gin.Context) {
	api.Respond(ctx, http.StatusOK, "Successfully obtained site config", s.Site.Global())
}
