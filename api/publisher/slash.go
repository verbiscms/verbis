// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

func (r *publish) handleTrailingSlash(ctx *gin.Context) (string, bool) {
	p := ctx.Request.URL.Path

	if len(ctx.Request.URL.Query()) > 0 {
		return p, false
	}

	// True if options enforce slash is set in admin
	trailing := r.Options.SeoEnforceSlash
	lastChar := p[len(p)-1:]

	uri, err := url.Parse(p)
	if err != nil {
		return p, false
	}

	base := path.Base(uri.Path)
	ext := filepath.Ext(base)
	if ext != "" {
		return p, false
	}

	// Must be homepage
	if p == "/" {
		return "/", false
	}

	if lastChar != "/" && trailing {
		ctx.Redirect(301, p+"/")
		return "", true
	}

	if lastChar == "/" && !trailing {
		ctx.Redirect(301, strings.TrimSuffix(p, "/"))
		return "", true
	}

	if lastChar == "/" {
		p = strings.TrimSuffix(p, "/")
	}

	return p, false
}
