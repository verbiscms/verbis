// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package frontend

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/render"
	"github.com/gin-gonic/gin"
)

// PublicHandler defines methods for the frontend to interact with the server
type PublicHandler interface {
	GetUploads(g *gin.Context)
	GetAssets(g *gin.Context)
	Serve(g *gin.Context)
}

// Public defines the handler for all frontend routes
type Public struct {
	*deps.Deps
	render render.Renderer
	render.ErrorHandler
}

// NewPublic - Construct
func NewPublic(d *deps.Deps) *Public {
	return &Public{
		Deps:         d,
		render:       render.NewRender(d),
		ErrorHandler: &render.Errors{Deps: d},
	}
}

// GetUploads retrieves images & media in the uploads folder, returns webp if accepts.
func (c *Public) GetUploads(g *gin.Context) {
	const op = "FrontendHandler.GetUploads"

	mimeType, file, err := c.render.Upload(g)
	if err != nil {
		c.NotFound(g)
		return
	}

	g.Data(200, *mimeType, *file)
}

// GetAssets retrieves assets from the theme path, returns webp if accepts.
func (c *Public) GetAssets(g *gin.Context) {
	const op = "FrontendHandler.GetAssets"

	mimeType, file, err := c.render.Asset(g)
	if err != nil {
		c.NotFound(g)
		return
	}

	g.Data(200, *mimeType, *file)
}

// Serve the front end website
func (c *Public) Serve(g *gin.Context) {
	const op = "FrontendHandler.Serve"

	page, err := c.render.Page(g)
	if errors.Code(err) == errors.NOTFOUND {
		c.NotFound(g)
		return
	} else {
		g.Data(500, "text/html", page)
		return
	}

	g.Data(200, "text/html", page)
}
