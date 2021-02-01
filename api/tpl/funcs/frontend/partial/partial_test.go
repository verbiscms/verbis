package partial

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/generic/dict"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/stretchr/testify/assert"
	"html/template"
	"os"
	"path/filepath"
	"testing"
)

func Setup(t *testing.T) template.FuncMap {
	wd, err := os.Getwd()
	assert.NoError(t, err)
	apiPath := filepath.Join(filepath.Dir(wd), "../../..")

	d := &deps.Deps{
		Paths: deps.Paths{
			Theme:   apiPath + "/test/testdata",
		},
	}

	ns := New(d, &internal.TemplateDeps{})

	funcs := template.FuncMap{
		"partial": ns.Partial,
		"dict": dict.New(d).Dict,
	}

	return funcs
}

func TestNamespace_Partial(t *testing.T) {

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
			want:  "error calling partial: Templates.Partial: no file exists with the path: html/wrongpath.cms",
		},
		"Bad Data": {
			input: `{{ partial "html/partial-baddata.cms" }}`,
			want:  "error calling partial: Unable to execute partial file: template: partial-baddata.cms:2:3:",
		},
		"File Type": {
			input: `{{ partial "images/gopher.png" }}`,
			want:  "error calling partial: Templates.Partial: template: gopher.png:8",
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
		t.Run(name, func(t *testing.T) {
			fnc := Setup(t)

			file := template.Must(template.New("test").Funcs(fnc).Parse(test.input))

			var b bytes.Buffer
			if err := file.Execute(&b, test.data); err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.Equal(t, test.want, b.String())
		})
	}
}
