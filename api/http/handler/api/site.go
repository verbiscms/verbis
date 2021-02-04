// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
)

// SiteHandler defines methods for the Site to interact with the server
type SiteHandler interface {
	GetSite(g *gin.Context)
	GetTheme(g *gin.Context)
	GetTemplates(g *gin.Context)
	GetLayouts(g *gin.Context)
}

// Site defines the handler for Posts
type Site struct {
	*deps.Deps
}

// newSite - Construct
func NewSite(d *deps.Deps) *Site {
	return &Site{d}
}

// GetSite gets site's general config
//
// Returns 200 if site config was obtained successfully.
func (c *Site) GetSite(g *gin.Context) {
	Respond(g, 200, "Successfully obtained site config", c.Store.Site.GetGlobalConfig())
}

// GetTheme gets the theme's config from the theme path
//
// Returns 200 if theme config was obtained successfully.
// Returns 500 if there was an error getting the theme config.
func (c *Site) GetTheme(g *gin.Context) {
	Respond(g, 200, "Successfully obtained theme config", c.Store.Site.GetThemeConfig())
}

// GetTemplates gets all page templates
//
// Returns 200 if the templates were obtained successfully.
// Returns 500 if there was an error getting the templates.
func (c *Site) GetTemplates(g *gin.Context) {
	const op = "SiteHandler.GetTemplates"

	templates, err := c.Store.Site.GetTemplates()
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained templates", templates)
}

// GetTemplates gets all layouts
//
// Returns 200 if the layouts were obtained successfully.
// Returns 500 if there was an error getting the layouts.
func (c *Site) GetLayouts(g *gin.Context) {
	const op = "SiteHandler.GetLayouts"

	templates, err := c.Store.Site.GetLayouts()
	if errors.Code(err) == errors.NOTFOUND {
		Respond(g, 200, errors.Message(err), err)
		return
	} else if err != nil {
		Respond(g, 500, errors.Message(err), err)
		return
	}

	Respond(g, 200, "Successfully obtained layouts", templates)
}
