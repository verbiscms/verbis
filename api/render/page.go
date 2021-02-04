// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/ainsleyclark/verbis/api/tpl/tplimpl"
	"github.com/gin-gonic/gin"
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

func (r *Render) Page(g *gin.Context) ([]byte, error) {
	const op = "Render.GetPage"

	api.ServeChan <- 1
	defer func() {
		<-api.ServeChan
	}()

	//url := g.Request.URL.Path
	url, hasRedirected := r.handleTrailingSlash(g)
	if hasRedirected {
		return nil, nil
	}

	post, err := r.store.Posts.GetBySlug(url)
	if err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No page found with the url: %s", url), Operation: op, Err: err}
	}

	// Check if the file has been cached
	var foundCache bool
	cacheKey := cache.GetPostKey(post.Id)
	if r.options.CacheServerTemplates {
		var cachedTemplate interface{}
		cachedTemplate, foundCache = cache.Store.Get(cacheKey)

		if cachedTemplate != nil && foundCache {
			return cachedTemplate.([]byte), nil
		}
	}

	// Check if the resource is public
	resource := post.Resource
	if resource != nil {
		for _, v := range r.theme.Resources {
			if v.Hidden && v.Name == *resource {
				return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("The post resource is not public: %v", resource), Operation: op, Err: err}
			}
		}
	}

	_, err = g.Cookie("verbis-session")
	if err != nil && post.Status != "published" {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Page not published, or user is not logged in", Operation: op, Err: err}
	}

	pt := "index"
	if post.PageTemplate != "default" {
		pt = r.theme.TemplateDir + "/" + post.PageTemplate
	}

	master := ""
	if post.PageLayout != "default" {
		master = r.theme.LayoutDir + "/" + post.PageLayout
	} else {
		pt = pt + r.theme.FileExtension
	}

	d := &deps.Deps{
		Store:   r.store,
		Config:  r.config,
		Site:    r.store.Site.GetGlobalConfig(),
		Options: r.store.Options.GetStruct(),
		Paths: deps.Paths{
			Base:    paths.Base(),
			Admin:   paths.Admin(),
			API:     paths.Api(),
			Theme:   paths.Theme(),
			Uploads: paths.Uploads(),
			Storage: paths.Storage(),
		},
		Theme: r.store.Site.GetThemeConfig(),
	}

	t := tplimpl.New(d)
	d.Tpl = t

	e := d.Tpl.Prepare(&tpl.Config{
		Root:      paths.Theme(),
		Extension: r.theme.FileExtension,
		Master:    master,
	})

	var b bytes.Buffer
	err = e.ExecutePost(&b, pt, g, &post)
	if err != nil {
		panic(err)
	}

	//tm := tpl.New(, g, &post)

	//gvFrontend := goview.New(goview.Config{
	//	Root:         paths.Theme(),
	//	Extension:    r.theme.FileExtension,
	//	Master:       master,
	//	Partials:     []string{},
	//	//Funcs:        tm.Funcs(),
	//	DisableCache: !environment.IsProduction(),
	//})
	//
	//var b bytes.Buffer
	//err = gvFrontend.RenderWriter(&b, pt, tm.Data())

	minified, err := r.minify.MinifyBytes(&b, "text/html")
	if err != nil || minified == nil {
		return b.Bytes(), nil
	}

	go func() {
		if r.options.CacheServerTemplates && !foundCache {
			cache.Store.Set(cacheKey, minified, cache.RememberForever)
		}
	}()

	return minified, nil
}

func (r *Render) handleTrailingSlash(g *gin.Context) (string, bool) {
	p := g.Request.URL.Path

	// True if options enforce slash is set in admin
	trailing := r.options.SeoEnforceSlash
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
		g.Redirect(301, p+"/")
		return "", true
	}

	if lastChar == "/" && !trailing {
		g.Redirect(301, strings.TrimSuffix(p, "/"))
		return "", true
	}

	if lastChar == "/" {
		p = strings.TrimSuffix(p, "/")
	}

	return p, false
}
