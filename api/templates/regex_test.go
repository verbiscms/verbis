package templates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RegexMatch(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"True 1": {
			tmpl: `{{ regexMatch "[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}" "test@verbiscms.com" }}`,
			want: true,
		},
		"True 2": {
			tmpl: `{{ regexMatch "[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}" "TesT@VERBISCMS.COM" }}`,
			want: true,
		},
		"False 1": {
			tmpl: `{{ regexMatch "[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}" "verbis" }}`,
			want: false,
		},
		"False 2": {
			tmpl: `{{ regexMatch "[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}" "verbis.com" }}`,
			want: false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_RegexFindAll(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Length 1": {
			tmpl: `{{ regexFindAll "v{2}" "vvvvvv" 1 }}`,
			want: "[vv]",
		},
		"Length 2": {
			tmpl: `{{ regexFindAll "v{2}" "vv" -1 }}`,
			want: "[vv]",
		},
		"None": {
			tmpl: `{{ regexFindAll "v{2}" "none" -1 }}`,
			want: "[]",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_RegexFind(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Found 1": {
			tmpl: `{{ regexFind "verbis.?" "verbis" }}`,
			want: "verbis",
		},
		"Found 2": {
			tmpl: `{{ regexFind "verbis.?" "verbiscmsverrbis" }}`,
			want: "verbisc",
		},
		"None": {
			tmpl: `{{ regexFind "verbis.?" "none" }}`,
			want: "",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_RegexReplaceAll(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"1": {
			tmpl: `{{ regexReplaceAll "a(x*)b" "-ab-axxb-" "${1}W" }}`,
			want: "-W-xxW-",
		},
		"2": {
			tmpl: `{{ regexReplaceAll "a(x*)b" "-ab-ab-" "${1}W" }}`,
			want: "-W-W-",
		},
		"3": {
			tmpl: `{{ regexReplaceAll "a(x*)b" "ababababab" "${1}W" }}`,
			want: "WWWWW",
		},
		"4": {
			tmpl: `{{ regexReplaceAll "a(x*)b" "----" "${1}W" }}`,
			want: "----",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_RegexReplaceAllLiteral(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"1": {
			tmpl: `{{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}" }}`,
			want: "-${1}-${1}-",
		},
		"2": {
			tmpl: `{{ regexReplaceAllLiteral "a(x*)b" "-ab-ab-" "${1}" }}`,
			want: "-${1}-${1}-",
		},
		"3": {
			tmpl: `{{ regexReplaceAllLiteral "a(x*)b" "ababababab" "${1}" }}`,
			want: "${1}${1}${1}${1}${1}",
		},
		"4": {
			tmpl: `{{ regexReplaceAllLiteral "a(x*)b" "----" "${1}" }}`,
			want: "----",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_RegexSplit(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Positive": {
			tmpl: `{{ regexSplit "v" "verbis" 1 }}`,
			want: "[verbis]",
		},
		"Negative": {
			tmpl: `{{ regexSplit "v" "verbis" -1 }}`,
			want: "[ erbis]",
		},
		"Multiple": {
			tmpl: `{{ regexSplit "v" "vvvvvvv" -1 }}`,
			want: "[       ]",
		},
		"None": {
			tmpl: `{{ regexSplit "v" "none" -1 }}`,
			want: "[none]",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_RegexQuoteMeta(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Stripped": {
			tmpl: `{{ regexQuoteMeta "verbis+" }}`,
			want: "verbis\\&#43;",
		},
		"None": {
			tmpl: `{{ regexQuoteMeta "verbis" }}`,
			want: "verbis",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			b, err := execute(f, test.tmpl, nil)
			assert.NoError(t, err)
			assert.Equal(t, b, test.want)
		})
	}
}
