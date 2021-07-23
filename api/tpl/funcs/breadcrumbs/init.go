// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package breadcrumbs

import (
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/tpl/internal"
	"github.com/verbiscms/verbis/api/verbis"
)

// Creates a new breadcrumbs Namespace
func New(d *deps.Deps, t *internal.TemplateDeps) *Namespace {
	return &Namespace{
		deps:   d,
		crumbs: t.Breadcrumbs,
	}
}

// Namespace defines the methods breadcrumbs posts to be used
// as template functions.
type Namespace struct {
	deps   *deps.Deps
	crumbs verbis.Breadcrumbs
}

const name = "breadcrumbs"

//  Creates a new Namespace and returns a new internal.FuncsNamespace
func Init(d *deps.Deps, t *internal.TemplateDeps) *internal.FuncsNamespace {
	ctx := New(d, t)

	ns := &internal.FuncsNamespace{
		Name:    name,
		Context: func(args ...interface{}) interface{} { return ctx },
	}

	ns.AddMethodMapping(ctx.Get,
		"breadcrumbs",
		[]string{"crumbs"},
		nil,
	)

	ns.AddMethodMapping(ctx.HTML,
		"breadcrumbsHTML",
		[]string{"crumbsHTML"},
		nil,
	)
	return ns
}
