package auth

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/gin-gonic/gin"
)

// Creates a new auth Namespace
func New(d *deps.Deps, t *core.TemplateDeps) *Namespace {
	return &Namespace{
		deps: d,
		ctx: t.Context,
	}
}

// Namespace defines the methods for auth to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	ctx  *gin.Context
}

const name = "auth"

//  Creates a new Namespace and returns a new core.FuncsNamespace
func Init(d *deps.Deps, t *core.TemplateDeps) *core.FuncsNamespace {
	ctx := New(d, t)

	ns := &core.FuncsNamespace{
		Name: name,
		Context: func(args ...interface{}) interface{} {
			return ctx
		},
	}

	ns.AddMethodMapping(ctx.Auth,
		"auth",
		nil,
		[][2]string{
			{`{{ toBool "true" }}`, `true`},
		},
	)

	ns.AddMethodMapping(ctx.Admin,
		"admin",
		nil,
		[][2]string{
			{`{{ auth }}`, `false`},
		},
	)

	return ns
}
