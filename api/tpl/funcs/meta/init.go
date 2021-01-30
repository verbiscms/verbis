package meta

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new meta Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for meta to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	post  *domain.PostData
}

const name = "safe"

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name: name,
			Context: func(args ...interface{}) interface{} {

				return ctx

			},
		}

		ns.AddMethodMapping(ctx.Header,
			"verbisHead",
			[]string{"head"},
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.MetaTitle,
			"metaTitle",
			[]string{"foot"},
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Footer,
			"verbisFoot",
			[]string{"foot"},
			[][2]string{},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
