package templates

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RandInt(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want []int
	}{
		"Valid": {
			tmpl: `{{ randInt 1 100 }}`,
			want: []int{1, 100},
		},
		"Valid 2": {
			tmpl: `{{ randInt 1 5 }}`,
			want: []int{1, 5},
		},
		"Large": {
			tmpl: `{{ randInt 1 99999999 }}`,
			want: []int{1, 99999999},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			str, err := execute(newTestSuite(), test.tmpl, nil)
			assert.NoError(t, err)
			got := cast.ToInt(str)
			if test.want[0] > got || test.want[1] < got {
				t.Error(fmt.Errorf("got %v expecting in between %v", got, test.want))
			}
		})
	}
}

func Test_RandFloat(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want []float64
	}{
		"Valid": {
			tmpl: `{{ randFloat 1 100 }}`,
			want: []float64{1, 100},
		},
		"Valid2": {
			tmpl: `{{ randFloat 1.555555 10.44444 }}`,
			want: []float64{1.555555, 10.44444},
		},
		"Large": {
			tmpl: `{{ randInt 1 99999999.99999 }}`,
			want: []float64{1, 99999999.99999},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			str, err := execute(newTestSuite(), test.tmpl, nil)
			assert.NoError(t, err)
			got := cast.ToFloat64(str)
			if test.want[0] > got || test.want[1] < got {
				t.Error(fmt.Errorf("got %v expecting in between %v", got, test.want))
			}
		})
	}
}

func Test_RandAlpha(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		len int
	}{
		"Valid": {
			tmpl: `{{ randAlpha 5  }}`,
			len: 5,
		},
		"Valid 2": {
			tmpl: `{{ randAlpha 10  }}`,
			len: 10,
		},
		"Valid 3": {
			tmpl: `{{ randAlpha 100  }}`,
			len: 100,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := execute(newTestSuite(), test.tmpl, nil)
			assert.NoError(t, err)
			assert.Equal(t, test.len, len(got))
		})
	}
}

func Test_RandAlphaNum(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		len int
	}{
		"Valid": {
			tmpl: `{{ randAlphaNum 5  }}`,
			len: 5,
		},
		"Valid 2": {
			tmpl: `{{ randAlphaNum 10  }}`,
			len: 10,
		},
		"Valid 3": {
			tmpl: `{{ randAlphaNum 100  }}`,
			len: 100,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := execute(newTestSuite(), test.tmpl, nil)
			assert.NoError(t, err)
			assert.Equal(t, test.len, len(got))
		})
	}
}