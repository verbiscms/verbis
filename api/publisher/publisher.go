// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/minify"
	"github.com/verbiscms/verbis/api/services/media"
)

// Publisher
type Publisher interface {
	Asset(g *gin.Context, webp bool) (*[]byte, domain.Mime, error)
	Upload(g *gin.Context, webp bool) (*[]byte, domain.Mime, error)
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
	media   media.Library
}

func (r *publish) SiteMap() SiteMapper {
	return r.sitemap
}

// NewRender - Construct
func NewRender(d *deps.Deps) Publisher {
	return &publish{
		d,
		minify.New(minify.Config{
			MinifyHTML: d.Options.MinifyHTML,
			MinifyCSS:  d.Options.MinifyCSS,
			MinifyJS:   d.Options.MinifyJS,
			MinifySVG:  d.Options.MinifySVG,
			MinifyJSON: d.Options.MinifyJSON,
			MinifyXML:  d.Options.MinifyXML,
		}),
		newHeaders(d.Options),
		NewSitemap(d),
		media.New(d.Options, d.Storage, d.Store.Media),
	}
}
