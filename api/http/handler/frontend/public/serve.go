// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Serve
//
// Serves the front end website by obtaining the post
// from the store. If the error code is anything
// but errors.NOTFOUND, a http.StatusInternalServerError response will
// be sent.
//
// Returns a http.StatusNotFound if the post was not found.
// Returns a http.StatusInternalServerError if the template file failed to execute.
func (p *Public) Serve(ctx *gin.Context) {
	const op = "FrontendHandler.Serve"

	page, err := p.Publisher.Page(ctx)
	if errors.Code(err) == errors.NOTFOUND {
		p.Publisher.NotFound(ctx)
		return
	} else if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html", page)
		return
	}

	ctx.Data(http.StatusOK, "text/html", page)
}
