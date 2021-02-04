package reflect

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testReflect struct {
	Name, Value string
}

var (
	ns = New(&deps.Deps{})
	tr = &testReflect{"hello", "verbis"}
)

func TestNamespace_KindIs(t *testing.T) {

	tt := map[string]struct {
		target string
		src    interface{}
		want   bool
	}{
		"True": {
			"ptr",
			tr,
			true,
		},
		"False": {
			"hello",
			tr,
			false,
		},
		"Nil": {
			"",
			nil,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.KindIs(test.target, test.src)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_KindOf(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Slice": {
			[]string{"hello"},
			"slice",
		},
		"Int": {
			123,
			"int",
		},
		"String": {
			"hello",
			"string",
		},
		"Struct": {
			tr,
			"ptr",
		},
		"Nil": {
			nil,
			"invalid",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.KindOf(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_TypeOf(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"Slice": {
			[]string{"hello"},
			"[]string",
		},
		"Int": {
			123,
			"int",
		},
		"String": {
			"hello",
			"string",
		},
		"Struct": {
			tr,
			"*reflect.testReflect",
		},
		"Nil": {
			nil,
			"<nil>",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.TypeOf(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_TypeIs(t *testing.T) {

	tt := map[string]struct {
		target string
		src    interface{}
		want   bool
	}{
		"True": {
			"*reflect.testReflect",
			tr,
			true,
		},
		"False": {
			"wrongval",
			tr,
			false,
		},
		"Nil": {
			"",
			nil,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.TypeIs(test.target, test.src)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_TypeIsLike(t *testing.T) {

	tt := map[string]struct {
		target string
		src    interface{}
		want   bool
	}{
		"True": {
			"reflect.testReflect",
			*tr,
			true,
		},
		"True Pointer": {
			"*reflect.testReflect",
			tr,
			true,
		},
		"False": {
			"wrongval",
			tr,
			false,
		},
		"Nil": {
			"",
			nil,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.TypeIsLike(test.target, test.src)
			assert.Equal(t, test.want, got)
		})
	}
}
