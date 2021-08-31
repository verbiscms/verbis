// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/services/fields"
	"github.com/verbiscms/verbis/api/tpl/internal"
)

// New creates a new fields Namespace.
func New(d *deps.Deps, t *internal.TemplateDeps) *Namespace {
	f := fields.NewService(d, t.Post)
	return &Namespace{
		deps:   d,
		fields: f,
	}
}

// Namespace defines the methods for fields to be used
// as template functions.
type Namespace struct {
	deps   *deps.Deps
	fields fields.FieldService
}

// name defines the identifier for the namespace.
const name = "fields"

// Init creates a new Namespace and returns a new
// internal.FuncsNamespace.
func Init(d *deps.Deps, t *internal.TemplateDeps) *internal.FuncsNamespace {
	ctx := New(d, t)

	ns := &internal.FuncsNamespace{
		Name:    name,
		Context: func(args ...interface{}) interface{} { return ctx },
	}

	ns.AddMethodMapping(ctx.fields.GetField,
		"field",
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetFieldObject,
		"fieldObject",
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetFields,
		"fields",
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetLayout,
		"layout",
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetLayouts,
		"layouts",
		[]string{},
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetRepeater,
		"repeater",
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetFlexible,
		"flexible",
		nil,
		[][2]string{},
	)

	return ns
}
