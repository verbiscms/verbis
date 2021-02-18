// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/http/handler/api/auth"
	"github.com/ainsleyclark/verbis/api/http/handler/api/cache"
	"github.com/ainsleyclark/verbis/api/http/handler/api/categories"
	"github.com/ainsleyclark/verbis/api/http/handler/api/fields"
	"github.com/ainsleyclark/verbis/api/http/handler/api/forms"
	"github.com/ainsleyclark/verbis/api/http/handler/api/media"
	"github.com/ainsleyclark/verbis/api/http/handler/api/options"
	"github.com/ainsleyclark/verbis/api/http/handler/api/posts"
	"github.com/ainsleyclark/verbis/api/http/handler/api/redirects"
	"github.com/ainsleyclark/verbis/api/http/handler/api/site"
	"github.com/ainsleyclark/verbis/api/http/handler/api/users"
	"github.com/ainsleyclark/verbis/api/http/handler/frontend"
	"github.com/ainsleyclark/verbis/api/http/handler/spa"
)

// Handler defines all of handler funcs for the app.
type Handler struct {
	Auth       api.AuthHandler
	Cache      api.CacheHandler
	Categories api.CategoryHandler
	Media      api.MediaHandler
	Options    api.OptionsHandler
	Posts      api.PostHandler
	Redirects  api.RedirectHandler
	Site       api.SiteHandler
	User       api.UserHandler
	Forms      api.FormHandler
	Fields     api.FieldHandler
	Frontend   frontend.PublicHandler
	SEO        frontend.SEOHandler
	SPA        spa.SPAHandler
}

// Construct
func New(d *deps.Deps) *Handler {
	return &Handler{
		Auth:       api.NewAuth(d),
		Cache:      api.NewCache(d),
		Categories: api.NewCategories(d),
		Fields:     api.NewFields(d),
		Forms:      api.NewForms(d),
		Media:      api.NewMedia(d),
		Options:    api.NewOptions(d),
		Posts:      api.NewPosts(d),
		Redirects:  api.NewRedirects(d),
		Site:       api.NewSite(d),
		User:       api.NewUser(d),
		SPA:        spa.NewSpa(d),
		Frontend:   frontend.NewPublic(d),
		SEO:        frontend.NewSEO(d),
	}
}


// Handler defines all of handler funcs for the API.
type ApiHandler struct {
	Auth       auth.Handler
	Cache      cache.Handler
	Categories categories.Handler
	Fields     fields.Handler
	Forms      forms.Handler
	Media      media.Handler
	Options    options.Handler
	Posts      posts.Handler
	Redirects  redirects.Handler
	Site       site.Handler
	Users      users.Handler
}

// NewApiHandler
//
// Returns a new API handler for routes.
func NewApiHandler(d *deps.Deps) *ApiHandler {
	return &ApiHandler{
		Auth:       &auth.Auth{Deps: d},
		Cache:      &cache.Cache{Deps: d},
		Categories: &categories.Categories{},
		Fields:     &fields.Fields{Deps: d},
		Forms:      &forms.Forms{Deps: d},
		Media:      &media.Media{Deps: d},
		Options:    &options.Options{Deps: d},
		Posts:      &posts.Posts{Deps: d},
		Redirects:  &redirects.Redirects{Deps: d},
		Site:       &site.Site{Deps: d},
		Users:      &users.Users{Deps: d},
	}
}