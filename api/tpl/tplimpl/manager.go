// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tplimpl

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl"
	"html/template"
	"sync"
)

const (
	// Delimiter left for the template
	DelimitersLeft = "{{"
	// Delimiter right for the template
	DelimitersRight = "}}"
)

// TemplateManager defines the service for executing and
// parsing Verbis templates. It's responsible for
// obtaining a template.FuncMap and Data to be
// used within the template.
type TemplateManager struct {
	deps *deps.Deps
}

// Creates a new TemplateManager
func New(d *deps.Deps) *TemplateManager {
	return &TemplateManager{
		deps: d,
	}
}

// Execute defines the data for rendering a template
// contains
type Execute struct {
	*TemplateManager
	config      tpl.TemplateConfig
	tplMap      map[string]*template.Template
	tplMutex    sync.RWMutex
	fileHandler fileHandler
	funcMap     template.FuncMap
}
