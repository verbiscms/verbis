// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/gin-contrib/location"
	"io/ioutil"
	"net/http"
)

type (
	// Data represents the main struct for sending back data to the
	// template for recovery.
	Data struct {
		Error   Error
		Request Request
		Post    *domain.PostData
		Stack   Stack
		Debug   bool
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
		Headers    map[string][]string
		Query      map[string][]string
		Body       string
		Cookies    []*http.Cookie
		IP         string
		DataLength int
		UserAgent  string
		Referer    string
	}
)

// getData
//
// Retrieves all the necessary template data to show
// the recovery page. If a template executor has
// been set, the template file stack will be
// retrieved,
func (r *Recover) getData() *Data {
	return &Data{
		Error:   r.getErrorData(),
		Request: r.getRequestData(),
		Post:    r.config.Post,
		Stack:   r.getStackData(),
		Debug:   environment.IsDebug(),
	}
}

// getStackData
//
// Check if the template exec has been set, if it has
// retrieve the file stack for the template. and
// prepend it to the stack
func (r *Recover) getStackData() Stack {
	stack := GetStack(StackDepth, StackSkip)
	if r.config.TplExec != nil && r.config.TplFile != "" {
		path := r.config.TplExec.Config().GetRoot() + "/" + r.config.TplFile + r.config.TplExec.Config().GetExtension()
		stack.Prepend(&File{
			File:     path,
			Line:     tplLineNumber(r.config.Error),
			Name:     r.config.TplFile,
			Contents: tplFileContents(path),
		})
	}
	return stack
}

// getErrorData
//
//
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
//
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
		Headers:    ctx.Request.Header,
		Query:      ctx.Request.URL.Query(),
		Body:       string(body),
		Cookies:    ctx.Request.Cookies(),
		IP:         ctx.ClientIP(),
		DataLength: ctx.Writer.Size(),
		UserAgent:  ctx.Request.UserAgent(),
		Referer:    ctx.Request.Referer(),
	}
}
