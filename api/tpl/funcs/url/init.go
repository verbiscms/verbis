// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package url

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/gin-gonic/gin"
)

// Creates a new reflect Namespace
func New(d *deps.Deps, t *internal.TemplateDeps) *Namespace {
	return &Namespace{
		deps: d,
		ctx:  t.Context,
	}
}

// Namespace defines the methods for reflect to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	ctx  *gin.Context
}

const name = "url"

//  Creates a new Namespace and returns a new internal.FuncsNamespace
func Init(d *deps.Deps, t *internal.TemplateDeps) *internal.FuncsNamespace {
	ctx := New(d, t)

	ns := &internal.FuncsNamespace{
		Name:    name,
		Context: func(args ...interface{}) interface{} { return ctx },
	}

	ns.AddMethodMapping(ctx.Base,
		"baseUrl",
		nil,
		[][2]string{
			{`{{ baseUrl }}`, `http://verbiscms.com`},
		},
	)

	ns.AddMethodMapping(ctx.Scheme,
		"scheme",
		nil,
		[][2]string{
			{`{{ scheme }}`, `http`},
		},
	)

	ns.AddMethodMapping(ctx.Host,
		"host",
		nil,
		[][2]string{
			{`{{ host }}`, `verbiscms.com`},
		},
	)

	ns.AddMethodMapping(ctx.Full,
		"fullUrl",
		nil,
		[][2]string{
			{`{{ fullUrl }}`, `http://verbiscms.com/page`},
		},
	)

	ns.AddMethodMapping(ctx.Path,
		"path",
		nil,
		[][2]string{
			{`{{ path }}`, `/page`},
		},
	)

	ns.AddMethodMapping(ctx.Query,
		"query",
		nil,
		[][2]string{
			{`{{ query "foo" }}`, `bar`},
		},
	)

	ns.AddMethodMapping(ctx.Pagination,
		"paginationPage",
		nil,
		[][2]string{
			{`{{ paginationPage }}`, `2`},
		},
	)

	return ns
}
