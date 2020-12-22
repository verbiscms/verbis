package templates

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Env(t *testing.T) {

	tt := map[string]struct {
		env  func() error
		tmpl string
		want string
	}{
		"Valid": {
			env:  func() error { return os.Setenv("verbis", "cms") },
			tmpl: `{{ env "verbis" }}`,
			want: "cms",
		},
		"Valid 2": {
			env:  func() error { return os.Setenv("foo", "bar") },
			tmpl: `{{ env "foo" }}`,
			want: "bar",
		},
		"Not found": {
			env:  func() error { return nil },
			tmpl: `{{ env "cms" }}`,
			want: "",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()

			err := test.env()
			assert.NoError(t, err)

			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_ExpandEnv(t *testing.T) {

	tt := map[string]struct {
		env  func() error
		tmpl string
		want string
	}{
		"Valid": {
			env:  func() error { return os.Setenv("path", "verbis") },
			tmpl: `{{ expandEnv "$path is my name" }}`,
			want: "verbis is my name",
		},
		"Valid 2": {
			env:  func() error { return os.Setenv("foo", "bar") },
			tmpl: `{{ expandEnv "hello $foo" }}`,
			want: "hello bar",
		},
		"Not found": {
			env:  func() error { return nil },
			tmpl: `{{ expandEnv "hello $test verbis" }}`,
			want: "hello  verbis",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()

			err := test.env()
			assert.NoError(t, err)

			runt(t, f, test.tmpl, test.want)
		})
	}
}
