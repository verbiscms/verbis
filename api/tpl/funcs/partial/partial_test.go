// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package partial

import (
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	mocks "github.com/verbiscms/verbis/api/mocks/tpl"
	"github.com/verbiscms/verbis/api/tpl/funcs/dict"
	"html/template"
	"os"
	"path/filepath"
	"testing"
)

func Setup(t *testing.T) *mocks.TemplateExecutor {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	m := &mocks.TemplateExecutor{}
	mc := &mocks.TemplateConfig{}

	m.On("Config").Return(mc)
	mc.On("GetRoot").Return(filepath.Join(wd, "testdata"))

	return m
}

func TestNamespace_Partial_ExecuteError(t *testing.T) {
	m := &mocks.TemplateExecutor{}
	mc := &mocks.TemplateConfig{}

	wd, err := os.Getwd()
	assert.NoError(t, err)

	m.On("Config").Return(mc)
	mc.On("GetRoot").Return(wd)

	orig := newTpl
	defer func() { newTpl = orig }()

	newTpl = func(name string) *template.Template {
		return template.New("name")
	}

	_, err = Partial(nil, m)("testdata/baddata.cms", nil)
	if err == nil {
		t.Fatal("expecting error")
		return
	}

	assert.Contains(t, err.Error(), "incomplete or empty template")
}

func TestNamespace_Partial(t *testing.T) {
	tt := map[string]struct {
		name     string
		data     interface{}
		multiple bool
		want     interface{}
	}{
		"Success": {
			`partial.cms`,
			nil,
			false,
			template.HTML(`<h1>This is a partial file.</h1>`),
		},
		"Wrong Path": {
			`wrongpath.cms`,
			nil,
			false,
			"no such file or directory",
		},
		"Bad Data": {
			`html/baddata.cms`,
			nil,
			false,
			template.HTML(""),
		},
		"File Mime": {
			`images/gopher.png`,
			nil,
			false,
			template.HTML(""),
		},
		"Dict": {
			`dict.cms`,
			map[string]interface{}{"Text": "cms"},
			false,
			template.HTML("cms"),
		},
		"Single Input": {
			`data.cms`,
			"verbis",
			false,
			template.HTML("verbis"),
		},
		"Multiple Inputs": {
			`data.cms`,
			[]interface{}{"hello", "verbis"},
			true,
			template.HTML("[hello verbis]"),
		},
		"Multiple Inputs 2": {
			`data.cms`,
			[]interface{}{"hello", "verbis", 1, 2, 3},
			true,
			template.HTML("[hello verbis 1 2 3]"),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			dic := dict.New(&deps.Deps{})

			p := Partial(template.FuncMap{
				"dict": dic.Dict,
			}, Setup(t))

			var got template.HTML
			var err error

			if test.multiple {
				slice, ok := test.data.([]interface{})
				assert.True(t, ok)
				got, err = p(test.name, slice...)
			} else {
				got, err = p(test.name, test.data)
			}

			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
