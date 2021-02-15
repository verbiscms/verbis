// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
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
