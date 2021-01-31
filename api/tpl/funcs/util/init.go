package util

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/core"
)

// Creates a new util Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for util to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "safe"

// Adds the namespace methods to the core.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *core.FuncsNamespace {
		ctx := New(d)

		ns := &core.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Len,
			"len",
			nil,
			[][2]string{
				{`{{ len "hello" }}`, `5`},
				{`{{ slice 1 2 3 | len  }}`, `3`},
			},
		)

		ns.AddMethodMapping(ctx.Explode,
			"explode",
			nil,
			[][2]string{
				{`{{ explode "," "hello there !" }}`, `[hello there !]`},
			},
		)

		ns.AddMethodMapping(ctx.Implode,
			"implode",
			nil,
			[][2]string{
				{`{{ slice 1 2 3 | explode "," }}`, `[1 2 3]`},
			},
		)

		return ns
	}

	core.AddFuncsNamespace(f)
}
