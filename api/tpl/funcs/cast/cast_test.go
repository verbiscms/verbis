package cast

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ns = New(&deps.Deps{})
)

func TestNamespace_ToSlice(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"String": {
			"a",
			[]interface{}{"a"},
		},
		"String Slice": {
			[]interface{}{"a", "b"},
			[]interface{}{"a", "b"},
		},
		"Int": {
			1,
			[]interface{}{1},
		},
		"Int Slice": {
			[]interface{}{1, 2},
			[]interface{}{1, 2},
		},
		"Map": {
			[]map[string]interface{}{{"a": 1}, {"a": 2}},
			[]interface{}{map[string]interface{}{"a": 1}, map[string]interface{}{"a": 2}},
		},
		"Nil": {
			nil,
			[]interface{}(nil),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.ToSlice(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
