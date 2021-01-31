package reflect

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/core"
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

// Adds the namespace methods to the core.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *core.FuncsNamespace {
		ctx := New(d)

		ns := &core.FuncsNamespace{
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

	core.AddFuncsNamespace(f)
}
