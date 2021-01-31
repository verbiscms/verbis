package partial

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/core"
)

// Creates a new partial Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for partials to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "partial"

// Adds the namespace methods to the core.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *core.FuncsNamespace {
		ctx := New(d)

		ns := &core.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Partial,
			"partial",
			nil,
			[][2]string{
				{`{{ slice "hello" "world" "!" }}`, `[hello world !]`},
			},
		)

		return ns
	}

	core.AddFuncsNamespace(f)
}
