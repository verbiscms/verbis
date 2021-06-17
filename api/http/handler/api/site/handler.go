// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/version"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler defines methods for the site to interact with the server.
type Handler interface {
	Global(ctx *gin.Context)
	Update(ctx *gin.Context)
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

// Update
//
// TODO, update comments
// Returns http.StatusOK if the user was created.
// Returns http.StatusInternalServerError if there was an error creating the user.
// Returns http.StatusBadRequest if the the validation failed or a user already exists.
func (s *Site) Update(ctx *gin.Context) {
	const op = "SiteHandler.Update"

	canUpdate, err := s.Updater.HasUpdate()
	if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, "Error obtaining latest version", &errors.Error{Code: errors.INVALID, Err: err, Operation: op})
		return
	}

	if !canUpdate {
		api.Respond(ctx, http.StatusBadRequest, "Verbis is up to date", &errors.Error{Code: errors.INVALID, Err: fmt.Errorf("verbis up to date running version %s", version.Version), Operation: op})
	}

	code, err := s.Updater.Update()
	if err != nil {
		api.Respond(ctx, http.StatusOK, "Verbis updated successfully to version "+s.Updater.LatestVersion(), nil)
	}

}
