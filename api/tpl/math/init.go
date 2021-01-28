package math

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new math Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for math to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "math"

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.add,
			"add",
			nil,
			[][2]string{
				{`{{ add 2 2 }}`, `4`},
			},
		)

		ns.AddMethodMapping(ctx.subtract,
			"subtract",
			nil,
			[][2]string{
				{`{{ subtract 100 10 }}`, `90`},
			},
		)

		ns.AddMethodMapping(ctx.subtract,
			"subtract",
			nil,
			[][2]string{
				{`{{ divide 16 4 }}`, `4`},
			},
		)

		ns.AddMethodMapping(ctx.multiply,
			"multiply",
			nil,
			[][2]string{
				{`{{ multiply 4 4 }}`, `16`},
			},
		)

		ns.AddMethodMapping(ctx.modulus,
			"mod",
			[]string{"modulus"},
			[][2]string{
				{`{{ mod 10 9 }}`, `1`},
			},
		)

		ns.AddMethodMapping(ctx.round,
			"round",
			nil,
			[][2]string{
				{`{{ round 10.2 }}`, `10`},
			},
		)

		ns.AddMethodMapping(ctx.ceil,
			"ceil",
			nil,
			[][2]string{
				{`{{ ceil 9.32 }}`, `10`},
			},
		)

		ns.AddMethodMapping(ctx.floor,
			"floor",
			nil,
			[][2]string{
				{`{{ floor 9.62 }}`, `9`},
			},
		)

		ns.AddMethodMapping(ctx.min,
			"min",
			nil,
			[][2]string{
				{`{{ min 20 1 100 }}`, `1`},
			},
		)

		ns.AddMethodMapping(ctx.max,
			"max",
			nil,
			[][2]string{
				{`{{ max 20 1 100 }}`, `100`},
			},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
