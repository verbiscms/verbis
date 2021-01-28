package slice

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
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

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.slice,
			"slice",
			nil,
			[][2]string{
				{`{{ slice "hello" "world" "!" }}`, `[hello world !]`},
			},
		)

		ns.AddMethodMapping(ctx.append,
			"append",
			nil,
			[][2]string{
				{`{{ append (slice "hello" "world" "!") "verbis" }}`, `[hello world ! verbis]`},
			},
		)

		ns.AddMethodMapping(ctx.prepend,
			"prepend",
			nil,
			[][2]string{
				{`{{ prepend (slice "hello" "world" "!") "verbis" }}`, `[verbis hello world !]`},
			},
		)

		ns.AddMethodMapping(ctx.first,
			"first",
			nil,
			[][2]string{
				{`{{ first (slice "hello" "world" "!") }}`, `hello`},
			},
		)

		ns.AddMethodMapping(ctx.last,
			"last",
			nil,
			[][2]string{
				{`{{ last (slice "hello" "world" "!") }}`, `!`},
			},
		)

		ns.AddMethodMapping(ctx.reverse,
			"reverse",
			nil,
			[][2]string{
				{`{{ reverse (slice "hello" "world" "!") }}`, `[! world hello]`},
			},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
