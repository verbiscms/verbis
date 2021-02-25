// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package publisher

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/recovery"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type page struct {
	*deps.Deps
	Context    *gin.Context
	Post       *domain.PostDatum
	Url        string
	CacheKey   string
	FoundCache bool
	Type       TypeOfPage
}

// Execute
//
// Executes the page template. If there was a user error
// within the tpl or the template was not found,
// a new recovery will be created displaying
// an error page. The page will be cached
// if the options allow for it.
func (p *page) Execute() ([]byte, error) {
	var buf bytes.Buffer

	exec := p.Prepare()
	template := p.Theme.TemplateDir + string(os.PathSeparator) + p.Post.PageTemplate
	failed, err := exec.ExecutePost(&buf, template, p.Context, p.Post)

	if err != nil {
		rec := recovery.New(p.Deps).Recover(recovery.Config{
			Code:    http.StatusInternalServerError,
			Context: p.Context,
			Error:   err,
			TplFile: failed,
			TplExec: exec,
			Post:    p.Post,
		})
		return rec, err
	}

	b := buf.Bytes()
	if p.CanCache() && !p.FoundCache {
		go p.Cache(b)
	}

	return b, nil
}

// Prepare
//
// Prepares the page template for execution using the
// post data and paths.
func (p *page) Prepare() tpl.TemplateExecutor {
	return p.Tmpl().Prepare(&tpl.Config{
		Root:      p.Paths.Theme,
		Extension: p.Theme.FileExtension,
		Master:    p.Theme.LayoutDir + string(os.PathSeparator) + p.Post.PageLayout,
	})
}

// IsHomepage
//
// Determines if the page is the index.
func (p *page) IsHomepage() bool {
	return p.Url == "/" || p.Url == ""
}

// IsResourcePublic
//
// Checks if the resource is nil (a page with no resources
// attached) and loops through the themes resources.
// If there is a match, an error will be returned.
//
// Returns errors.NOTFOUND if the resource is not public.
func (p *page) IsResourcePublic() error {
	const op = "Publisher.Page.IsResourcePublic"

	resource := p.Post.Resource
	if resource != nil {
		for _, v := range p.Theme.Resources {
			if v.Hidden && v.Name == *resource {
				return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("The post resource is not public: %v", resource), Operation: op, Err: fmt.Errorf("resource not public")}
			}
		}
	}

	return nil
}

// Cache
//
// Cache the post with keys.
func (p *page) Cache(b []byte) {
	cache.Store.Set(p.CacheKey, b, cache.RememberForever)
}

// GetCached
//
// Obtains the post from the store, if there was none
// found, false with nil bytes will be returned.
func (p *page) GetCached() ([]byte, bool) {
	var c interface{}
	c, ok := cache.Store.Get(p.CacheKey)
	if ok && c != nil {
		p.FoundCache = true
		return c.([]byte), true
	}
	return nil, false
}

// CanCache
//
// Determines if the post can be cached by using the
// CacheServerTemplates bool set in the options.
// If there is a query attached to the post,
// the page cannot be cached.
func (p *page) CanCache() bool {
	return !p.HasQuery() && p.Options.CacheServerTemplates
}

// HasQuery
//
// Determines if the url has a query parameter.
func (p *page) HasQuery() bool {
	return len(p.Context.Request.URL.Query()) > 0
}

// CheckSession
//
// Checks if the user has a verbis-session cookie, if the
// user is not logged in, or the post is not public,
// and error will be returned.
//
// Returns errors.NOTFOUND if the page is not public.
func (p *page) CheckSession() error {
	const op = "Publisher.Page.CheckSession"

	_, err := p.Context.Cookie("verbis-session")
	if err != nil && p.Post.IsPublic() {
		return &errors.Error{Code: errors.NOTFOUND, Message: "Page not published, or user is not logged in", Operation: op, Err: err}
	}

	return nil
}

// HandleTrailingSlash
//
//
func (p *page) HandleTrailingSlash() (string, bool) {
	pth := p.Context.Request.URL.Path

	if p.HasQuery() {
		return pth, false
	}

	// True if options enforce slash is set in admin
	trailing := p.Options.SeoEnforceSlash
	lastChar := pth[len(pth)-1:]

	uri, err := url.Parse(pth)
	if err != nil {
		return pth, false
	}

	base := path.Base(uri.Path)
	ext := filepath.Ext(base)
	if ext != "" {
		return pth, false
	}

	// Must be homepage
	if p.IsHomepage() {
		return "/", false
	}

	if lastChar != "/" && trailing {
		p.Context.Redirect(301, pth+"/")
		return "", true
	}

	if lastChar == "/" && !trailing {
		p.Context.Redirect(301, strings.TrimSuffix(pth, "/"))
		return "", true
	}

	if lastChar == "/" {
		pth = strings.TrimSuffix(pth, "/")
	}

	return pth, false
}
