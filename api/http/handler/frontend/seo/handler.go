// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/publisher"
	"github.com/gin-gonic/gin"
)

// Handler defines methods for seo routes to
// interact with the server.
type Handler interface {
	Robots(ctx *gin.Context)
	SiteMapIndex(ctx *gin.Context)
	SiteMapResource(ctx *gin.Context)
	SiteMapXSL(ctx *gin.Context, index bool)
}

// SEO defines the handler for all seo routes,
// such as sitemaps and robots.txt
type SEO struct {
	*deps.Deps
	publisher publisher.Publisher
}

// New
//
// Creates a new seo handler.
func New(d *deps.Deps) *SEO {
	return &SEO{
		Deps:      d,
		publisher: publisher.NewRender(d),
	}
}
