package users

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/core"
)

// Creates a new users Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for users to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "users"

// Adds the namespace methods to the core.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *core.FuncsNamespace {
		ctx := New(d)

		ns := &core.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Find,
			"user",
			nil,
			nil,
		)

		ns.AddMethodMapping(ctx.List,
			"users",
			nil,
			nil,
		)

		return ns
	}

	core.AddFuncsNamespace(f)
}
