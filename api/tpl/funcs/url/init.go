package url

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/core"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

// Creates a new reflect Namespace
func New(d *deps.Deps, ctx *gin.Context) *Namespace {
	return &Namespace{
		deps: d,
		ctx:  ctx,
	}
}

// Namespace defines the methods for reflect to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
	ctx  *gin.Context
}

const name = "reflect"

// Adds the namespace methods to the core.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps, t *core.TemplateDeps) *core.FuncsNamespace {
		ctx := New(d, t.Context)

		ns := &core.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Base,
			"baseUrl",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Scheme,
			"scheme",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Host,
			"host",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Full,
			"fullUrl",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Path,
			"path",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Query,
			"query",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Pagination,
			"paginationPage",
			nil,
			[][2]string{},
		)

		return ns
	}

	color.Green.Println(f)

	//core.AddFuncsNamespace(f)
}
