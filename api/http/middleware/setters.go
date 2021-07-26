// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/cache"
	"github.com/verbiscms/verbis/api/config"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/logger"
	"net/http"
	"path/filepath"
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
		setTheme(d, ctx)

		ctx.Next()
	}
}

func setOptions(d *deps.Deps, ctx *gin.Context) {
	var opts domain.Options
	cachedOpts, err := cache.Get(ctx, cache.OptionsKey)
	if err != nil {
		opts = d.Store.Options.Struct()
		err = cache.Set(ctx, cache.OptionsKey, opts, cache.Options{
			Expiration: time.Minute * 15,
			Tags:       nil,
		})
		if err != nil {
			logger.Panic(err)
			ctx.Next()
			return
		}
	} else {
		opts = cachedOpts.(domain.Options)
	}
	d.SetOptions(&opts)
}

var themeFetcher = config.Fetch

// setTheme
func setTheme(d *deps.Deps, ctx *gin.Context) {
	var theme domain.ThemeConfig
	cachedTheme, err := cache.Get(ctx, cache.ThemeConfigKey)
	if err != nil {
		cfg := themeFetcher(filepath.Join(d.Paths.Themes, d.Options.ActiveTheme))
		if cfg == nil {
			logger.Panic(err)
			ctx.Next()
			return
		}
		err = cache.Set(ctx, cache.ThemeConfigKey, *cfg, cache.Options{
			Expiration: time.Minute * 15,
		})
		if err != nil {
			logger.Panic(err)
			ctx.Next()
			return
		}
		theme = *cfg
	} else {
		theme = cachedTheme.(domain.ThemeConfig)
	}
	d.Config = &theme
}
