// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SiteMapIndex obtains the sitemap index file from the sitemap
// model Obtains the []bytes to send back as data when
// /sitemap.xml is visited.
//
// Returns a http.StatusNotFound if there was an error obtaining the XML file.
// or there was no resource items found.
func (s *SEO) SiteMapIndex(ctx *gin.Context) {
	const op = "FrontendHandler.SiteMapIndex"

	sitemap, err := s.publisher.SiteMap().Index()
	if err != nil {
		s.publisher.NotFound(ctx)
		return
	}

	ctx.Data(http.StatusOK, "application/xml; charset=utf-8", sitemap)
}
