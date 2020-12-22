package templates

import (
	"html/template"
	"testing"
)

type noStringer struct{}

func Test_SafeHTML(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			input: "&lt;h1&gt;Hello&lt;/h1&gt;",
			want:  "<h1>Hello</h1>",
		},
		"int": {
			input: 64,
			want:  "64",
		},
		"Error": {
			input: noStringer{},
			want:  "unable to cast templates.noStringer{} of type templates.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			tpl := `{{ safeHTML .Safe }}`
			runtv(t, f, tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}

func Test_SafeHTMLAttr(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			input: `dir="ltd"`,
			want:  template.HTMLAttr(`dir="ltd"`),
		},
		"Error": {
			input: noStringer{},
			want:  "unable to cast templates.noStringer{} of type templates.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			tpl := `{{ safeHTMLAttr .Safe }}`
			runtv(t, f, tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}

func Test_SafeCSS(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			input: `a[href =~ "//verbiscms.com"]#foo`,
			want:  template.CSS(`a[href =~ "//verbiscms.com"]#foo`),
		},
		"Error": {
			input: noStringer{},
			want:  "unable to cast templates.noStringer{} of type templates.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			tpl := `{{ safeCSS .Safe }}`
			runtv(t, f, tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}

func Test_SafeJS(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			input: `alert("Hello, Verbis!");`,
			want:  template.JS(`alert("Hello, Verbis!");`),
		},
		"Error": {
			input: noStringer{},
			want:  "unable to cast templates.noStringer{} of type templates.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			tpl := `{{ safeJS .Safe }}`
			runtv(t, f, tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}

func Test_SafeJSStr(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			input: `Verbis CMS \x21`,
			want:  template.JSStr(`Verbis CMS \x21`),
		},
		"Error": {
			input: noStringer{},
			want:  "unable to cast templates.noStringer{} of type templates.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			tpl := `{{ safeJSStr .Safe }}`
			runtv(t, f, tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}

func Test_SafeURL(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Valid": {
			input: `verbis=H%71&title=(CMS)`,
			want:  template.URL(`verbis=H%71&title=(CMS)`),
		},
		"Error": {
			input: noStringer{},
			want:  "unable to cast templates.noStringer{} of type templates.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			tpl := `{{ safeURL .Safe }}`
			runtv(t, f, tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}
