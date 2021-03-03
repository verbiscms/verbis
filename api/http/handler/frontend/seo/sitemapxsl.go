// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SiteMapXSL
//
// Serves the XSL files for use with any .xml file that
// is used to serve the sitemap.
//
// Returns a http.StatusNotFound if there was an error obtaining the XSL.
func (s *SEO) SiteMapXSL(ctx *gin.Context, index bool) {
	const op = "FrontendHandler.SiteMapIndexXSL"

	sitemap, err := s.Publisher.SiteMap().XSL(index)
	if err != nil {
		s.Publisher.NotFound(ctx)
		return
	}

	ctx.Data(http.StatusOK, "application/xml; charset=utf-8", sitemap)
}
