// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import "github.com/gin-gonic/gin"

// SiteMapIndex obtains the sitemap index file from the sitemap
// model Obtains the []bytes to send back as data when
// /sitemap.xml is visited.
//
// Returns a 404 if there was an error obtaining the XML file.
// or there was no resource items found.
func (c *SEO) SiteMapIndex(ctx *gin.Context) {
	const op = "FrontendHandler.SiteMapIndex"

	sitemap, err := c.Sitemap.GetIndex()
	if err != nil {
		c.Publisher.NotFound(ctx)
		return
	}

	ctx.Data(200, "application/xml; charset=utf-8", sitemap)
}