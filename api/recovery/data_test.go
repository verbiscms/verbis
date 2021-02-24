// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/tpl"
	"github.com/ainsleyclark/verbis/api/recovery/trace"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var (
	cookie = http.Cookie{Name: "test"}
)

func (t *RecoverTestSuite) TestRecover_GetData() {

	err := os.Setenv("APP_DEBUG", "false")
	t.NoError(err)

	want := &Data{
		Error: Error{
			Code:      errors.TEMPLATE,
			Message:   "message",
			Operation: "operation",
			Err:       "operation: error",
		},
		Request: Request{
			Url:        "http://localhost:8080/test",
			Method:     "GET",
			Headers:    map[string][]string{"Header": {"test"}},
			Query:      map[string][]string{"page": {"test"}},
			Body:       "test",
			Cookies:    []*http.Cookie{},
			IP:         "",
			DataLength: 0,
			Referer:    "",
		},
		Post:  nil,
		Debug: true,
	}

	r := Recover{
		deps:   nil,
		err:    &errors.Error{Code: errors.TEMPLATE, Message: "message", Operation: "operation", Err: fmt.Errorf("error")},
		config: Config{},
		tracer: trace.New(),
	}

	t.RequestSetup(bytes.NewBuffer([]byte("test")), nil, func(ctx *gin.Context) {
		r.config.Context = ctx
	})

	got := r.getData()
	t.Equal(want.Error, got.Error)
	t.Equal(want.Request, got.Request)
	t.Equal(want.Post, got.Post)
	t.Equal(want.Debug, got.Debug)
}

func (t *RecoverTestSuite) TestRecover_GetStackData() {

	m := &mocks.TemplateExecutor{}
	mc := &mocks.TemplateConfig{}
	mc.On("GetRoot").Return("test")
	mc.On("GetExtension").Return("cms")
	m.On("Config").Return(mc)

	tt := map[string]struct {
		input Config
		want  func(s trace.Stack) trace.Stack
	}{
		"Nil Exec": {
			Config{TplExec: nil},
			func(s trace.Stack) trace.Stack {
				return s
			},
		},
		"Nil TplFile": {
			Config{TplFile: ""},
			func(s trace.Stack) trace.Stack {
				return s
			},
		},
		"With Template": {
			Config{
				TplFile: "test",
				TplExec: m,
			},
			func(s trace.Stack) trace.Stack {
				s.Append(&trace.File{File: "tt", Line: 0, Contents: "tt"})
				return s
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			r := Recover{config: test.input, tracer: trace.New()}
			stack := trace.New().Trace(StackDepth, 1)
			t.Equal(len(test.want(stack)), len(r.getStackData()))
		})
	}
}

func (t *RecoverTestSuite) TestRecover_GetErrorData() {

	tt := map[string]struct {
		input *errors.Error
		want  Error
	}{
		"Simple": {
			&errors.Error{Code: errors.TEMPLATE, Message: "message", Operation: "operation", Err: fmt.Errorf("error")},
			Error{Code: errors.TEMPLATE, Message: "message", Operation: "operation", Err: "operation: error"},
		},
		"Nil Error": {
			&errors.Error{Code: errors.TEMPLATE, Message: "message", Operation: "operation", Err: nil},
			Error{Code: errors.TEMPLATE, Message: "message", Operation: "operation", Err: "operation: <template> message"},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			r := Recover{err: test.input}
			t.Equal(test.want, r.getErrorData())
		})
	}
}

func (t *RecoverTestSuite) TestRecover_GetRequestData() {
	want := Request{
		Url:        "http://localhost:8080/test",
		Method:     "GET",
		Headers:    map[string][]string{"Header": {"test"}, "Cookie": {"test="}},
		Query:      map[string][]string{"page": {"test"}},
		Body:       "test",
		Cookies:    []*http.Cookie{&cookie},
		IP:         "",
		DataLength: -1,
		Referer:    "",
	}

	t.RequestSetup(bytes.NewBuffer([]byte("test")), &cookie, func(ctx *gin.Context) {
		r := Recover{config: Config{Context: ctx}}
		got := r.getRequestData()
		t.Equal(want, got)
	})
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("test error")
}

func (t *RecoverTestSuite) TestRecover_GetRequestData_NilBody() {
	want := Request{
		Url:        "http://localhost:8080/test",
		Method:     "GET",
		Headers:    map[string][]string{"Header": {"test"}, "Cookie": {"test="}},
		Query:      map[string][]string{"page": {"test"}},
		Body:       "",
		Cookies:    []*http.Cookie{&cookie},
		IP:         "",
		DataLength: -1,
		Referer:    "",
	}

	t.RequestSetup(errReader(0), &cookie, func(ctx *gin.Context) {
		r := Recover{config: Config{Context: ctx}}
		got := r.getRequestData()
		t.Equal(want, got)
	})
}
