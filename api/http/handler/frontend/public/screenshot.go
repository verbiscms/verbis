// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Screenshot
//
// Retrieves the screenshot of the theme passed.
//
// Returns http.StatusInternalServerError if there was an error getting the layouts.
// Returns http.StatusOK if the themes were obtained successfully or there were none found.
func (p *Public) Screenshot(ctx *gin.Context) {
	const op = "FrontendHandler.Screenshot"

	theme := ctx.Param("theme")
	if theme == "" {
		p.Publisher.NotFound(ctx)
		return
	}

	file := ctx.Param("file")
	if file == "" {
		p.Publisher.NotFound(ctx)
		return
	}

	screenshot, mime, err := p.Theme.Screenshot(theme, file)
	if errors.Code(err) == errors.NOTFOUND {
		p.Publisher.NotFound(ctx)
		return
	}

	ctx.Data(http.StatusOK, mime, screenshot)
}
