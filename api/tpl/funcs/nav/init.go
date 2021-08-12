// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/tpl/internal"
	"github.com/verbiscms/verbis/api/verbis/nav"
)

// Creates a new breadcrumbs Namespace
func New(d *deps.Deps, t *internal.TemplateDeps) *Namespace {
	nav, err := nav.New(d, t.Post)
	if err != nil {
		logger.WithError(err).Panic()
	}
	return &Namespace{
		deps: d,
		nav:  nav,
	}
}

// Namespace defines the methods breadcrumbs posts to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	nav  nav.Getter
}

const name = "nav"

// Init creates a new Namespace and returns a new internal.FuncsNamespace
func Init(d *deps.Deps, t *internal.TemplateDeps) *internal.FuncsNamespace {
	ctx := New(d, t)

	ns := &internal.FuncsNamespace{
		Name:    name,
		Context: func(args ...interface{}) interface{} { return ctx },
	}

	ns.AddMethodMapping(ctx.Get,
		"nav",
		nil,
		nil,
	)

	ns.AddMethodMapping(ctx.HTML,
		"navHTML",
		nil,
		nil,
	)

	return ns
}
