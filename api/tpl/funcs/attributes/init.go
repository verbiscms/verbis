package attributes

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new attributes Namespace
func New(d *deps.Deps, t *core.TemplateDeps) *Namespace {
	return &Namespace{
		deps: d,
		post: t.Post,
		//auth: nil,
	}
}

// Namespace defines the methods for attributes to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	post *domain.PostData
	//auth *auth.Namespace
}

const name = "attributes"

//  Creates a new Namespace and returns a new core.FuncsNamespace
func Init(d *deps.Deps, t *core.TemplateDeps) *core.FuncsNamespace {
	ctx := New(d, t)

	ns := &core.FuncsNamespace{
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
