// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/minify"
	"github.com/gin-gonic/gin"
)

// Renderer
type Renderer interface {
	Asset(g *gin.Context) (string, *[]byte, error)
	Upload(g *gin.Context) (string, *[]byte, error)
	Page(g *gin.Context) ([]byte, error)
	NotFound(g *gin.Context)
}

// Render
type Render struct {
	*deps.Deps
	minify minify.Minifier
	cacher headerWriter
}

// NewRender - Construct
func NewRender(d *deps.Deps) *Render {
	options := d.Store.Options.GetStruct()
	return &Render{
		d,
		minify.New(minify.Config{
			MinifyHTML: options.MinifyHTML,
			MinifyCSS:  options.MinifyCSS,
			MinifyJS:   options.MinifyJS,
			MinifySVG:  options.MinifySVG,
			MinifyJSON: options.MinifyJSON,
			MinifyXML:  options.MinifyXML,
		}),
		newHeaders(options),
	}
}
