// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/minify"
	"github.com/gin-gonic/gin"
)

// Publisher
type Publisher interface {
	Asset(g *gin.Context) (string, *[]byte, error)
	Upload(g *gin.Context) (domain.Mime, *[]byte, error)
	Page(g *gin.Context) ([]byte, error)
	NotFound(g *gin.Context)
	SiteMap() SiteMapper
}

// publish
type publish struct {
	*deps.Deps
	minify  minify.Minifier
	cacher  headerWriter
	sitemap *Sitemap
}

func (r *publish) SiteMap() SiteMapper {
	return r.sitemap
}

// NewRender - Construct
func NewRender(d *deps.Deps) Publisher {
	options := d.Store.Options.GetStruct()
	return &publish{
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
		NewSitemap(d),
	}
}
