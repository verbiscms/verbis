// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
)

// GenericNamespaceRegistry represents the slice of generic
// functions that provide namespaces.
var GenericNamespaceRegistry []func(d *deps.Deps) *FuncsNamespace

// GenericNamespaceRegistry represents the slice of generic
// functions that provide namespaces.
type FuncNamespaces []*FuncsNamespace

// TemplateDeps represents the data to be passed to templates
// that rely on either context or domain.PostData such as
// "url", "fields" or "meta".
type TemplateDeps struct {
	// The context to be used used for obtaining URL's & query parameters etc...
	Context *gin.Context
	// The post to be used for rendering meta information for the page
	Post *domain.PostData
	// The config of the executor used in partials to obtain the root path.
	Cfg tpl.TemplateConfig
}

// FuncsNamespace represents a template function namespace.
type FuncsNamespace struct {
	// The name of the namespace, for example "math" or "slice"
	Name string
	// The method receiver of the namespace
	Context func(v ...interface{}) interface{}
	// Additional information about the namespace such as aliases and examples.
	MethodMappings map[string]FuncMethodMapping
}

// FuncMethodMapping represents individual methods found in
// each template namespaces.
type FuncMethodMapping struct {
	Method   interface{}
	Name     string
	Aliases  []string
	Examples [][2]string
}

// AddFuncsNamespace
//
// Appends a FuncsNamespace to the registry
func AddFuncsNamespace(ns func(d *deps.Deps) *FuncsNamespace) {
	GenericNamespaceRegistry = append(GenericNamespaceRegistry, ns)
}

// AddMethodMapping
//
// Adds a FuncsNamespace to the GenericNamespaceRegistry
// If any duplicates are found in the registry a panic
// will occur.
func (t *FuncsNamespace) AddMethodMapping(m interface{}, name string, aliases []string, examples [][2]string) {
	if t.MethodMappings == nil {
		t.MethodMappings = make(map[string]FuncMethodMapping)
	}

	for _, e := range examples {
		if e[0] == "" {
			panic(t.Name + ": Empty example for " + name)
		}
	}

	for _, a := range aliases {
		if a == "" {
			panic(t.Name + ": Empty alias for " + name)
		}
	}

	t.MethodMappings[name] = FuncMethodMapping{
		Method:   m,
		Name:     name,
		Aliases:  aliases,
		Examples: examples,
	}
}
