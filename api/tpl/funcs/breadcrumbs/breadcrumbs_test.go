// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package breadcrumbs

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/logger"
	mocks "github.com/verbiscms/verbis/api/mocks/tpl"
	"github.com/verbiscms/verbis/api/verbis"
	"html/template"
	"io"
	"io/ioutil"
	"testing"
)

var (
	crumbs = verbis.Breadcrumbs{
		Enabled: true,
		Title:   "Items",
	}
)

func TestNamespace_Get(t *testing.T) {
	ns := Namespace{
		deps:   nil,
		crumbs: crumbs,
	}
	got := ns.Get()
	assert.Equal(t, crumbs, got)
}

func TestNamespace_HTML(t *testing.T) {
	tt := map[string]struct {
		mock func(th *mocks.TemplateHandler)
		fn   func(ns *Namespace) interface{}
		want interface{}
	}{
		"Success": {
			func(th *mocks.TemplateHandler) {
				th.On("ExecuteTpl", &bytes.Buffer{}, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					arg := args.Get(0).(io.Writer)
					_, err := arg.Write([]byte("test"))
					assert.NoError(t, err)
				}).Return(nil)
			},
			func(ns *Namespace) interface{} {
				return ns.HTML()
			},
			template.HTML("test"),
		},
		"Execute Error": {
			func(th *mocks.TemplateHandler) {
				th.On("ExecuteTpl", &bytes.Buffer{}, mock.Anything, mock.Anything).Return(fmt.Errorf("error"))
			},
			func(ns *Namespace) interface{} {
				return ns.HTML()
			},
			template.HTML(""),
		},
		"Wrong File": {
			nil,
			func(ns *Namespace) interface{} {
				orig := TemplateName
				defer func() {
					TemplateName = orig
				}()
				TemplateName = "wrong"

				return ns.HTML()
			},
			template.HTML(""),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			logger.SetOutput(ioutil.Discard)

			tplMock := &mocks.TemplateHandler{}
			ns := &Namespace{
				deps:   &deps.Deps{},
				crumbs: crumbs,
			}
			ns.deps.SetTmpl(tplMock)

			if test.mock != nil {
				test.mock(tplMock)
			}

			got := test.fn(ns)

			assert.Equal(t, test.want, got)
		})
	}
}
