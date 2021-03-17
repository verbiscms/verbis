// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
func (e *Execute) Execute(w io.Writer, name string, data interface{}) (string, error) {
	e.funcMap = e.GenericFuncMap()
	return e.executeRender(w, name, data)
}

// ExecutePost
//
// Satisfies the tpl.TemplateExecutor interface by executing
// a template with a io.Writer, the name of the template
// the context and the domain.PostDatum. Data is not
// needed to be  passed as data is obtained from
// the variables package.
func (e *Execute) ExecutePost(w io.Writer, name string, ctx *gin.Context, post *domain.PostDatum) (string, error) {
	data := e.Data(ctx, post)
	e.funcMap = e.FuncMap(ctx, post, e.config)
	return e.executeRender(w, name, data)
}

// Exists
//
// Satisfies the tpl.TemplateExecutor interface by determining
// if a template file exists with the given name.
func (e *Execute) Exists(name string) bool {
	_, err := e.fileHandler(e.config, name)
	return err == nil
}

// Config
//
// Satisfies the tpl.TemplateExecutor interface by returning
// the Execute configuration to obtain the root, extension
// and master layout.
func (e *Execute) Config() tpl.TemplateConfig {
	return e.config
}

// Executor
//
// Satisfies the tpl.TemplateExecutor interface by returning
// itself for use with recovery
func (e *Execute) Executor() tpl.TemplateExecutor {
	return e
}

// Data
//
// Satisfies the tpl.TemplateDataGetter interface by returning
// data for the front end that relies on context and post
// data.
func (t *TemplateManager) Data(ctx *gin.Context, post *domain.PostDatum) interface{} {
	return variables.Data(t.deps, ctx, post)
}

// FuncMap
//
// Satisfies the tpl.TemplateFuncGetter interface by returning
// functions that relies on context and post data such as
// `Meta` and `url`. Generic functions are also included.
func (t *TemplateManager) FuncMap(ctx *gin.Context, post *domain.PostDatum, cfg tpl.TemplateConfig) template.FuncMap {
	td := &internal.TemplateDeps{
		Context: ctx,
		Post:    post,
		Cfg:     cfg,
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
