package templates

import (
	"github.com/stretchr/testify/assert"
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
		name  string
		input string
		want  string
	}{
		"Success": {
			input: `{{ partial "html/partial.cms" }}`,
			want:  `<h1>This is a partial file.</h1>`,
		},
		"Wrong Path": {
			input: `{{ partial "html/wrongpath.cms" }}`,
			want:  `<h1>This is a partial file.</h1>`,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			runt(t, f, test.input, test.want)
		})
	}
}
