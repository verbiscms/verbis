package templates

import "testing"

func Test_Len(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ len "verbis" }}`,
			want: "6",
		},
		"Valid 2": {
			tmpl: `{{ len "verbis cms" }}`,
			want: "10",
		},
		"Int": {
			tmpl: `{{ len 1234 }}`,
			want: 4,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Replace(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ replace "verbis-cms-is-amazing" "-" " " }}`,
			want: "verbis cms is amazing",
		},
		"Valid 2": {
			tmpl: `{{ replace "verbis" "v" "" }}`,
			want: "erbis",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Substr(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ substr "verbiscms" 0 2 }}`,
			want: "ve",
		},
		"Valid 2": {
			tmpl: `{{ substr "hello world" 0 5 }}`,
			want: "hello",
		},
		"Strings as Params": {
			tmpl: `{{ substr "hello world" "0" "5" }}`,
			want: "hello",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Trunc(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Positive": {
			tmpl: `{{ trunc "hello world" 5 }}`,
			want: "hello",
		},
		"Negative": {
			tmpl: `{{ trunc "hello world" -5 }}`,
			want: "world",
		},
		"Strings as Params": {
			tmpl: `{{ trunc "hello world" "-5" }}`,
			want: "world",
		},
		"Original": {
			tmpl: `{{ trunc "hello world" -1000 }}`,
			want: "hello world",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Ellipsis(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ ellipsis "hello world" 5 }}`,
			want: "hello...",
		},
		"Valid 2": {
			tmpl: `{{ ellipsis "hello world this is Verbis CMS" 11 }}`,
			want: "hello world...",
		},
		"Short String": {
			tmpl: `{{ ellipsis "cms" 3 }}`,
			want: "cms",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}
