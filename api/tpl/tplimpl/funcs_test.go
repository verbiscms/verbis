// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tplimpl

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/ainsleyclark/verbis/api/tpl/variables"
	"golang.org/x/net/html"
	"html/template"
)

// Test all config template function mappings
func (t *TplTestSuite) TestFuncs() {
	tm, ctx, post := t.Setup()

	v := variables.Data(tm.deps, ctx, post)
	td := &internal.TemplateDeps{Context: ctx, Post: post, Cfg: nil}
	funcs := tm.FuncMap(ctx, post, nil)

	for _, ns := range tm.getNamespaces(td) {
		for _, mm := range ns.MethodMappings {
			for _, e := range mm.Examples {
				file, err := template.New("test").Funcs(funcs).Parse(e[0])
				if err != nil {
					t.Error(err, fmt.Sprintf("test failed for %s", mm.Name))
					continue
				}

				var tpl bytes.Buffer
				err = file.Execute(&tpl, v)
				if err != nil {
					t.Error(err, fmt.Sprintf("test failed for %s", mm.Name))
				}

				t.Equal(e[1], html.UnescapeString(tpl.String()))
			}
		}
	}
}

func (t *TplTestSuite) TestFuncs_FuncMap() {
	tt := map[string]struct {
		namespaces internal.FuncNamespaces
		want       template.FuncMap
		panics     bool
	}{
		"Success": {
			internal.FuncNamespaces{
				&internal.FuncsNamespace{Name: "namespace", MethodMappings: map[string]internal.FuncMethodMapping{
					"func": {
						Method: nil,
						Name:   "func",
					},
				}},
			},
			template.FuncMap{"func": nil},
			false,
		},
		"Duplicate Func": {
			internal.FuncNamespaces{
				&internal.FuncsNamespace{Name: "namespace", MethodMappings: map[string]internal.FuncMethodMapping{
					"test":    {Method: nil, Name: "replace"},
					"replace": {Method: nil, Name: "replace"},
				}},
			},
			nil,
			true,
		},
		"Duplicate Alias": {
			internal.FuncNamespaces{
				&internal.FuncsNamespace{Name: "namespace", MethodMappings: map[string]internal.FuncMethodMapping{
					"test": {Method: nil, Name: "test", Aliases: []string{"test"}},
				}},
			},
			template.FuncMap{},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			tm, _, _ := t.Setup()
			if test.panics {
				t.Panics(func() {
					tm.getFuncs(test.namespaces)
				})
				return
			}
			t.Equal(test.want, tm.getFuncs(test.namespaces))
		})
	}
}
