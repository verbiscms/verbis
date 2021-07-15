// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/http/handler/api/auth"
	"github.com/ainsleyclark/verbis/api/http/handler/api/cache"
	"github.com/ainsleyclark/verbis/api/http/handler/api/categories"
	"github.com/ainsleyclark/verbis/api/http/handler/api/fields"
	"github.com/ainsleyclark/verbis/api/http/handler/api/forms"
	"github.com/ainsleyclark/verbis/api/http/handler/api/media"
	"github.com/ainsleyclark/verbis/api/http/handler/api/options"
	"github.com/ainsleyclark/verbis/api/http/handler/api/posts"
	"github.com/ainsleyclark/verbis/api/http/handler/api/redirects"
	"github.com/ainsleyclark/verbis/api/http/handler/api/roles"
	"github.com/ainsleyclark/verbis/api/http/handler/api/site"
	"github.com/ainsleyclark/verbis/api/http/handler/api/storage"
	"github.com/ainsleyclark/verbis/api/http/handler/api/system"
	"github.com/ainsleyclark/verbis/api/http/handler/api/themes"
	"github.com/ainsleyclark/verbis/api/http/handler/api/users"
	"github.com/ainsleyclark/verbis/api/http/handler/frontend/public"
	"github.com/ainsleyclark/verbis/api/http/handler/frontend/seo"
	"github.com/ainsleyclark/verbis/api/http/handler/spa"
)

// APIHandler defines all handler functions for API
// routes.
type APIHandler struct {
	Auth       auth.Handler
	Cache      cache.Handler
	Categories categories.Handler
	Fields     fields.Handler
	Forms      forms.Handler
	Media      media.Handler
	Options    options.Handler
	Posts      posts.Handler
	Redirects  redirects.Handler
	Roles      roles.Handler
	Site       site.Handler
	Storage    storage.Handler
	System     system.Handler
	Themes     themes.Handler
	Users      users.Handler
}

// NewAPI returns a new API handler.
func NewAPI(d *deps.Deps) *APIHandler {
	return &APIHandler{
		Auth:       auth.New(d),
		Cache:      cache.New(d),
		Categories: categories.New(d),
		Fields:     fields.New(d),
		Forms:      forms.New(d),
		Media:      media.New(d),
		Options:    options.New(d),
		Posts:      posts.New(d),
		Redirects:  redirects.New(d),
		Roles:      roles.New(d),
		Site:       site.New(d),
		Storage:    storage.New(d),
		System:     system.New(d),
		Themes:     themes.New(d),
		Users:      users.New(d),
	}
}

// FrontendHandler defines all handler functions for
// frontend routes.
type FrontendHandler struct {
	Public public.Handler
	SEO    seo.Handler
}

// NewFrontend returns a new frontend handler.
func NewFrontend(d *deps.Deps) *FrontendHandler {
	return &FrontendHandler{
		Public: public.New(d),
		SEO:    seo.New(d),
	}
}

// SPAHandler defines all handler functions for SPA
// routes.
type SPAHandler struct {
	spa.Handler
}

// NewSPA returns a new SPA handler.
func NewSPA(d *deps.Deps) *SPAHandler {
	return &SPAHandler{
		spa.New(d),
	}
}
