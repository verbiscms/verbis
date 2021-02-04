package internal

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

func TestAddFuncsNamespace(t *testing.T) {
	fns := []func(d *deps.Deps) *FuncsNamespace{
		func(d *deps.Deps) *FuncsNamespace {
			return &FuncsNamespace{
				Name: "test",
			}
		},
	}
	AddFuncsNamespace(fns[0])
	fnc1 := runtime.FuncForPC(reflect.ValueOf(fns[0]).Pointer()).Name()
	fnc2 := runtime.FuncForPC(reflect.ValueOf(GenericNamespaceRegistry[0]).Pointer()).Name()
	assert.Equal(t, fnc1, fnc2)
	GenericNamespaceRegistry = nil
}

func TestFuncsNamespace_AddMethodMapping(t *testing.T) {

	fns := FuncsNamespace{Name: "test"}

	tt := map[string]struct {
		name     string
		aliases  []string
		examples [][2]string
		want     interface{}
		panic    bool
	}{
		"Success": {
			"test",
			nil,
			nil,
			FuncMethodMapping{Name: "test"},
			false,
		},
		"Empty Example": {
			"test",
			nil,
			[][2]string{{""}, {""}},
			FuncMethodMapping{Name: "test"},
			true,
		},
		"Empty Alias": {
			"test",
			[]string{""},
			nil,
			FuncMethodMapping{Name: "test"},
			true,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			if test.panic {
				assert.Panics(t, func() {
					fns.AddMethodMapping(nil, test.name, test.aliases, test.examples)
				})
				return
			}

			fns.AddMethodMapping(nil, test.name, test.aliases, test.examples)
			got := fns.MethodMappings[test.name]
			assert.Equal(t, test.want, got)
		})
	}
}
