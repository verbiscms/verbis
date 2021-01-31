package slice

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/core"
)

// Creates a new slice Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for slices to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "slice"

// Adds the namespace methods to the core.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *core.FuncsNamespace {
		ctx := New(d)

		ns := &core.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Slice,
			"slice",
			nil,
			[][2]string{
				{`{{ slice "hello" "world" "!" }}`, `[hello world !]`},
			},
		)

		ns.AddMethodMapping(ctx.Append,
			"append",
			nil,
			[][2]string{
				{`{{ append (slice "hello" "world" "!") "verbis" }}`, `[hello world ! verbis]`},
			},
		)

		ns.AddMethodMapping(ctx.Prepend,
			"prepend",
			nil,
			[][2]string{
				{`{{ prepend (slice "hello" "world" "!") "verbis" }}`, `[verbis hello world !]`},
			},
		)

		ns.AddMethodMapping(ctx.First,
			"first",
			nil,
			[][2]string{
				{`{{ first (slice "hello" "world" "!") }}`, `hello`},
			},
		)

		ns.AddMethodMapping(ctx.Last,
			"last",
			nil,
			[][2]string{
				{`{{ last (slice "hello" "world" "!") }}`, `!`},
			},
		)

		ns.AddMethodMapping(ctx.Reverse,
			"reverse",
			nil,
			[][2]string{
				{`{{ reverse (slice "hello" "world" "!") }}`, `[! world hello]`},
			},
		)

		return ns
	}

	core.AddFuncsNamespace(f)
}
