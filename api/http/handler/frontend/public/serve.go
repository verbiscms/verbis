// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package public

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
)

// Serve
//
// Serves the front end website by obtaining the post
// from the store. If the error code is anything
// but errors.NOTFOUND, a 500 response will
// be sent.
//
// Returns a 404 if the post was not found.
// Returns a 500 if the template file failed to execute.
func (p *Public) Serve(ctx *gin.Context) {
	const op = "FrontendHandler.Serve"

	page, err := p.Publisher.Page(ctx)
	if errors.Code(err) == errors.NOTFOUND {
		p.Publisher.NotFound(ctx)
		return
	} else if err != nil {
		ctx.Data(500, "text/html", page)
		return
	}

	ctx.Data(200, "text/html", page)
}
