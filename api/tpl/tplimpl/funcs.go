package tplimpl

import (
	// Context/Post specific functions
	"github.com/ainsleyclark/verbis/api/tpl/funcs/attributes"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/auth"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/fields"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/meta"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/partial"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/url"
	"github.com/ainsleyclark/verbis/api/tpl/internal"

	// Generic functions
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/cast"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/categories"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/date"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/debug"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/dict"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/math"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/media"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/os"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/paths"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/posts"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/rand"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/reflect"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/safe"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/slice"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/strings"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/users"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/util"

	// Package
	"html/template"
)

// getFuncs
//
// Loops over the internal.FuncNamespaces passed and returns
// a new template.FuncMap. If duplicates are found for
// either the main method name or an alias, the func
// will panic.
func (t *TemplateManager) getFuncs(fs internal.FuncNamespaces) template.FuncMap {
	funcMap := template.FuncMap{}

	for _, ns := range fs {
		for _, mm := range ns.MethodMappings {
			if _, exists := funcMap[mm.Name]; exists {
				panic(ns.Name + " is a duplicate template func")
			}
			funcMap[mm.Name] = mm.Method
			for _, alias := range mm.Aliases {
				if _, exists := funcMap[alias]; exists {
					panic(alias + " is a duplicate template func")
				}
				funcMap[alias] = mm.Method
			}
		}
	}

	return funcMap
}

// getNamespaces
//
// Merges the generic and frontend namespaces and returns
// a slice of namespaces, both generic and frontend.
func (t *TemplateManager) getNamespaces(td *internal.TemplateDeps) internal.FuncNamespaces {
	ns := t.getGenericNamespaces()
	ns = append(ns, t.getFrontendNamespaces(td)...)
	return ns
}

// getGenericNamespaces
//
// Returns all generic namespaces, ones that are not
// dependant on a post or context.
func (t *TemplateManager) getGenericNamespaces() internal.FuncNamespaces {
	var fs internal.FuncNamespaces
	for _, nsf := range internal.GenericNamespaceRegistry {
		fs = append(fs, nsf(t.deps))
	}
	return fs
}

// getFrontendNamespaces
//
// Returns all frontend namespaces, ones that are
// dependant on a post or context.
func (t *TemplateManager) getFrontendNamespaces(td *internal.TemplateDeps) internal.FuncNamespaces {
	return internal.FuncNamespaces{
		attributes.Init(t.deps, td),
		auth.Init(t.deps, td),
		fields.Init(t.deps, td),
		meta.Init(t.deps, td),
		partial.Init(t.deps, td),
		url.Init(t.deps, td),
	}
}
