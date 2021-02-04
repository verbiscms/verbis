// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new reflect Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for reflect to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "reflect"

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.KindIs,
			"kindIs",
			nil,
			[][2]string{
				{`{{ kindIs "int" 123 }}`, `true`},
			},
		)

		ns.AddMethodMapping(ctx.KindOf,
			"kindOf",
			nil,
			[][2]string{
				{`{{ kindOf 123 }}`, `int`},
			},
		)

		ns.AddMethodMapping(ctx.TypeOf,
			"typeOf",
			nil,
			[][2]string{
				{`{{ typeOf .Post }}`, `domain.PostData`},
			},
		)

		ns.AddMethodMapping(ctx.TypeIs,
			"typeIs",
			nil,
			[][2]string{
				{`{{ trim "    hello verbis     " }}`, `hello verbis`},
			},
		)

		ns.AddMethodMapping(ctx.TypeIsLike,
			"typeIsLike",
			nil,
			[][2]string{
				{`{{ trim "    hello verbis     " }}`, `hello verbis`},
			},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
