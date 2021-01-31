package partial

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new partial Namespace
func New(d *deps.Deps, t *internal.TemplateDeps) *Namespace {
	return &Namespace{
		deps: d,
		tpld: t,
	}
}

// Namespace defines the methods for partials to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	tpld *internal.TemplateDeps
}

const name = "partial"

//  Creates a new Namespace and returns a new internal.FuncsNamespace
func Init(d *deps.Deps, t *internal.TemplateDeps) *internal.FuncsNamespace {
	ctx := New(d, t)

	ns := &internal.FuncsNamespace{
		Name:    name,
		Context: func(args ...interface{}) interface{} { return ctx },
	}

	ns.AddMethodMapping(ctx.Partial,
		"partial",
		nil,
		[][2]string{},
	)

	return ns
}
