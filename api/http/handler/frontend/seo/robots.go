// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Robots
//
// Obtains the Seo Robots field from the options struct
// which is set in the settings, and returns the
// robots.txt file.
//
// Returns a http.StatusNotFound if the options don't allow serving of
// robots.txt
func (s *SEO) Robots(ctx *gin.Context) {
	const op = "FrontendHandler.Robots"

	if !s.Deps.Options.SeoRobotsServe {
		s.Publisher.NotFound(ctx)
		return
	}

	ctx.Data(http.StatusOK, "text/plain", []byte(s.Deps.Options.SeoRobots))
}
