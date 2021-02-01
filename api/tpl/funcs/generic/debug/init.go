package debug

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new debug Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for debug to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "paths"

// Adds the namespace debug to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Debug,
			"debug",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Dump,
			"dump",
			nil,
			[][2]string{},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
