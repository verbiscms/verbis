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
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/recovery"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

type page struct {
	*deps.Deps
	ctx        *gin.Context
	post       *domain.PostDatum
	url        *url.URL
	cacheKey   string
	foundCache bool
	home       int
}

var (
	// NoPostFound is returned by page when lookup failed.
	NoPostFound = errors.New("no post found")
)

// NewPage
//
//
func newPage(d *deps.Deps, ctx *gin.Context) (page, bool, error) {
	const op = "Page.New"

	p := page{
		Deps: d,
		ctx:  ctx,
		home: d.Options.Homepage,
	}

	uri, err := url.Parse(ctx.Request.URL.Path)
	if err != nil {
		return page{}, false, fmt.Errorf("change me")
	}
	p.url = uri

	if p.HandleRedirect() {
		return page{}, true, nil
	}

	if p.HandleTrailingSlash() {
		return page{}, true, nil
	}

	post, err := p.resolve()
	if err != nil {
		return page{}, false, &errors.Error{Code: errors.NOTFOUND, Message: "No post found with the path: " + p.url.Path, Operation: op, Err: NoPostFound}
	}

	p.post = post
	p.cacheKey = cache.GetPostKey(post.Id)

	return p, false, nil
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
	template := p.Config.TemplateDir + string(os.PathSeparator) + p.post.PageTemplate
	failed, err := exec.ExecutePost(&buf, template, p.ctx, p.post)

	if err != nil {
		rec := recovery.New(p.Deps).Recover(recovery.Config{
			Code:    http.StatusInternalServerError,
			Context: p.ctx,
			Error:   err,
			TplFile: failed,
			TplExec: exec,
			Post:    p.post,
		})
		return rec, err
	}

	b := buf.Bytes()
	if p.CanCache() && !p.foundCache {
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
		Root:      p.ThemePath(),
		Extension: p.Config.FileExtension,
		Master:    p.Config.LayoutDir + string(os.PathSeparator) + p.post.PageLayout,
	})
}

// IsHomepage
//
// Determines if the page is the index.
func (p *page) IsHomepage() bool {
	return p.url.Path == "/" || p.url.Path == ""
}

// IsResourcePublic
//
// Checks if the resource is nil (a page with no resources
// attached) and loops through the themes resources.
// If there is a match, an error will be returned.
//
// Returns errors.NOTFOUND if the resource is not public.
func (p *page) IsResourcePublic() error {
	const op = "Page.IsResourcePublic"

	if p.post.HasResource() {
		for _, v := range p.Config.Resources {
			if v.Hidden && v.Name == p.post.Resource {
				return &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("The post resource is not public: %v", p.post.Resource), Operation: op, Err: fmt.Errorf("resource not public")}
			}
		}
	}

	return nil
}

// Cache
//
// Cache the post with keys.
func (p *page) Cache(b []byte) {
	cache.Store.Set(p.cacheKey, b, cache.RememberForever)
}

// GetCached
//
// Obtains the post from the store, if there was none
// found, false with nil bytes will be returned.
func (p *page) GetCached() ([]byte, bool) {
	var c interface{}
	c, ok := cache.Store.Get(p.cacheKey)
	if ok && c != nil {
		p.foundCache = true
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
	return len(p.ctx.Request.URL.Query()) > 0
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

	_, err := p.ctx.Cookie("verbis-session")
	if err != nil && !p.post.IsPublic() {
		return &errors.Error{Code: errors.NOTFOUND, Message: "Page not published, or user is not logged in", Operation: op, Err: err}
	}

	return nil
}

// resolve
//
// Returns a new post, or error by trimming leading forward
// slashes. It performs a lookup by comparing the last
// part of the URL, e.g /news/posts will be stripped
// and 'posts' will be obtained from the store.
//
// Returns errors.NOTFOUND If the permalink does not match
// the trimmed url or the slug could not be found from
// the store.
func (p *page) resolve() (*domain.PostDatum, error) {
	const op = "Page.Resolve"

	var notFoundErr = &errors.Error{
		Code:      errors.NOTFOUND,
		Message:   "No post found with the path: " + p.url.Path,
		Operation: op,
		Err:       NoPostFound,
	}

	urlTrimmed := strings.TrimSuffix(p.url.Path, "/")
	urlParts := strings.Split(urlTrimmed, "/")
	last := urlParts[len(urlParts)-1]

	homepage := p.Deps.Options.Homepage

	if last == "" {
		post, err := p.Store.Posts.Find(homepage, false)
		if err != nil {
			return nil, notFoundErr
		}
		return &post, nil
	}

	post, err := p.Store.Posts.FindBySlug(last)
	if err != nil {
		return nil, notFoundErr
	}

	if strings.TrimSuffix(post.Permalink, "/") != urlTrimmed {
		return nil, notFoundErr
	}

	return &post, nil
}

// HandleTrailingSlash
//
// Returns a bool indicating if a redirect has occurred by
// comparing the path and the enforce slash in the opts.
// If the URL contains a query parameter, or the post
// is the homepage the function will not redirect.
func (p *page) HandleTrailingSlash() bool {
	if p.HasQuery() {
		return false
	}

	trailing := p.Options.SeoEnforceSlash
	lastChar := p.url.Path[len(p.url.Path)-1:]

	// Must be homepage.
	if p.IsHomepage() {
		return false
	}

	if lastChar != "/" && trailing {
		p.ctx.Redirect(http.StatusMovedPermanently, p.url.Path+"/")
		return true
	}

	if lastChar == "/" && !trailing {
		p.ctx.Redirect(http.StatusMovedPermanently, strings.TrimSuffix(p.url.Path, "/"))
		return true
	}

	if lastChar == "/" {
		p.url.Path = strings.TrimSuffix(p.url.Path, "/")
	}

	return false
}

// HandleRedirect
//
// Returns a bool indicating if a redirect has occurred by
// stripping out unnecessary forward slashes from the
// URL.
// Logs errors.INTERNAL if there was an error compiling the
// regex used for comparing the path.
func (p *page) HandleRedirect() bool {
	const op = "Page.HandleRedirect"

	if !strings.Contains(p.url.Path, "//") {
		return false
	}

	re, err := regexp.Compile("/+") //nolint
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INTERNAL, Message: "Error compiling regex", Operation: op, Err: err})
	}

	p.ctx.Redirect(http.StatusMovedPermanently, re.ReplaceAllLiteralString(p.url.Path, "/"))

	return true
}
