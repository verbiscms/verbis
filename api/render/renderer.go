// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
)

// Renderer
type Renderer interface {
	Asset(g *gin.Context) (*string, *[]byte, error)
	Upload(g *gin.Context) (*string, *[]byte, error)
	Page(g *gin.Context) ([]byte, error)
}

// Render
type Render struct {
	store   *models.Store
	config  config.Configuration
	minify  minifier
	cacher  headerWriter
	options domain.Options
	theme   domain.ThemeConfig
}

// NewRender - Construct
func NewRender(m *models.Store, config config.Configuration) *Render {
	options := m.Options.GetStruct()
	return &Render{
		store:   m,
		config:  config,
		minify:  newMinify(options),
		cacher:  newHeaders(options),
		options: options,
		theme:   m.Site.GetThemeConfig(),
	}
}
