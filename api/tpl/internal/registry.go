package internal

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-gonic/gin"
	"html/template"
)

var FuncsNamespaceRegistry []func(d *deps.Deps) *FuncsNamespace

func AddFuncsNamespace(ns func(d *deps.Deps) *FuncsNamespace) {
	FuncsNamespaceRegistry = append(FuncsNamespaceRegistry, ns)
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

func GetFuncMap(d *deps.Deps) template.FuncMap {
	funcMap := template.FuncMap{}

	for _, nsf := range FuncsNamespaceRegistry {
		ns := nsf(d)
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

func GetPostFuncMap(d *deps.Deps, t *TemplateDeps) {

	//ns := fields.Init(d, t)
	//ns.
}

type TemplateDeps struct {
	Context *gin.Context
	Post    *domain.PostData
}
