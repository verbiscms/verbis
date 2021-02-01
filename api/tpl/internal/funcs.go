package internal

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-gonic/gin"
	"html/template"
)

var GenericNamespaceRegistry []func(d *deps.Deps) *FuncsNamespace

type FuncNamespaces []*FuncsNamespace

type TemplateDeps struct {
	Context *gin.Context
	Post    *domain.PostData
	Funcs   template.FuncMap
}

func AddFuncsNamespace(ns func(d *deps.Deps) *FuncsNamespace) {
	GenericNamespaceRegistry = append(GenericNamespaceRegistry, ns)
}

type FuncsNamespace struct {
	Name           string
	Context        func(v ...interface{}) interface{}
	MethodMappings map[string]FuncMethodMapping
}

type FuncMethodMapping struct {
	Method   interface{}
	Name     string
	Aliases  []string
	Examples [][2]string
}

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
