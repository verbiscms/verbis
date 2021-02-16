// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tplimpl

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/ainsleyclark/verbis/api/tpl/variables"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Setup(t *testing.T) (*TemplateManager, *gin.Context, *domain.PostData) {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(rr)
	ctx.Request, _ = http.NewRequest("GET", "/page", nil)
	engine.Use(location.Default())

	engine.GET("/page", func(g *gin.Context) {
		ctx = g
		return
	})

	req, err := http.NewRequest("GET", "http://verbiscms.com/page?page=2&foo=bar", nil)
	assert.NoError(t, err)
	engine.ServeHTTP(rr, req)

	os.Setenv("foo", "bar")

	post := &domain.PostData{
		Post: domain.Post{
			Id:           1,
			Slug:         "/page",
			Title:        "My Verbis Page",
			Status:       "published",
			Resource:     nil,
			PageTemplate: "single",
			PageLayout:   "main",
			UserId:       1,
		},
		Fields: []domain.PostField{
			{PostId: 1, Type: "text", Name: "text", Key: "", OriginalValue: "Hello World!"},
		},
	}

	d := &deps.Deps{
		Store:   nil,
		Config:  nil,
		Site:    domain.Site{},
		Options: &domain.Options{
			GeneralLocale: "en-gb",
		},
		Paths:   deps.Paths{},
		Theme:   &domain.ThemeConfig{},
		Running: false,
	}

	return &TemplateManager{deps: d}, ctx, post
}

// Test all internal template function mappings
func TestFuncs(t *testing.T) {
	tm, ctx, post := Setup(t)

	v := variables.Data(tm.deps, ctx, post)
	td := &internal.TemplateDeps{Context: ctx, Post:    post, Cfg:     nil,}
	funcs := tm.FuncMap(ctx, post, nil)

	for _, ns := range tm.getNamespaces(td) {
		for _, mm := range ns.MethodMappings {
			for _, e := range mm.Examples {
				file, err := template.New("test").Funcs(funcs).Parse(e[0])
				if err != nil {
					t.Errorf("test failed for %s: %s", mm.Name, err.Error())
					continue
				}

				var tpl bytes.Buffer
				err = file.Execute(&tpl, v)
				if err != nil {
					t.Error(err)
				}

				assert.Equal(t, e[1], html.UnescapeString(tpl.String()))
			}
		}
	}
}

func TestFuncs_FuncMap(t *testing.T) {

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
		t.Run(name, func(t *testing.T) {
			tm, _, _ := Setup(t)

			if test.panics {
				assert.Panics(t, func() {
					tm.getFuncs(test.namespaces)
				})
				return
			}

			assert.Equal(t, test.want, tm.getFuncs(test.namespaces))
		})
	}
}
