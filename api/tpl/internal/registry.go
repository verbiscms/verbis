package internal

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/core"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/meta"
	"github.com/gin-gonic/gin"
	"html/template"
)


func getGenericMap(d *deps.Deps) template.FuncMap {
	funcMap := template.FuncMap{}

	for _, nsf := range core.GenericNamespaceRegistry {
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



func GetFuncMap(d *deps.Deps, ctx *gin.Context, post *domain.PostData) template.FuncMap {
	funcMap := getGenericMap(d)

	t := &core.TemplateDeps{
		Context: ctx,
		Post:    post,
	}

	nss := core.MutableNamespaceRegistry{
		//attributes.Init(d, t),
		//auth.Init(d, t),
		//fields.Init(d, t),
		meta.Init(d, t),
	}

	for _, ns := range nss {
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
