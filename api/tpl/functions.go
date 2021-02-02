package tpl

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl/funcs"
	"github.com/ainsleyclark/verbis/api/tpl/variables"
	"github.com/gin-gonic/gin"
	"html/template"
)

// TemplateManager defines the service for executing and
// parsing Verbis templates. It's responsible for
// obtaining a template.FuncMap and Data to be
// used within the template.
type TemplateManager struct {
	mapper funcs.Mapper
	data   variables.Reader
}

// Creates a new TemplateManager
func New(d *deps.Deps, ctx *gin.Context, post *domain.PostData) *TemplateManager {
	if post == nil {
		panic("Post data is nil, cannot set up TemplateManager")
	}
	if ctx == nil {
		panic("Post data is nil, cannot set up TemplateManager")
	}
	return &TemplateManager{
		mapper: funcs.New(d, ctx, post),
		data:   variables.New(d, ctx, post),
	}
}

// Funcs
//
// Obtains the template.FuncMap from the registry
func (t *TemplateManager) Funcs() template.FuncMap {
	return t.mapper.FuncMap()
}

// Data
//
// Returns a variables.TemplateData
func (t *TemplateManager) Data() variables.TemplateData {
	return t.data.Get()
}
