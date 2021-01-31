package funcs

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
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
	"github.com/ainsleyclark/verbis/api/tpl/funcs/mutable/attributes"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/mutable/auth"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/mutable/fields"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/mutable/meta"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/mutable/partial"
	_ "github.com/ainsleyclark/verbis/api/tpl/funcs/mutable/partial"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/mutable/url"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/gin-gonic/gin"
	"html/template"
)

type Mapper interface {
	Map() template.FuncMap
}

type Funcs struct {
	deps *deps.Deps
	ctx  *gin.Context
	post *domain.PostData
}

// Creates a new Funcs
func New(d *deps.Deps, ctx *gin.Context, post *domain.PostData) *Funcs {
	return &Funcs{
		deps: d,
		ctx:  ctx,
		post: post,
	}
}

func (f *Funcs) Map() template.FuncMap {
	funcMap := template.FuncMap{}

	for _, ns := range f.getNamespaces() {
		if _, exists := funcMap[ns.Name]; exists {
			panic(ns.Name + " is a duplicate template func")
		}
		for _, mm := range ns.MethodMappings {
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

func (f *Funcs) getNamespaces() []*internal.FuncsNamespace {
	var fs []*internal.FuncsNamespace
	for _, nsf := range internal.GenericNamespaceRegistry {
		ns := nsf(f.deps)
		fs = append(fs, ns)
	}

	for _, nsf := range f.getMutableNamespaces() {
		fs = append(fs, nsf)
	}

	return fs
}

func (f *Funcs) getMutableNamespaces() internal.MutableNamespaceRegistry {
	t := &internal.TemplateDeps{
		Context: f.ctx,
		Post:    f.post,
	}
	return internal.MutableNamespaceRegistry{
		attributes.Init(f.deps, t),
		auth.Init(f.deps, t),
		fields.Init(f.deps, t),
		meta.Init(f.deps, t),
		partial.Init(f.deps, t),
		url.Init(f.deps, t),
	}
}

