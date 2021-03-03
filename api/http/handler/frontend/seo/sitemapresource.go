// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SiteMapResource
//
// Obtains the sitemap pages from the sitemap model by
// using the resource in the URL. Obtains the
// []bytes to send back as data when
// /:resource/sitemap.xml is visited.
//
// Returns a 404 if there was an error obtaining the XML
// file or there was no resource items found.
func (s *SEO) SiteMapResource(ctx *gin.Context) {
	const op = "FrontendHandler.SiteMap"

	sitemap, err := s.Publisher.SiteMap().Pages(ctx.Param("resource"))
	if err != nil {
		s.Publisher.NotFound(ctx)
		return
	}

	ctx.Data(http.StatusOK, "application/xml; charset=utf-8", sitemap)
}
