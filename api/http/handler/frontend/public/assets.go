// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Assets
//
// Returns assets from the theme path and returns webp
// file if the browser accepts it.
//
// Returns a http.StatusNotFound if the asset was not found.
func (p *Public) Assets(ctx *gin.Context) {
	const op = "FrontendHandler.GetAssets"

	file, mimeType, err := p.publisher.Asset(ctx)
	if err != nil {
		p.publisher.NotFound(ctx)
		return
	}

	ctx.Data(http.StatusOK, mimeType.String(), *file)
}
