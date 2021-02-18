// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
)

// Layouts
//
// Returns 200 if the layouts were obtained successfully.
// Returns 500 if there was an error getting the layouts.
func (s *Site) Layouts(ctx *gin.Context) {
	const op = "SiteHandler.Layouts"

	templates, err := s.Store.Site.GetLayouts()
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, 200, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, 500, errors.Message(err), err)
		return
	}

	api.Respond(ctx, 200, "Successfully obtained layouts", templates)
}