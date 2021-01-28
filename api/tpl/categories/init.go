package categories

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new categories Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for slice's for
// template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "categories"

// Adds the namespace methods to the internal.FuncsNamespace slice.
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.getCategory,
			"category",
			nil,
			nil,
		)

		ns.AddMethodMapping(ctx.getCategoryByName,
			"categoryByName",
			nil,
			nil,
		)

		ns.AddMethodMapping(ctx.getCategoryParent,
			"categoryParent",
			nil,
			nil,
		)

		ns.AddMethodMapping(ctx.getCategory,
			"categories",
			nil,
			nil,
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
