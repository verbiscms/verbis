package tplimpl

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/ainsleyclark/verbis/api/tpl/variables"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"sync"
)

// Prepare
//
// Satisfies the tpl.TemplateHandler interface by accepting
// a tpl.Config data and set's up the template so it's
// ready for execution.
func (t *TemplateManager) Prepare(c tpl.TemplateConfig) tpl.TemplateExecutor {
	return &Execute{
		t,
		c,
		make(map[string]*template.Template),
		sync.RWMutex{},
		DefaultFileHandler(),
		template.FuncMap{},
	}
}

// Execute
//
// Satisfies the tpl.TemplateExecutor interface by executing
// a template with a io.Writer, the name of the template
// and any data to be passed. As this function is not
// attached to any context, the generic function map
// is used.
func (e *Execute) Execute(w io.Writer, name string, data interface{}) error {
	e.funcMap = e.GenericFuncMap()
	return e.executeRender(w, name, data)
}

// ExecutePost
//
// Satisfies the tpl.TemplateExecutor interface by executing
// a template with a io.Writer, the name of the template
// the context and the domain.PostData. Data is not
// needed to be  passed as data is obtained from
// the variables package.
func (e *Execute) ExecutePost(w io.Writer, name string, ctx *gin.Context, post *domain.PostData) error {
	data := e.Data(ctx, post)
	e.funcMap = e.FuncMap(ctx, post)
	return e.executeRender(w, name, data)
}

// Exists
//
// Satisfies the tpl.TemplateExecutor interface by determining
// if a template file exists with the given name.
func (e *Execute) Exists(template string) bool {
	_, err := e.fileHandler(e.config, template)
	if err != nil {
		return false
	}
	return true
}

// Data
//
// Satisfies the tpl.TemplateDataGetter interface by returning
// data for the front end that relies on context and post
// data.
func (t *TemplateManager) Data(ctx *gin.Context, post *domain.PostData) interface{} {
	return variables.Data(t.deps, ctx, post)
}

// FuncMap
//
// Satisfies the tpl.TemplateFuncGetter interface by returning
// functions that relies on context and post data such as
// `Meta` and `Url`. Generic functions are also included.
func (t *TemplateManager) FuncMap(ctx *gin.Context, post *domain.PostData) template.FuncMap {
	td := &internal.TemplateDeps{
		Context: ctx,
		Post:    post,
	}
	funcs := t.getFuncs(t.getNamespaces(td))
	return funcs
}

// GenericFuncMap
//
// Satisfies the tpl.TemplateFuncGetter interface by returning
// functions that do not rely on any data that is processed
// at runtime. These functions are loaded on initialisation
// and stored in memory.
func (t *TemplateManager) GenericFuncMap() template.FuncMap {
	return t.getFuncs(t.getGenericNamespaces())
}

