package tpl

import (
	"html/template"
)

type noStringer struct{}

func (t *TplTestSuite) Test_SafeHTML() {

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
			want:  "unable to cast tpl.noStringer{} of type tpl.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			tpl := `{{ safeHTML .Safe }}`
			t.RunTWithData(tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}

func (t *TplTestSuite) Test_SafeHTMLAttr() {

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
			want:  "unable to cast tpl.noStringer{} of type tpl.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			tpl := `{{ safeHTMLAttr .Safe }}`
			t.RunTWithData(tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}

func (t *TplTestSuite) Test_SafeCSS() {

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
			want:  "unable to cast tpl.noStringer{} of type tpl.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			tpl := `{{ safeCSS .Safe }}`
			t.RunTWithData(tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}

func (t *TplTestSuite) Test_SafeJS() {

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
			want:  "unable to cast tpl.noStringer{} of type tpl.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			tpl := `{{ safeJS .Safe }}`
			t.RunTWithData(tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}

func (t *TplTestSuite) Test_SafeJSStr() {

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
			want:  "unable to cast tpl.noStringer{} of type tpl.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			tpl := `{{ safeJSStr .Safe }}`
			t.RunTWithData(tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}

func (t *TplTestSuite) Test_SafeURL() {

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
			want:  "unable to cast tpl.noStringer{} of type tpl.noStringer to string",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			tpl := `{{ safeURL .Safe }}`
			t.RunTWithData(tpl, test.want, map[string]interface{}{"Safe": test.input})
		})
	}
}
