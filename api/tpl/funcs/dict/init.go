package dict

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/core"
)

// Creates a new dict Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for dicts to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "dict"

// Adds the namespace methods to the core.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *core.FuncsNamespace {
		ctx := New(d)

		ns := &core.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Dict,
			"dict",
			nil,
			[][2]string{
				{`{{ dict "colour" "green" "height" 20 }}`, `map[colour:green height:20]`},
			},
		)

		return ns
	}

	core.AddFuncsNamespace(f)
}
