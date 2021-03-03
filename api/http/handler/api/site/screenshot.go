// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Themes
//
// Returns http.StatusInternalServerError if there was an error getting the layouts.
// Returns http.StatusOK if the themes were obtained successfully or there were none found.
func (s *Site) Screenshot(ctx *gin.Context) {
	const op = "SiteHandler.Layouts"

	theme := ctx.Param("theme")
	if theme == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	file := ctx.Param("file")
	if file == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	screenshot, mime, err := s.Site.Screenshot(s.Paths.Base, theme, file)
	if errors.Code(err) == errors.NOTFOUND {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Data(http.StatusOK, mime, screenshot)
}
