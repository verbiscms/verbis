package funcs

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/frontend/attributes"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/frontend/auth"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/frontend/fields"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/frontend/meta"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/frontend/partial"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/frontend/url"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/cast"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/categories"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/date"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/debug"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/dict"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/math"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/media"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/os"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/paths"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/posts"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/rand"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/reflect"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/safe"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/slice"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/strings"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/users"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/generic/util"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/gin-gonic/gin"
	"html/template"
)

// Funcs represents the dependency for interacting with the various
// template functions.
type Funcs struct {
	deps *deps.Deps
	*internal.TemplateDeps
}

// Creates a new Funcs
func New(d *deps.Deps) *Funcs {
	return &Funcs{deps: d}
}

// FuncMap
//
// Returns the frontend and generic funcMap's
func (f *Funcs) FuncMap(ctx *gin.Context, post *domain.PostData) template.FuncMap {
	f.TemplateDeps = &internal.TemplateDeps{
		Context: ctx,
		Post:    post,
	}
	funcs := f.getFuncs(f.getNamespaces())
	f.TemplateDeps.Funcs = funcs
	return funcs
}

// FuncMap
//
// Returns the generic funcMap
func (f *Funcs) GenericFuncMap() template.FuncMap {
	return f.getFuncs(f.getGenericNamespaces())
}

// getFuncs
//
// Loops over the internal.FuncNamespaces passed and returns
// a new template.FuncMap. If duplicates are found for
// either the main method name or an alias, the func
// will panic.
func (f *Funcs) getFuncs(fs internal.FuncNamespaces) template.FuncMap {
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
func (f *Funcs) getNamespaces() internal.FuncNamespaces {
	ns := f.getGenericNamespaces()
	ns = append(ns, f.getFrontendNamespaces()...)
	return ns
}

// getGenericNamespaces
//
// Returns all generic namespaces, ones that are not
// dependant on a post or context.
func (f *Funcs) getGenericNamespaces() internal.FuncNamespaces {
	var fs internal.FuncNamespaces
	for _, nsf := range internal.GenericNamespaceRegistry {
		fs = append(fs, nsf(f.deps))
	}
	return fs
}

// getFrontendNamespaces
//
// Returns all frontend namespaces, ones that are
// dependant on a post or context.
func (f *Funcs) getFrontendNamespaces() internal.FuncNamespaces {
	return internal.FuncNamespaces{
		attributes.Init(f.deps, f.TemplateDeps),
		auth.Init(f.deps, f.TemplateDeps),
		fields.Init(f.deps, f.TemplateDeps),
		meta.Init(f.deps, f.TemplateDeps),
		partial.Init(f.deps, f.TemplateDeps),
		url.Init(f.deps, f.TemplateDeps),
	}
}
