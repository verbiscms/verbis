// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/tpl/internal"
)

// New creates a new categories Namespace.
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for categories to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

// name defines the identifier for the namespace.
const name = "categories"

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Find,
			"category",
			nil,
			nil,
		)

		ns.AddMethodMapping(ctx.ByName,
			"categoryByName",
			nil,
			nil,
		)

		ns.AddMethodMapping(ctx.Parent,
			"categoryParent",
			nil,
			nil,
		)

		ns.AddMethodMapping(ctx.List,
			"categories",
			nil,
			nil,
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
