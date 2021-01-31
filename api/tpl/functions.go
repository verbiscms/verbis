package tpl

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	_ "github.com/ainsleyclark/verbis/api/tpl/init"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/ainsleyclark/verbis/api/tpl/variables"
	"github.com/gin-gonic/gin"
	"html/template"
)

type TemplateManager struct {
	deps *deps.Deps
	ctx *gin.Context
	post *domain.PostData
}

func New(d *deps.Deps, ctx *gin.Context, post *domain.PostData) *TemplateManager {
	return &TemplateManager{
		deps: d,
		ctx: ctx,
		post: post,
	}
}

func (t *TemplateManager) Funcs() template.FuncMap {
	return internal.GetFuncMap(t.deps, t.ctx, t.post)
}

func (t *TemplateManager) Data() variables.TemplateData {
	return variables.TemplateData{}
}