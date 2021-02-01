package rand

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
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

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Int,
			"randInt",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Float,
			"randFloat",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Alpha,
			"randAlpha",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.AlphaNum,
			"randAlphaNum",
			nil,
			[][2]string{},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
