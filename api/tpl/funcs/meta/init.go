// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package meta

import (
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/tpl/internal"
	"github.com/verbiscms/verbis/api/verbis"
	"html/template"
)

// New creates a new meta Namespace.
func New(d *deps.Deps, t *internal.TemplateDeps) *Namespace {
	if t.Post.SeoMeta.Seo == nil {
		t.Post.SeoMeta.Seo = &domain.PostSeo{Private: false, ExcludeSitemap: false, Canonical: ""}
	}
	if t.Post.SeoMeta.Meta == nil {
		t.Post.SeoMeta.Meta = &domain.PostMeta{Title: "", Description: ""}
	}
	return &Namespace{
		deps:   d,
		post:   t.Post,
		crumbs: t.Breadcrumbs,
	}
}

// Namespace defines the methods for meta to be used
// as template functions.
type Namespace struct {
	deps   *deps.Deps
	post   *domain.PostDatum
	crumbs verbis.Breadcrumbs
	funcs  template.FuncMap
}

// name defines the identifier for the namespace.
const name = "meta"

// Init Creates a new Namespace and returns a new
// internal.FuncsNamespace.
func Init(d *deps.Deps, t *internal.TemplateDeps) *internal.FuncsNamespace {
	ctx := New(d, t)

	ns := &internal.FuncsNamespace{
		Name: name,
		Context: func(args ...interface{}) interface{} {
			return ctx
		},
	}

	ns.AddMethodMapping(ctx.Header,
		"verbisHead",
		[]string{"head"},
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.MetaTitle,
		"metaTitle",
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.Footer,
		"verbisFoot",
		[]string{"foot"},
		[][2]string{},
	)

	return ns
}
