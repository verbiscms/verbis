package tpl

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

type testReflect struct {
	Name, Value string
}

var tr = &testReflect{"hello", "verbis"}

func (t *TplTestSuite) Test_KindIs() {

	tt := map[string]struct {
		input interface{}
		want  interface{}
		tpl   string
	}{
		"True": {
			input: tr,
			want:  true,
			tpl:   `{{ kindIs "ptr" . }}`,
		},
		"False": {
			input: tr,
			want:  false,
			tpl:   `{{ kindIs "hello" . }}`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RunTWithData(test.tpl, test.want, test.input)
		})
	}
}

func (t *TplTestSuite) Test_KindOf() {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Slice": {
			input: []string{"hello"},
			want:  "slice",
		},
		"Int": {
			input: 123,
			want:  "int",
		},
		"String": {
			input: "hello",
			want:  "string",
		},
		"Struct": {
			input: tr,
			want:  "ptr",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RunTWithData(`{{ kindOf . }}`, test.want, test.input)
		})
	}
}

func (t *TplTestSuite) Test_TypeOf() {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Slice": {
			input: []string{"hello"},
			want:  "[]string",
		},
		"Int": {
			input: 123,
			want:  "int",
		},
		"String": {
			input: "hello",
			want:  "string",
		},
		"Struct": {
			input: tr,
			want:  "*tpl.testReflect",
		},
		"Post": {
			input: &domain.Post{},
			want:  "*domain.Post",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RunTWithData(`{{ typeOf . }}`, test.want, test.input)
		})
	}
}

func (t *TplTestSuite) Test_TypeIs() {

	tt := map[string]struct {
		input interface{}
		want  interface{}
		tpl   string
	}{
		"True": {
			input: tr,
			want:  true,
			tpl:   `{{ typeIs "*tpl.testReflect" . }}`,
		},
		"False": {
			input: tr,
			want:  false,
			tpl:   `{{ typeIs "wrongval" . }}`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RunTWithData(test.tpl, test.want, test.input)
		})
	}
}

func (t *TplTestSuite) Test_TypeIsLike() {

	tt := map[string]struct {
		input interface{}
		want  interface{}
		tpl   string
	}{
		"True": {
			input: *tr,
			want:  true,
			tpl:   `{{ typeIsLike "tpl.testReflect" . }}`,
		},
		"True Pointer": {
			input: tr,
			want:  true,
			tpl:   `{{ typeIsLike "*tpl.testReflect" . }}`,
		},
		"False": {
			input: tr,
			want:  false,
			tpl:   `{{ typeIsLike "wrongval" . }}`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RunTWithData(test.tpl, test.want, test.input)
		})
	}
}
