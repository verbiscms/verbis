// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/recovery/trace"
	"github.com/gin-contrib/location"
	"io/ioutil"
	"net/http"
)

type (
	// Data represents the main struct for sending back data to the
	// template for recovery.
	Data struct {
		Error      Error
		StatusCode int
		Request    Request
		Post       *domain.PostDatum
		Stack      trace.Stack
		Context    Context
		Debug      bool
	}
	// Error represents a errors.Error in friendly form (strings) to
	// for the recovery template.
	Error struct {
		Code      string
		Message   string
		Operation string
		Err       string
	}
	// Request represents the data obtained from the context with
	// detailed information about the http request made.
	Request struct {
		Url        string
		Method     string
		IP         string
		Referer    string
		DataLength int
		Body       string
		Headers    map[string][]string
		Query      map[string][]string
		Cookies    []*http.Cookie
	}
	// Context represents general information about Verbis to
	// help with debugging.
	Context struct {
		Version string
		Site    domain.Site
		Options domain.Options
	}
)

const (
	// The amount of files in the stack to be retrieved.
	StackDepth = 200
	// How many files to move up in the runtime.Caller
	// before obtaining the stack.
	StackSkip = 2
)

// getData
//
// Retrieves all the necessary template data to show
// the recovery page. If a template executor has
// been set, the template file stack will be
// retrieved,
func (r *Recover) getData() *Data {
	return &Data{
		Error:      r.getErrorData(),
		StatusCode: r.config.Code,
		Request:    r.getRequestData(),
		Post:       r.config.Post,
		Stack:      r.getStackData(),
		Context:    r.getContextData(),
		// TEMPORARY
		Debug: true,
	}
}

// getStackData
//
// Check if the template exec has been set, if it has
// retrieve the file stack for the template. and
// prepend it to the stack.
func (r *Recover) getStackData() trace.Stack {
	stack := r.tracer.Trace(StackDepth, StackSkip)
	if r.config.TplExec != nil && r.config.TplFile != "" {
		root := r.config.TplExec.Config().GetRoot()
		ext := r.config.TplExec.Config().GetExtension()
		path := root + "/" + r.config.TplFile + ext

		stack.Prepend(&trace.File{
			File:     path,
			Line:     tplLineNumber(r.err),
			Function: r.config.TplFile,
			Contents: tplFileContents(path),
			Language: "handlebars",
		})
	}
	return stack
}

// getErrorData
//
// Returns error friendly Error data for the template.
func (r *Recover) getErrorData() Error {
	return Error{
		Code:      r.err.Code,
		Message:   r.err.Message,
		Operation: r.err.Operation,
		Err:       r.err.Error(),
	}
}

// getRequestData
//
// Returns error friendly Request data for the template.
func (r *Recover) getRequestData() Request {
	ctx := r.config.Context

	// Retrieve the request body contents.
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		body = nil
	}

	return Request{
		Url:        location.Get(ctx).String() + ctx.Request.URL.Path,
		Method:     ctx.Request.Method,
		IP:         ctx.ClientIP(),
		Referer:    ctx.Request.Referer(),
		DataLength: ctx.Writer.Size(),
		Body:       string(body),
		Headers:    ctx.Request.Header,
		Query:      ctx.Request.URL.Query(),
		Cookies:    ctx.Request.Cookies(),
	}
}

// getContextData
//
// Returns error friendly request Context for the template.
func (r *Recover) getContextData() Context {
	return Context{
		Version: api.App.Version,
		Site:    r.deps.Site,
		Options: *r.deps.Options,
	}
}
