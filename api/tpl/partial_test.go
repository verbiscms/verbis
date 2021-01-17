package tpl

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

func (t *TplTestSuite) Test_Partial() {

	// Save current function and restore
	oldThemePath := themePath
	defer func() {
		themePath = oldThemePath
	}()

	wd, err := os.Getwd()
	t.NoError(err)
	apiPath := filepath.Join(filepath.Dir(wd))
	themePath = apiPath + "/test/testdata"

	tt := map[string]struct {
		input string
		want  string
		data  interface{}
	}{
		"Success": {
			input: `{{ partial "html/partial.cms" }}`,
			want:  `<h1>This is a partial file.</h1>`,
		},
		"Wrong Path": {
			input: `{{ partial "html/wrongpath.cms" }}`,
			want:  "error calling partial: Templates.partial: no file exists with the path: html/wrongpath.cms",
		},
		"Bad Data": {
			input: `{{ partial "html/partial-baddata.cms" }}`,
			want:  "error calling partial: Unable to execute partial file: template: partial-baddata.cms:2:3:",
		},
		"File Type": {
			input: `{{ partial "images/gopher.png" }}`,
			want:  "error calling partial: Templates.partial: template: gopher.png:8",
		},
		"Dict": {
			input: `{{ partial "html/partial-dict.cms" (dict "Text" "cms") }}`,
			want:  "cms",
		},
		"Single Input": {
			input: `{{ partial "html/partial-data.cms" "verbis" }}`,
			want:  "verbis",
		},
		"Multiple Inputs": {
			input: `{{ partial "html/partial-data.cms" "hello" "verbis" }}`,
			want:  "[hello verbis]",
		},
		"Multiple Inputs 2": {
			input: `{{ partial "html/partial-data.cms" "hello" "verbis" 1 2 3 }}`,
			want:  "[hello verbis 1 2 3]",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			tt := template.Must(template.New("test").Funcs(t.GetFunctions()).Parse(test.input))

			var b bytes.Buffer
			if err := tt.Execute(&b, test.data); err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, b.String())
		})
	}
}
