package slice

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

type Namespace struct {
	deps *deps.Deps
}

const name = "slice"

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		ctx := New(d)

		ns := &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.slice,
			"slice",
			nil,
			[][2]string{
				{`{{ $mySlice := slice 1 2 3 }}`, `1234`},
			},
		)

		ns.AddMethodMapping(ctx.append,
			"append",
			nil,
			[][2]string{
				{`{{ $mySlice := slice 1 2 3 }}`, `<p>Blockhead</p>`},
			},
		)

		ns.AddMethodMapping(ctx.prepend,
			"prepend",
			nil,
			[][2]string{
				{`{{chomp "<p>Blockhead</p>\n" | safeHTML }}`, `<p>Blockhead</p>`},
			},
		)

		ns.AddMethodMapping(ctx.first,
			"first",
			nil,
			[][2]string{
				{`{{chomp "<p>Blockhead</p>\n" | safeHTML }}`, `<p>Blockhead</p>`},
			},
		)

		ns.AddMethodMapping(ctx.last,
			"last",
			nil,
			[][2]string{
				{`{{chomp "<p>Blockhead</p>\n" | safeHTML }}`, `<p>Blockhead</p>`},
			},
		)

		ns.AddMethodMapping(ctx.reverse,
			"reverse",
			nil,
			[][2]string{
				{`{{ $mySlice := slice 1 2 3 }}{{ reverse $mySlice }}`, `321`},
			},
		)

		return ns
	}

	internal.AddTemplateFuncsNamespace(f)
}
