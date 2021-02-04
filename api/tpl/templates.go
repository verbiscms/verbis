// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tpl

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
)

// TemplateHandler is the main template renderer for Verbis
// It's responsible for preparing and executing templates
// and obtaining information such as function maps and
// post specific data.
type TemplateHandler interface {
	TemplateFuncGetter
	TemplateDataGetter
	Prepare(c TemplateConfig) TemplateExecutor
}

// TemplateExecute represents the functions for executing
// template.
type TemplateExecutor interface {
	Exists(template string) bool
	Execute(w io.Writer, name string, data interface{}) error
	ExecutePost(w io.Writer, name string, ctx *gin.Context, post *domain.PostData) error
	Config() TemplateConfig
}

// TemplateFuncGetter represents the functions for obtaining
// template.FuncMap's for use in Verbis templates.
type TemplateFuncGetter interface {
	FuncMap(ctx *gin.Context, post *domain.PostData, cfg TemplateConfig) template.FuncMap
	GenericFuncMap() template.FuncMap
}

// TemplateDataGetter represents the the Data function
// for obtaining post relevant data to send back to
// the template.
type TemplateDataGetter interface {
	Data(ctx *gin.Context, post *domain.PostData) interface{}
}

// TemplateConfig represents the functions for obtaining
// the executor configuration including "root",
// "master" and "extension".
type TemplateConfig interface {
	GetRoot() string
	GetExtension() string
	GetMaster() string
}

// Config represents the options for passing
type Config struct {
	Root      string
	Extension string
	Master    string
}

// GetRoot
//
// Returns the view root
func (c Config) GetRoot() string {
	return c.Root
}

// GetExtension
//
// Returns the template extension
func (c Config) GetExtension() string {
	return c.Extension
}

// GetMaster
//
// Returns the template master layout
func (c Config) GetMaster() string {
	return c.Master
}
