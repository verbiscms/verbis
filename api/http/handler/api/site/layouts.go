// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Layouts
//
// Returns http.StatusInternalServerError if there was an error getting the layouts.
// Returns http.StatusOK if the layouts were obtained successfully or there were none found.
func (s *Site) Layouts(ctx *gin.Context) {
	const op = "SiteHandler.Layouts"

	templates, err := s.Site.Layouts()
	if errors.Code(err) == errors.NOTFOUND {
		api.Respond(ctx, http.StatusOK, errors.Message(err), err)
		return
	} else if err != nil {
		api.Respond(ctx, http.StatusInternalServerError, errors.Message(err), err)
		return
	}

	api.Respond(ctx, http.StatusOK, "Successfully obtained layouts", templates)
}
