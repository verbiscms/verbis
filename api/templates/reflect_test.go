package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"testing"
)

type testReflect struct {
	Name, Value string
}

var tr = &testReflect{"hello", "verbis"}

func Test_KindIs(t *testing.T) {

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
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), test.tpl, test.want, test.input)
		})
	}
}

func Test_KindOf(t *testing.T) {

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
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), `{{ kindOf . }}`, test.want, test.input)
		})
	}
}

func Test_TypeOf(t *testing.T) {

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
			want:  "*templates.testReflect",
		},
		"Post": {
			input: &domain.Post{},
			want:  "*domain.Post",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), `{{ typeOf . }}`, test.want, test.input)
		})
	}
}

func Test_TypeIs(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
		tpl   string
	}{
		"True": {
			input: tr,
			want:  true,
			tpl:   `{{ typeIs "*templates.testReflect" . }}`,
		},
		"False": {
			input: tr,
			want:  false,
			tpl:   `{{ typeIs "wrongval" . }}`,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), test.tpl, test.want, test.input)
		})
	}
}

func Test_TypeIsLike(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
		tpl   string
	}{
		"True": {
			input: *tr,
			want:  true,
			tpl:   `{{ typeIsLike "templates.testReflect" . }}`,
		},
		"True Pointer": {
			input: tr,
			want:  true,
			tpl:   `{{ typeIsLike "*templates.testReflect" . }}`,
		},
		"False": {
			input: tr,
			want:  false,
			tpl:   `{{ typeIsLike "wrongval" . }}`,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			runtv(t, newTestSuite(), test.tpl, test.want, test.input)
		})
	}
}
