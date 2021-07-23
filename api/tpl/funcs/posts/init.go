// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/tpl/internal"
)

// Creates a new posts Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for posts to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "posts"

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
			"post",
			nil,
			nil,
		)

		ns.AddMethodMapping(ctx.List,
			"posts",
			nil,
			nil,
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
