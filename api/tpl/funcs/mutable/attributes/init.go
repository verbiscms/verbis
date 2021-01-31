package attributes

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/mutable/auth"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new attributes Namespace
func New(d *deps.Deps, t *internal.TemplateDeps) *Namespace {
	return &Namespace{
		deps: d,
		tpld: t,
		auth: auth.New(d, t),
	}
}

// Namespace defines the methods for attributes to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	tpld *internal.TemplateDeps
	auth *auth.Namespace
}

const name = "attributes"

//  Creates a new Namespace and returns a new internal.FuncsNamespace
func Init(d *deps.Deps, t *internal.TemplateDeps) *internal.FuncsNamespace {
	ctx := New(d, t)

	ns := &internal.FuncsNamespace{
		Name:    name,
		Context: func(args ...interface{}) interface{} { return ctx },
	}

	ns.AddMethodMapping(ctx.Body,
		"body",
		nil,
		[][2]string{
			{`{{ body }}`, `page page-id-1 page-title-my-verbis-page page-template-single page-layout-main`},
		},
	)

	ns.AddMethodMapping(ctx.Lang,
		"lang",
		nil,
		[][2]string{
			{`{{ lang }}`, `en-gb`},
		},
	)

	return ns
}
