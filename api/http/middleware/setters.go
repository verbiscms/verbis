// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	"net/http"
	"strings"
	"time"
)

// Setters ensures that the theme options and domain configuration
// is always up to date on deps.Deps. If the request is a file
// or the method is not http.MethodGet, the middleware will
// break and continue to the next request.
// Panics with error if the options could not be to to the cache
// or the theme configuration could not be found.
func Setters(d *deps.Deps) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method != http.MethodGet {
			ctx.Next()
			return
		}
		if strings.Contains(ctx.Request.URL.Path, ".") {
			ctx.Next()
			return
		}
		setOptions(d, ctx)
		setTheme(d)
		ctx.Next()
	}
}

// setOptions retrieves the options from the
// Options Service and sets it to deps.
func setOptions(d *deps.Deps, ctx *gin.Context) {
	var opts domain.Options
	cachedOpts, err := d.Cache.Get(ctx, cache.OptionsKey)
	if err == nil {
		opts = cachedOpts.(domain.Options)
	} else {
		opts = d.Store.Options.Struct()
		d.Cache.Set(ctx, cache.OptionsKey, opts, cache.Options{
			Expiration: time.Minute * 15,
		})
	}
	d.Options = &opts
}

// setTheme retrieves the theme configuration from the
// Theme Service and sets it to deps.
func setTheme(d *deps.Deps) {
	config, err := d.Theme.Config()
	if err != nil {
		logger.Panic(err)
		return
	}
	d.Config = &config
}
