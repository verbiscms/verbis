package templates

import "testing"

func Test_Add(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ add 1 2 3 }}`,
			want: "6",
		},
		"Valid 2": {
			tmpl: `{{ add 10 10 }}`,
			want: "20",
		},
		"Strings": {
			tmpl: `{{ add "10" "10" }}`,
			want: "20",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Subtract(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ subtract 10 1 }}`,
			want: "9",
		},
		"Valid 2": {
			tmpl: `{{ subtract 100 50 }}`,
			want: "50",
		},
		"Strings": {
			tmpl: `{{ subtract "10" "5" }}`,
			want: "5",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Divide(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ divide 16 2 }}`,
			want: "8",
		},
		"Valid 2": {
			tmpl: `{{ divide 100 50 }}`,
			want: "2",
		},
		"Strings": {
			tmpl: `{{ divide "10" "5" }}`,
			want: "2",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Multiply(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ multiply 10 10 }}`,
			want: "100",
		},
		"Valid 2": {
			tmpl: `{{ multiply 2 4 4 }}`,
			want: "32",
		},
		"Strings": {
			tmpl: `{{ multiply "10" "5" "2" }}`,
			want: "100",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Modulus(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ mod 10 2 }}`,
			want: "0",
		},
		"Valid 2": {
			tmpl: `{{ mod 16 3 }}`,
			want: "1",
		},
		"Strings": {
			tmpl: `{{ mod "100" "4" }}`,
			want: "0",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Round(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ round 10.1234 }}`,
			want: "10",
		},
		"Valid 2": {
			tmpl: `{{ round 16 }}`,
			want: "16",
		},
		"Strings": {
			tmpl: `{{ round "100.564988" }}`,
			want: "101",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Ceil(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ ceil 10.1234 }}`,
			want: "11",
		},
		"Valid 2": {
			tmpl: `{{ ceil 16 }}`,
			want: "16",
		},
		"Strings": {
			tmpl: `{{ ceil "100.564988" }}`,
			want: "101",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Floor(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ floor 10.1234 }}`,
			want: "10",
		},
		"Valid 2": {
			tmpl: `{{ floor 16 }}`,
			want: "16",
		},
		"Strings": {
			tmpl: `{{ floor "100.564988" }}`,
			want: "100",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Min(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ min 1 2 3 4 5 6 7 8 9 10 }}`,
			want: "1",
		},
		"Valid 2": {
			tmpl: `{{ min 102 3004 323 2848 }}`,
			want: "102",
		},
		"Smaller Comparison": {
			tmpl: `{{ min 102 2 40 2848 }}`,
			want: "2",
		},
		"Strings": {
			tmpl: `{{ min "1" "1000" }}`,
			want: "1",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}

func Test_Max(t *testing.T) {

	tt := map[string]struct {
		tmpl string
		want interface{}
	}{
		"Valid": {
			tmpl: `{{ max 1 2 3 4 5 6 7 8 9 10 }}`,
			want: "10",
		},
		"Valid 2": {
			tmpl: `{{ max 102 3004 323 2848 }}`,
			want: "3004",
		},
		"Larger Comparison": {
			tmpl: `{{ max 102 9999 40 2848 }}`,
			want: "9999",
		},
		"Strings": {
			tmpl: `{{ max "1" "1000" }}`,
			want: "1000",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.tmpl, test.want)
		})
	}
}
