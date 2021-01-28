package dict

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new date Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for dicts to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "dict"

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.dict,
			"dict",
			nil,
			[][2]string{
				{`{{ dict "colour" "green" "height" 20 }}`, `map[colour:green height:20]`},
			},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
