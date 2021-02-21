// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import "github.com/gin-gonic/gin"

// Assets
//
// Returns assets from the theme path and returns webp
// file if the browser accepts it.
//
// Returns a 404 if the asset was not found.
func (p *Public) Assets(ctx *gin.Context) {
	const op = "FrontendHandler.GetAssets"

	mimeType, file, err := p.Publisher.Asset(ctx)
	if err != nil {
		p.Publisher.NotFound(ctx)
		return
	}

	ctx.Data(200, mimeType, *file)
}
