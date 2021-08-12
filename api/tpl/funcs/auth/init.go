// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/tpl/internal"
)

// New creates a new auth Namespace.
func New(d *deps.Deps, t *internal.TemplateDeps) *Namespace {
	return &Namespace{
		deps: d,
		ctx:  t.Context,
	}
}

// Namespace defines the methods for auth to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	ctx  *gin.Context
}

// name defines the identifier for the namespace.
const name = "auth"

// Init creates a new Namespace and returns a new
// internal.FuncsNamespace
func Init(d *deps.Deps, t *internal.TemplateDeps) *internal.FuncsNamespace {
	ctx := New(d, t)

	ns := &internal.FuncsNamespace{
		Name: name,
		Context: func(args ...interface{}) interface{} {
			return ctx
		},
	}

	ns.AddMethodMapping(ctx.Auth,
		"auth",
		nil,
		[][2]string{
			{`{{ toBool "true" }}`, `true`},
		},
	)

	ns.AddMethodMapping(ctx.Admin,
		"admin",
		nil,
		[][2]string{
			{`{{ auth }}`, `false`},
		},
	)

	return ns
}
