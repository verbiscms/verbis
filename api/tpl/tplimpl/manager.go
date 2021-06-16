// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tplimpl

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/tpl"
	"html/template"
	"io"
	"sync"
)

const (
	// DelimitersLeft left for the template.
	DelimitersLeft = "{{"
	// DelimitersRight right for the template.
	DelimitersRight = "}}"
)

// TemplateManager defines the service for executing and
// parsing Verbis templates. It's responsible for
// obtaining a template.FuncMap and Data to be
// used within the template.
type TemplateManager struct {
	deps *deps.Deps
}

// New creates a new TemplateManager.
func New(d *deps.Deps) *TemplateManager {
	return &TemplateManager{
		deps: d,
	}
}

// Execute defines the data for rendering a template
// contains.
type Execute struct {
	*TemplateManager
	config      tpl.TemplateConfig
	tplMap      map[string]*template.Template
	tplMutex    sync.RWMutex
	fileHandler fileHandler
	funcMap     template.FuncMap
}

// ExecuteTpl
//
// Is a helper function for executing standard templates
// that are embedded.
func (t *TemplateManager) ExecuteTpl(w io.Writer, text string, data interface{}) error {
	const op = "TemplateEngine.ExecuteTpl"

	tmpl, err := template.New("").Funcs(t.GenericFuncMap()).Parse(text)
	if err != nil {
		return &errors.Error{Code: errors.TEMPLATE, Message: "Error parsing template", Operation: op, Err: err}
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return &errors.Error{Code: errors.TEMPLATE, Message: "Error executing template", Operation: op, Err: err}
	}

	return nil
}
