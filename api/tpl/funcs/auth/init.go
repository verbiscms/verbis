package auth

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/gin-gonic/gin"
)

// Creates a new auth Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for auth to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	ctx  *gin.Context
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

	internal.AddFuncsNamespace(f)
}
