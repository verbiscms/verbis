// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package partial

import (
	"github.com/ainsleyclark/verbis/api/deps"
	mocks "github.com/ainsleyclark/verbis/api/mocks/tpl"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/dict"
	"github.com/stretchr/testify/assert"
	"html/template"
	"os"
	"path/filepath"
	"testing"
)

const (
	TestPath = "/test/testdata"
)

func Setup(t *testing.T) *mocks.TemplateExecutor {
	wd, err := os.Getwd()
	assert.NoError(t, err)
	apiPath := filepath.Join(filepath.Dir(wd), "../..")

	m := &mocks.TemplateExecutor{}
	mc := &mocks.TemplateConfig{}

	m.On("Config").Return(mc)
	mc.On("GetRoot").Return(apiPath + TestPath)

	return m
}

func TestNamespace_Partial(t *testing.T) {

	tt := map[string]struct {
		name     string
		data     interface{}
		multiple bool
		want     interface{}
	}{
		"Success": {
			`html/partial.cms`,
			nil,
			false,
			template.HTML(`<h1>This is a partial file.</h1>`),
		},
		"Wrong Path": {
			`html/wrongpath.cms`,
			nil,
			false,
			"Templates.Partial: no file exists with the path: html/wrongpath.cms",
		},
		"Bad Data": {
			`html/partial-baddata.cms`,
			nil,
			false,
			template.HTML(""),
		},
		//"Error Executing": {
		//	`html/partial.cms`,
		//	[]interface{}{make(chan int),make(chan int)},
		//	false,
		//	template.HTML("unable to execute partial file"),
		//},
		"File Type": {
			`images/gopher.png`,
			nil,
			false,
			template.HTML(""),
		},
		"Dict": {
			`html/partial-dict.cms`,
			map[string]interface{}{"Text": "cms"},
			false,
			template.HTML("cms"),
		},
		"Single Input": {
			`html/partial-data.cms`,
			"verbis",
			false,
			template.HTML("verbis"),
		},
		"Multiple Inputs": {
			`html/partial-data.cms`,
			[]interface{}{"hello", "verbis"},
			true,
			template.HTML("[hello verbis]"),
		},
		"Multiple Inputs 2": {
			`html/partial-data.cms`,
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
