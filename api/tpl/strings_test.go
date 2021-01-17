package tpl

func (t *TplTestSuite) Test_Replace() {

	tt := map[string]struct {
		tpl  string
		want interface{}
	}{
		"Valid": {
			tpl:  `{{ replace "verbis-cms-is-amazing" "-" " " }}`,
			want: "verbis cms is amazing",
		},
		"Valid 2": {
			tpl:  `{{ replace "verbis" "v" "" }}`,
			want: "erbis",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RunT(test.tpl, test.want)
		})
	}
}

func (t *TplTestSuite) Test_Substr() {

	tt := map[string]struct {
		tpl  string
		want interface{}
	}{
		"Valid": {
			tpl:  `{{ substr "verbiscms" 0 2 }}`,
			want: "ve",
		},
		"Valid 2": {
			tpl:  `{{ substr "hello world" 0 5 }}`,
			want: "hello",
		},
		"Strings as Params": {
			tpl:  `{{ substr "hello world" "0" "5" }}`,
			want: "hello",
		},
		"Negative Start": {
			tpl:  `{{ substr "hello world" "-1" "5" }}`,
			want: "hello",
		},
		"Negative End": {
			tpl:  `{{ substr "hello world" "5" "-1" }}`,
			want: " world",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RunT(test.tpl, test.want)
		})
	}
}

func (t *TplTestSuite) Test_Trunc() {

	tt := map[string]struct {
		tpl  string
		want interface{}
	}{
		"Positive": {
			tpl:  `{{ trunc "hello world" 5 }}`,
			want: "hello",
		},
		"Negative": {
			tpl:  `{{ trunc "hello world" -5 }}`,
			want: "world",
		},
		"Strings as Params": {
			tpl:  `{{ trunc "hello world" "-5" }}`,
			want: "world",
		},
		"Original": {
			tpl:  `{{ trunc "hello world" -1000 }}`,
			want: "hello world",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RunT(test.tpl, test.want)
		})
	}
}

func (t *TplTestSuite) Test_Ellipsis() {

	tt := map[string]struct {
		tpl  string
		want interface{}
	}{
		"Valid": {
			tpl:  `{{ ellipsis "hello world" 5 }}`,
			want: "hello...",
		},
		"Valid 2": {
			tpl:  `{{ ellipsis "hello world this is Verbis CMS" 11 }}`,
			want: "hello world...",
		},
		"Short String": {
			tpl:  `{{ ellipsis "cms" 3 }}`,
			want: "cms",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RunT(test.tpl, test.want)
		})
	}
}
