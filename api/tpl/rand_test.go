package tpl

import (
	"fmt"
	"github.com/spf13/cast"
)

func (t *TplTestSuite) Test_RandInt() {

	tt := map[string]struct {
		tpl string
		want []int
	}{
		"Valid": {
			tpl: `{{ randInt 1 100 }}`,
			want: []int{1, 100},
		},
		"Valid 2": {
			tpl: `{{ randInt 1 5 }}`,
			want: []int{1, 5},
		},
		"Large": {
			tpl: `{{ randInt 1 99999999 }}`,
			want: []int{1, 99999999},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			str, err := t.Execute(test.tpl, nil)
			t.NoError(err)

			got := cast.ToInt(str)
			if test.want[0] > got || test.want[1] < got {
				t.Error(fmt.Errorf("got %v expecting in between %v", got, test.want))
			}
		})
	}
}

func (t *TplTestSuite) Test_RandFloat() {

	tt := map[string]struct {
		tpl string
		want []float64
	}{
		"Valid": {
			tpl: `{{ randFloat 1 100 }}`,
			want: []float64{1, 100},
		},
		"Valid2": {
			tpl: `{{ randFloat 1.555555 10.44444 }}`,
			want: []float64{1.555555, 10.44444},
		},
		"Large": {
			tpl: `{{ randInt 1 99999999.99999 }}`,
			want: []float64{1, 99999999.99999},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			str, err := t.Execute(test.tpl, nil)
			t.NoError(err)

			got := cast.ToFloat64(str)
			if test.want[0] > got || test.want[1] < got {
				t.Error(fmt.Errorf("got %v expecting in between %v", got, test.want))
			}
		})
	}
}

func (t *TplTestSuite) Test_RandAlpha() {

	tt := map[string]struct {
		tpl string
		len  int
	}{
		"Valid": {
			tpl: `{{ randAlpha 5  }}`,
			len:  5,
		},
		"Valid 2": {
			tpl: `{{ randAlpha 10  }}`,
			len:  10,
		},
		"Valid 3": {
			tpl: `{{ randAlpha 100  }}`,
			len:  100,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got, err := t.Execute(test.tpl, nil)
			t.NoError(err)
			t.Equal(test.len, len(got))
		})
	}
}

func (t *TplTestSuite) Test_RandAlphaNum() {

	tt := map[string]struct {
		tpl string
		len  int
	}{
		"Valid": {
			tpl: `{{ randAlphaNum 5  }}`,
			len:  5,
		},
		"Valid 2": {
			tpl: `{{ randAlphaNum 10  }}`,
			len:  10,
		},
		"Valid 3": {
			tpl: `{{ randAlphaNum 100  }}`,
			len:  100,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got, err := t.Execute(test.tpl, nil)
			t.NoError(err)
			t.Equal(test.len, len(got))
		})
	}
}
