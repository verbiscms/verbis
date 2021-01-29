package cast

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/spf13/cast"
)

// Creates a new cast Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for cast to be used
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

		ns.AddMethodMapping(cast.ToBool,
			"toBool",
			nil,
			[][2]string{
				{`{{ toBool "true" }}`, `true`},
			},
		)

		ns.AddMethodMapping(cast.ToString,
			"toString",
			nil,
			[][2]string{
				{`{{ toString 1 }}`, `1`},
			},
		)

		ns.AddMethodMapping(ctx.ToSlice,
			"toSlice",
			nil,
			[][2]string{
				{`{{ toSlice "a" }}`, `["a"]`},
				{`{{ toSlice 1 }}`, `[1]`},
			},
		)

		ns.AddMethodMapping(cast.ToTime,
			"toTime",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(cast.ToDuration,
			"toTime",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(cast.ToInt,
			"toInt",
			nil,
			[][2]string{
				{`{{ toInt "1" }}`, `1`},
			},
		)

		ns.AddMethodMapping(cast.ToInt8,
			"toInt8",
			nil,
			[][2]string{
				{`{{ toInt8 "1" }}`, `1`},
			},
		)

		ns.AddMethodMapping(cast.ToInt16,
			"toInt16",
			nil,
			[][2]string{
				{`{{ toInt16 "1" }}`, `1`},
			},
		)

		ns.AddMethodMapping(cast.ToInt32,
			"toInt32",
			nil,
			[][2]string{
				{`{{ toInt32 "1" }}`, `1`},
			},
		)

		ns.AddMethodMapping(cast.ToInt64,
			"toInt64",
			nil,
			[][2]string{
				{`{{ toInt64 "1" }}`, `1`},
			},
		)

		ns.AddMethodMapping(cast.ToUint,
			"toUInt",
			nil,
			[][2]string{
				{`{{ toUInt "1" }}`, `1`},
			},
		)

		ns.AddMethodMapping(cast.ToUint8,
			"toUInt8",
			nil,
			[][2]string{
				{`{{ toUInt8 "1" }}`, `1`},
			},
		)

		ns.AddMethodMapping(cast.ToUint16,
			"toUInt16",
			nil,
			[][2]string{
				{`{{ toUInt16 "1" }}`, `1`},
			},
		)

		ns.AddMethodMapping(cast.ToUint32,
			"toUInt32",
			nil,
			[][2]string{
				{`{{ toUInt32 "1" }}`, `1`},
			},
		)

		ns.AddMethodMapping(cast.ToUint64,
			"toUInt64",
			nil,
			[][2]string{
				{`{{ toUInt64 "1" }}`, `1`},
			},
		)

		ns.AddMethodMapping(cast.ToFloat32,
			"toFloat32",
			nil,
			[][2]string{
				{`{{ toFloat32 "1.1" }}`, `1.1`},
			},
		)

		ns.AddMethodMapping(cast.ToFloat64E,
			"toFloat64",
			nil,
			[][2]string{
				{`{{ toFloat64 "1.1" }}`, `1.1`},
			},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
