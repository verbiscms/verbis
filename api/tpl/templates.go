package tpl

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
)


type TemplateHandler interface {
	TemplateFuncGetter
	TemplateDataGetter
	Prepare(c TemplateConfig) TemplateExecutor
}

type TemplateExecutor interface {
	Exists(template string) bool
	Execute(w io.Writer, name string, data interface{}) error
	ExecutePost(w io.Writer, name string, ctx *gin.Context, post *domain.PostData) error
}

// TemplateMapper represents the functions for obtaining template.FuncMap's
// for use in Verbis templates.
type TemplateFuncGetter interface {
	FuncMap(ctx *gin.Context, post *domain.PostData) template.FuncMap
	GenericFuncMap() template.FuncMap
}

// TemplateData represents the functions for obtaining template.FuncMap's
// for use in Verbis templates.
type TemplateDataGetter interface {
	Data(ctx *gin.Context, post *domain.PostData) interface{}
}

type TemplateConfig interface {
	GetRoot() string
	GetExtension() string
	GetMaster() string
}

// Config represents the options for passing
type Config struct {
	Root      string //view root
	Extension string //template extension
	Master    string //template master
}

func (c Config) GetRoot() string {
	return c.Root
}

func (c Config) GetExtension() string {
	return c.Extension
}

func (c Config) GetMaster() string {
	return c.Master
}


