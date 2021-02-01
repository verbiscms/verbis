package meta

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new meta Namespace
func New(d *deps.Deps, t *internal.TemplateDeps) *Namespace {
	return &Namespace{
		deps: d,
		post: t.Post,
	}
}

// Namespace defines the methods for meta to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	post *domain.PostData
}

const name = "meta"

//  Creates a new Namespace and returns a new internal.FuncsNamespace
func Init(d *deps.Deps, t *internal.TemplateDeps) *internal.FuncsNamespace {
	ctx := New(d, t)

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
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.Footer,
		"verbisFoot",
		[]string{"foot"},
		[][2]string{},
	)

	return ns
}
