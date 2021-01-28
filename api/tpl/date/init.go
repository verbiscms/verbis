package date

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"time"
)

// Creates a new date Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for dates to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "date"

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(time.Now(),
			"now",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.date,
			"date",
			nil,
			[][2]string{
				{`{{ date "02 Jan 2006" 643408779 }}`, `22 May 1990`},
			},
		)

		ns.AddMethodMapping(ctx.dateInZone,
			"dateInZone",
			nil,
			[][2]string{
				{`{{ dateInZone "02/01/2006" 643408779 "Europe/London" }}`, `22 May 1990`},
			},
		)

		ns.AddMethodMapping(ctx.ago,
			"ago",
			nil,
			[][2]string{},
		)

		ns.AddMethodMapping(ctx.duration,
			"duration",
			nil,
			[][2]string{
				{`{{ duration "85" }}`, `1m25s`},
			},
		)

		ns.AddMethodMapping(ctx.htmlDate,
			"htmlDate",
			nil,
			[][2]string{
				{`{{ htmlDate 643408779 }}`, `1990-05-22`},
			},
		)

		ns.AddMethodMapping(ctx.htmlDateInZone,
			"htmlDateInZone",
			nil,
			[][2]string{
				{`{{ htmlDateInZone 643408779 "GMT" }}`, `1990-05-22`},
			},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
