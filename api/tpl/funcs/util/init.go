package util

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
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

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
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

		ns.AddMethodMapping(ctx.ToSlice,
			"toSlice",
			nil,
			[][2]string{
				{`{{ toSlice 1 }}`, `[1]`},
			},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
