// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new date Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for the os to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "os"

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Env,
			"env",
			nil,
			[][2]string{
				{`{{ env "foo" }}`, `bar`},
			},
		)

		ns.AddMethodMapping(ctx.ExpandEnv,
			"expandEnv",
			nil,
			[][2]string{
				{`{{ expandEnv "Welcome to $foo" }}`, `Welcome to bar`},
			},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
