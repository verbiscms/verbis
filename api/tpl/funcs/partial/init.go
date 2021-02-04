// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package partial

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new partial Namespace
func New(d *deps.Deps, t *internal.TemplateDeps) *Namespace {
	return &Namespace{
		deps: d,
		tpld: t,
	}
}

// Namespace defines the methods for partials to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	tpld *internal.TemplateDeps
}

const name = "partial"

//  Creates a new Namespace and returns a new internal.FuncsNamespace
func Init(d *deps.Deps, t *internal.TemplateDeps) *internal.FuncsNamespace {
	ctx := New(d, t)

	ns := &internal.FuncsNamespace{
		Name:    name,
		Context: func(args ...interface{}) interface{} { return ctx },
	}

	ns.AddMethodMapping(ctx.Partial,
		"partial",
		[]string{"include"},
		[][2]string{},
	)

	return ns
}
