package attributes

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/auth"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new attributes Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for attributes to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	post *domain.PostData
	auth *auth.Namespace
}

const name = "attributes"

// Adds the namespace attributes to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name: name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Body,
			"body",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Lang,
			"lang",
			nil,
			[][2]string{},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
