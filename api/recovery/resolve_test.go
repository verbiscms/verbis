// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	mocks "github.com/ainsleyclark/verbis/api/mocks/tpl"
	"github.com/ainsleyclark/verbis/api/tpl"
)

func (t *RecoverTestSuite) TestRecover_Resolver() {

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
				mh.On("Prepare", tpl.Config{Root: t.deps.Paths.Web, Extension: VerbisErrorExtension, Master: VerbisErrorLayout}).Return(m).Once()
			},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			handlerMock := &mocks.TemplateHandler{}
			templateMock := &mocks.TemplateExecutor{}
			test.mock(handlerMock, templateMock, test.path)

			t.deps.SetTmpl(handlerMock)

			r := Recover{deps: t.deps, config: Config{Code: test.code}}
			path, exec, custom := r.resolveErrorPage(test.custom)

			t.Equal(test.custom, custom)
			t.Equal(templateMock, exec)
			t.Equal(test.path, path)
		})
	}
}
