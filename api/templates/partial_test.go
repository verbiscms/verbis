package templates

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"html/template"
	"os"
	"path/filepath"
	"testing"
)

func Test_Partial(t *testing.T) {

	// Save current function and restore
	oldThemePath := themePath
	defer func() {
		themePath = oldThemePath
	}()

	wd, err := os.Getwd()
	assert.NoError(t, err)
	apiPath := filepath.Join(filepath.Dir(wd))
	themePath = apiPath + "/test/testdata"

	tt := map[string]struct {
		input string
		want  string
		error string
	}{
		"Success": {
			input: `{{ partial "html/partial.cms" }}`,
			want:  `<h1>This is a partial file.</h1>`,
			error: "",
		},
		"Wrong Path": {
			input: `{{ partial "html/wrongpath.cms" }}`,
			want:  ``,
			error: "No file exists with the path: html/wrongpath.cms",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()

			tt := template.Must(template.New("test").Funcs(f.GetFunctions()).Parse(test.input))

			var b bytes.Buffer
			err := tt.Execute(&b, map[string]string{})
			if test.error != "" {
				assert.Contains(t, err.Error(), test.error)
			}

			assert.Equal(t, b.String(), test.want)
		})
	}
}
