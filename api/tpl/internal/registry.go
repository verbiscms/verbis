package internal

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"html/template"
)

var TemplateFuncsNamespaceRegistry []func(d *deps.Deps) *TemplateFuncsNamespace

func AddTemplateFuncsNamespace(ns func(d *deps.Deps) *TemplateFuncsNamespace) {
	TemplateFuncsNamespaceRegistry = append(TemplateFuncsNamespaceRegistry, ns)
}

type TemplateFuncsNamespace struct {
	Name           string
	Context        func(v ...interface{}) interface{}
	MethodMappings map[string]TemplateFuncMethodMapping
}

type TemplateFuncMethodMapping struct {
	Method   interface{}
	Name     string
	Aliases  []string
	Examples [][2]string
}

func (t *TemplateFuncsNamespace) AddMethodMapping(m interface{}, name string, aliases []string, examples [][2]string) {
	if t.MethodMappings == nil {
		t.MethodMappings = make(map[string]TemplateFuncMethodMapping)
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

	t.MethodMappings[name] = TemplateFuncMethodMapping{
		Method:   m,
		Name: name,
		Aliases:  aliases,
		Examples: examples,
	}
}


func GetFuncMap(d *deps.Deps) template.FuncMap {
	funcMap := template.FuncMap{}

	for _, nsf := range TemplateFuncsNamespaceRegistry {
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
