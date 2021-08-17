// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/logger"
	tpl "github.com/verbiscms/verbis/api/mocks/tpl"
	mocks "github.com/verbiscms/verbis/api/mocks/verbis/nav"
	"github.com/verbiscms/verbis/api/verbis/nav"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	menu = nav.Menu{Options: nav.Options{Menu: "main-menu"}}
	args = nav.Args{"menu": "main-menu"}
)

func TestNamespace_Get(t *testing.T) {
	tt := map[string]struct {
		mock func(m *mocks.Getter)
		want interface{}
	}{
		"Success": {
			func(m *mocks.Getter) {
				m.On("Get", args).Return(menu, nil)
			},
			menu,
		},
		"Error": {
			func(m *mocks.Getter) {
				m.On("Get", args).Return(menu, fmt.Errorf("error"))
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			navMock := &mocks.Getter{}
			ns := &Namespace{
				deps: &deps.Deps{},
				nav:  navMock,
			}
			if test.mock != nil {
				test.mock(navMock)
			}

			got, err := ns.Get(args)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_HTML(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)
	testPath := filepath.Join(wd, "testdata")

	tt := map[string]struct {
		mock func(th *tpl.TemplateHandler, n *mocks.Getter)
		fn   func(ns *Namespace) (interface{}, error)
		want interface{}
	}{
		"Success": {
			func(th *tpl.TemplateHandler, m *mocks.Getter) {
				m.On("Get", args).Return(menu, nil)
				th.On("ExecuteTpl", &bytes.Buffer{}, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					arg := args.Get(0).(io.Writer)
					_, err := arg.Write([]byte("test"))
					assert.NoError(t, err)
				}).Return(nil)
			},
			func(ns *Namespace) (interface{}, error) {
				return ns.HTML(args)
			},
			template.HTML("test"),
		},
		"Partial": {
			func(th *tpl.TemplateHandler, m *mocks.Getter) {
				m.On("Get", nav.Args{"menu": "main-menu", "partial": "nav.cms"}).Return(nav.Menu{Options: nav.Options{Partial: "nav.cms"}}, nil)
				th.On("ExecuteTpl", &bytes.Buffer{}, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					arg := args.Get(0).(io.Writer)
					_, err := arg.Write([]byte(args.Get(1).(string)))
					assert.NoError(t, err)
				}).Return(nil)
			},
			func(ns *Namespace) (interface{}, error) {
				return ns.HTML(nav.Args{"menu": "main-menu", "partial": "nav.cms"})
			},
			template.HTML("<h1>\n  Hello\n</h1>"),
		},
		"menusDB Error": {
			func(th *tpl.TemplateHandler, m *mocks.Getter) {
				m.On("Get", args).Return(menu, fmt.Errorf("nav error"))
			},
			func(ns *Namespace) (interface{}, error) {
				return ns.HTML(args)
			},
			"nav error",
		},
		"Execute Error": {
			func(th *tpl.TemplateHandler, m *mocks.Getter) {
				m.On("Get", args).Return(menu, nil)
				th.On("ExecuteTpl", &bytes.Buffer{}, mock.Anything, mock.Anything).Return(fmt.Errorf("execute error"))
			},
			func(ns *Namespace) (interface{}, error) {
				return ns.HTML(args)
			},
			"execute error",
		},
		"Wrong File": {
			func(th *tpl.TemplateHandler, m *mocks.Getter) {
				m.On("Get", args).Return(menu, nil)
			},
			func(ns *Namespace) (interface{}, error) {
				orig := TemplateName
				defer func() {
					TemplateName = orig
				}()
				TemplateName = "wrong"

				return ns.HTML(args)
			},
			template.HTML(""),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			logger.SetOutput(ioutil.Discard)

			tplMock := &tpl.TemplateHandler{}
			navMock := &mocks.Getter{}
			ns := &Namespace{
				deps: &deps.Deps{},
				nav:  navMock,
				themeGetter: func() string {
					return testPath
				},
			}
			ns.deps.SetTmpl(tplMock)

			if test.mock != nil {
				test.mock(tplMock, navMock)
			}

			got, err := test.fn(ns)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.Equal(t, test.want, got)
		})
	}
}
