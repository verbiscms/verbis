// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Uploads
//
// Returns assets from the uploads dir and returns webp
// file if the browser accepts it.
//
// Returns a http.StatusNotFound if the upload was not found.
func (p *Public) Uploads(ctx *gin.Context) {
	const op = "FrontendHandler.GetUploads"

	webp := true
	if ctx.Query("webp") == "false" {
		webp = false
	}

	mime, file, err := p.publisher.Upload(ctx, webp)
	if err != nil {
		p.publisher.NotFound(ctx)
		return
	}

	ctx.Data(http.StatusOK, mime.String(), *file)
}
