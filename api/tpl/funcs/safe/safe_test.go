// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package safe

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
)

var (
	ns = New(&deps.Deps{})
)

type noStringer struct{}

func TestNamespace_SafeHTML(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			"<h1>hello verbis &amp; world!</h1>",
			template.HTML("<h1>hello verbis &amp; world!</h1>"),
		},
		"int": {
			64,
			template.HTML("64"),
		},
		"Error": {
			noStringer{},
			"unable to cast safe.noStringer{}",
		},
		"Nil": {
			nil,
			template.HTML(""),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.HTML(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_SafeHTMLAttr(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			`dir="ltd"`,
			template.HTMLAttr(`dir="ltd"`),
		},
		"Error": {
			noStringer{},
			"unable to cast safe.noStringer{}",
		},
		"Nil": {
			nil,
			template.HTMLAttr(""),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.HTMLAttr(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_SafeCSS(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			`a[href =~ "//verbiscms.com"]#foo`,
			template.CSS(`a[href =~ "//verbiscms.com"]#foo`),
		},
		"Error": {
			noStringer{},
			"unable to cast safe.noStringer{}",
		},
		"Nil": {
			nil,
			template.CSS(""),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.CSS(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_SafeJS(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			`alert("Hello, Verbis!");`,
			template.JS(`alert("Hello, Verbis!");`),
		},
		"Error": {
			noStringer{},
			"unable to cast safe.noStringer{}",
		},
		"Nil": {
			nil,
			template.JS(""),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.JS(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_SafeJSStr(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			`Verbis CMS \x21`,
			template.JSStr(`Verbis CMS \x21`),
		},
		"Error": {
			noStringer{},
			"unable to cast safe.noStringer{}",
		},
		"Nil": {
			nil,
			template.JSStr(""),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.JSStr(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_SafeURL(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			`verbis=H%71&title=(CMS)`,
			template.URL(`verbis=H%71&title=(CMS)`),
		},
		"Error": {
			noStringer{},
			"unable to cast safe.noStringer{}",
		},
		"Nil": {
			nil,
			template.URL(""),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.Url(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
