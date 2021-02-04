// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"strings"
)

// Creates a new strings Namespace
func New(d *deps.Deps) *Namespace {
	return &Namespace{deps: d}
}

// Namespace defines the methods for slices to be used
// as template functions.
type Namespace struct {
	deps *deps.Deps
}

const name = "strings"

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func init() {
	f := func(d *deps.Deps) *internal.FuncsNamespace {
		ctx := New(d)

		ns := &internal.FuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(strings.TrimSpace,
			"trim",
			nil,
			[][2]string{
				{`{{ trim "    hello verbis     " }}`, `hello verbis`},
			},
		)

		ns.AddMethodMapping(strings.ToUpper,
			"upper",
			nil,
			[][2]string{
				{`{{ upper "hello verbis" }}`, `HELLO VERBIS`},
			},
		)

		ns.AddMethodMapping(strings.ToLower,
			"lower",
			nil,
			[][2]string{
				{`{{ lower "hELLo VERBIS" }}`, `hello verbis`},
			},
		)

		ns.AddMethodMapping(strings.Title,
			"title",
			nil,
			[][2]string{
				{`{{ title "hello verbis" }}`, `Hello Verbis`},
			},
		)

		ns.AddMethodMapping(ctx.Replace,
			"replace",
			nil,
			[][2]string{
				{`{{ replace " " "-" "hello verbis cms" }}`, `hello-verbis-cms`},
			},
		)

		ns.AddMethodMapping(ctx.Substr,
			"substr",
			nil,
			[][2]string{
				{`{{ substr "hello verbis" 0 5 }}`, `hello`},
			},
		)

		ns.AddMethodMapping(ctx.Trunc,
			"trunc",
			nil,
			[][2]string{
				{`{{ trunc "hello verbis" 5 }}`, `hello`},
				{`{{ trunc "hello verbis" -6 }}`, `verbis`},
			},
		)

		ns.AddMethodMapping(ctx.Ellipsis,
			"ellipsis",
			nil,
			[][2]string{
				{`{{ ellipsis "hello verbis cms!" 12 }}`, `hello verbis...`},
			},
		)

		ns.AddMethodMapping(ctx.Match,
			"regexMatch",
			nil,
			[][2]string{
				{`{{ regexMatch "^Verbis" "Verbis CMS" }}`, `true`},
			},
		)

		ns.AddMethodMapping(ctx.FindAll,
			"regexFindAll",
			nil,
			[][2]string{
				{`{{ regexFindAll "[1,3,5,7]" "123456789" -1 }}`, `[1 3 5 7]`},
			},
		)

		ns.AddMethodMapping(ctx.Find,
			"regexFind",
			nil,
			[][2]string{
				{`{{ regexFind "verbis.?" "verbiscms" }}`, `verbisc`},
			},
		)

		ns.AddMethodMapping(ctx.ReplaceAll,
			"regexReplaceAll",
			nil,
			[][2]string{
				{`{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "${1}W" }}`, `-W-xxW-`},
			},
		)

		ns.AddMethodMapping(ctx.ReplaceAllLiteral,
			"regexReplaceAllLiteral",
			nil,
			[][2]string{
				{`{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}" }}`, `-${1}-${1}-`},
			},
		)

		ns.AddMethodMapping(ctx.Split,
			"regexSplit",
			nil,
			[][2]string{
				{`{{ regexSplit "b+" "verbis" -1 }}`, `[ver is]`},
			},
		)

		ns.AddMethodMapping(ctx.QuoteMeta,
			"regexQuoteMeta",
			nil,
			[][2]string{
				{`{{ regexQuoteMeta "verbis+?" }}`, "verbis\\+\\?"},
			},
		)

		return ns
	}

	internal.AddFuncsNamespace(f)
}
