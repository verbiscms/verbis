// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNamespace_RegexMatch(t *testing.T) {

	tt := map[string]struct {
		regex string
		str   string
		want  bool
	}{
		"True 1": {
			`[A-Za-z0-9._%+-]`,
			"test@verbiscms.com",
			true,
		},
		"True 2": {
			`[A-Za-z0-9._%+-]`,
			"TesT@VERBISCMS.COM",
			true,
		},
		"False 1": {
			`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}`,
			"verbis",
			false,
		},
		"False 2": {
			`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}`,
			"verbis.com",
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Match(test.regex, test.str)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_RegexFindAll(t *testing.T) {

	var s []string

	tt := map[string]struct {
		regex string
		str   string
		n     int
		want  interface{}
	}{
		"Length 1": {
			"v{2}",
			"vvvvvv",
			1,
			[]string{"vv"},
		},
		"Length 2": {
			"v{2}",
			"vv",
			-1,
			[]string{"vv"},
		},
		"None": {
			"v{2}",
			"none",
			-1,
			s,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.FindAll(test.regex, test.str, test.n)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_RegexFind(t *testing.T) {

	tt := map[string]struct {
		regex string
		str   string
		want  interface{}
	}{
		"Found 1": {
			"verbis.?",
			"verbis",
			"verbis",
		},
		"Found 2": {
			"verbis.?",
			"verbiscmsverrbis",
			"verbisc",
		},
		"None": {
			"verbis.?",
			"none",
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Find(test.regex, test.str)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_RegexReplaceAll(t *testing.T) {

	tt := map[string]struct {
		regex string
		str   string
		repl  string
		want  interface{}
	}{
		"1": {
			"a(x*)b",
			"-ab-axxb-",
			"${1}W",
			"-W-xxW-",
		},
		"2": {
			"a(x*)b",
			"-ab-ab-",
			"${1}W",
			"-W-W-",
		},
		"3": {
			"a(x*)b",
			"ababababab",
			"${1}W",
			"WWWWW",
		},
		"4": {
			"a(x*)b",
			"----",
			"${1}W",
			"----",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.ReplaceAll(test.regex, test.str, test.repl)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_RegexReplaceAllLiteral(t *testing.T) {

	tt := map[string]struct {
		regex string
		str   string
		repl  string
		want  interface{}
	}{
		"1": {
			"a(x*)b",
			"-ab-axxb-",
			"${1}",
			"-${1}-${1}-",
		},
		"2": {
			"a(x*)b",
			"-ab-ab-",
			"${1}",
			"-${1}-${1}-",
		},
		"3": {
			"a(x*)b",
			"ababababab",
			"${1}",
			"${1}${1}${1}${1}${1}",
		},
		"4": {
			"a(x*)b",
			"----",
			"${1}",
			"----",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.ReplaceAllLiteral(test.regex, test.str, test.repl)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_RegexSplit(t *testing.T) {

	tt := map[string]struct {
		regex string
		str   string
		i     int
		want  []string
	}{
		"Positive": {
			"v",
			"verbis",
			1,
			[]string{"verbis"},
		},
		"Negative": {
			"v",
			"verbis",
			-1,
			[]string{"", "erbis"},
		},
		"Multiple": {
			"v",
			"vvvvvvv",
			-1,
			[]string{"", "", "", "", "", "", "", ""},
		},
		"None": {
			"v",
			"none",
			-1,
			[]string{"none"},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Split(test.regex, test.str, test.i)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_RegexQuoteMeta(t *testing.T) {

	tt := map[string]struct {
		input string
		want  interface{}
	}{
		"Stripped": {
			"verbis+",
			"verbis\\+",
		},
		"None": {
			"verbis",
			"verbis",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.QuoteMeta(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
