// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/errors"
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

	screenshot, mime, err := p.Theme.Screenshot(ctx.Param("theme"), ctx.Param("file"))
	if errors.Code(err) == errors.NOTFOUND {
		p.publisher.NotFound(ctx)
		return
	}

	ctx.Data(http.StatusOK, mime.String(), screenshot)
}
