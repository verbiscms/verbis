// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	app "github.com/verbiscms/verbis/api"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/http/handler/api"
	"net/http"
	"strings"
)

const (
	// InstallRoute defines the URL for installing
	// Verbis.
	InstallRoute = "/admin/install"
)

var (
	// excludedInstall are the post routes excluded from
	// being redirected.
	excludedInstall = []string{
		app.HTTPAPIRoute + "/install/validate",
		app.HTTPAPIRoute + "/install",
	}
)

// Installed redirects to the install URL if the application
// is not in install mode or the current path is not the
// install URL.
func Installed(d *deps.Deps) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if d.Installed {
			ctx.Next()
			return
		}

		url := ctx.Request.URL.String()

		if url == app.AdminInstallPath {
			ctx.Next()
			return
		}

		if strings.Contains(url, ".") {
			ctx.Next()
			return
		}

		for _, exclude := range excludedInstall {
			if ctx.Request.Method == http.MethodPost && strings.Contains(url, exclude) {
				ctx.Next()
				return
			}
		}

		if strings.Contains(url, app.HTTPAPIRoute) {
			api.AbortJSON(ctx, http.StatusBadRequest, "Verbis not installed", nil)
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, app.AdminInstallPath)
		ctx.Abort()
	}
}
