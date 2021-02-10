// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/tpl"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	d = &deps.Deps{
		Paths: deps.Paths{
			Theme: "theme",
		},
		Theme: &domain.ThemeConfig{
			FileExtension: "cms",
			TemplateDir:   "template",
		},
	}
)

func (t *RecoverTestSuite) Test_GetError() {

	tt := map[string]struct {
		input interface{}
		want  *errors.Error
	}{
		"Non Pointer": {
			errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("error")},
			&errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("error")},
		},
		"Pointer": {
			&errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("error")},
			&errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("error")},
		},
		"Standard Error": {
			fmt.Errorf("error"),
			&errors.Error{Code: errors.TEMPLATE, Message: "error", Operation: "", Err: fmt.Errorf("error")},
		},
		"Nil Input": {
			nil,
			&errors.Error{Code: errors.TEMPLATE, Message: "Internal Verbis error, please report", Operation: "Internal", Err: nil},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := getError(test.input)
			t.Equal(test.want.Operation, got.Operation)
			t.Equal(test.want.Message, got.Message)
			t.Equal(test.want.Code, got.Code)
			if test.want.Err != nil {
				t.Equal(test.want.Err.Error(), got.Err.Error())
			}
		})
	}
}

func (t *RecoverTestSuite) Test_TplLineNumber() {

	tt := map[string]struct {
		input *errors.Error
		want  int
	}{
		"Found": {
			&errors.Error{Err: fmt.Errorf(`template: templates/home:4: function "wrong" not defined`)},
			4,
		},
		"Found Second": {
			&errors.Error{Err: fmt.Errorf(`template: templates/home:10: function "wrong" not defined`)},
			10,
		},
		"Not Found": {
			&errors.Error{Err: fmt.Errorf(`template: templates/home10: function "wrong" not defined`)},
			-1,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := tplLineNumber(test.input)
			t.Equal(test.want, got)
		})
	}
}

func (t *RecoverTestSuite) Test_TplFileContents() {

	tt := map[string]struct {
		input string
		want  string
	}{
		"Found": {
			t.apiPath + "/test/testdata/html/partial.cms",
			"<h1>This is a partial file.</h1>",
		},
		"Not Found": {
			"wrong path",
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := tplFileContents(test.input)
			t.Equal(test.want, got)
		})
	}
}

func TestRecover_Resolver(t *testing.T) {

	tt := map[string]struct {
		input  bool
		code   int
		path   string
		mock   func(mh *mocks.TemplateHandler, m *mocks.TemplateExecutor, path string)
		custom bool
	}{
		"Default Code": {
			true,
			0,
			"error-500",
			func(mh *mocks.TemplateHandler, m *mocks.TemplateExecutor, path string) {
				mh.On("Prepare", tpl.Config{Root: "theme/template", Extension: "cms"}).Return(m)
				m.On("Exists", path).Return(true)
			},
			true,
		},
		"Custom error-404.cms": {
			true,
			404,
			"error-404",
			func(mh *mocks.TemplateHandler, m *mocks.TemplateExecutor, path string) {
				mh.On("Prepare", tpl.Config{Root: "theme/template", Extension: "cms"}).Return(m)
				m.On("Exists", path).Return(true)
			},
			true,
		},
		"Custom error-500.cms": {
			true,
			500,
			"error-500",
			func(mh *mocks.TemplateHandler, m *mocks.TemplateExecutor, path string) {
				mh.On("Prepare", tpl.Config{Root: "theme/template", Extension: "cms"}).Return(m)
				m.On("Exists", path).Return(true)
			},
			true,
		},
		"Custom error.cms": {
			true,
			500,
			"error",
			func(mh *mocks.TemplateHandler, m *mocks.TemplateExecutor, path string) {
				mh.On("Prepare", tpl.Config{Root: "theme/template", Extension: "cms"}).Return(m)
				m.On("Exists", "error-500").Return(false).Once()
				m.On("Exists", path).Return(true).Once()
			},
			true,
		},
		"Verbis Error Exec": {
			true,
			500,
			"templates/error",
			func(mh *mocks.TemplateHandler, m *mocks.TemplateExecutor, path string) {
				fe := &mocks.TemplateExecutor{}
				fe.On("Exists", "error-500").Return(false).Once()
				fe.On("Exists", "error").Return(false).Once()
				mh.On("Prepare", tpl.Config{Root: "theme/template", Extension: "cms"}).Return(fe).Once()
				mh.On("Prepare", tpl.Config{Root: d.Paths.Web, Extension: VerbisErrorExtension, Master: VerbisErrorLayout}).Return(m).Once()
			},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			handlerMock := &mocks.TemplateHandler{}
			templateMock := &mocks.TemplateExecutor{}
			test.mock(handlerMock, templateMock, test.path)

			d.SetTmpl(handlerMock)

			r := Recover{deps: d, config: Config{Code: test.code}}
			path, exec, custom := r.resolveErrorPage(test.custom)
			assert.Equal(t, test.custom, custom)
			assert.Equal(t, templateMock, exec)
			assert.Equal(t, test.path, path)
		})
	}
}
