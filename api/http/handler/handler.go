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
	"github.com/ainsleyclark/verbis/api/http/handler/api/site"
	"github.com/ainsleyclark/verbis/api/http/handler/api/users"
	"github.com/ainsleyclark/verbis/api/http/handler/frontend/public"
	"github.com/ainsleyclark/verbis/api/http/handler/frontend/seo"
	"github.com/ainsleyclark/verbis/api/http/handler/spa"
	"github.com/ainsleyclark/verbis/api/render"
)

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

// NewApi
//
// Returns a new API handler for routes.
func NewApi(d *deps.Deps) *ApiHandler {
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

type FrontendHandler struct {
	Public public.Handler
	SEO    seo.Handler
}

func NewFrontend(d *deps.Deps) *FrontendHandler {
	p := render.NewRender(d)
	return &FrontendHandler{
		Public: &public.Public{Deps: d, Publisher: p},
		SEO:    &seo.SEO{Deps: d, Publisher: p},
	}
}

type SPAHandler struct {
	spa.Handler
}

func NewSPA(d *deps.Deps) *SPAHandler {
	return &SPAHandler{
		&spa.SPA{Deps: d},
	}
}
