package safe

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new safe Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for safe to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "safe"

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.HTML,
			"safeHTML",
			nil,
			[][2]string{
				{`{{ "<p>verbis&cms</p>" | safeHTML }}`, `verbis&cms`},
			},
		)

		ns.AddMethodMapping(ctx.HTMLAttr,
			"safeHTMLAttr",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.CSS,
			"safeCSS",
			nil,
			[][2]string{
				{`{{ "<p>verbis&cms</p>" | safeHTML }}`, `verbis&amp;cms`},
			},
		)

		ns.AddMethodMapping(ctx.JS,
			"safeJS",
			nil,
			[][2]string{
				{`{{ "(2*2)" | safeJS }}`, `(2*2)`},
			},
		)

		ns.AddMethodMapping(ctx.JSStr,
			"safeJSStr",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.Url,
			"safeURL",
			nil,
			[][2]string{
				{`{{ "https://verbiscms.com" | safeUrl }}`, `https://verbiscms.com`},
			},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
