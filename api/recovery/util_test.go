package recovery

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func Test_TplLineNumber(t *testing.T) {

	tt := map[string]struct {
		input *errors.Error
		want  int
	}{
		"Found": {
			&errors.Error{Err: fmt.Errorf(`template: templates/home:4: function "wrong" not defined`)},
			4,
		},
		"Found Second": {
			&errors.Error{Err: fmt.Errorf(`template: templates/home:10: function "wrong" not defined`)},
			10,
		},
		"Not Found": {
			&errors.Error{Err: fmt.Errorf(`template: templates/home10: function "wrong" not defined`)},
			-1,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := tplLineNumber(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_TplFileContents(t *testing.T) {

	wd, err := os.Getwd()
	assert.NoError(t, err)
	apiPath := filepath.Join(filepath.Dir(wd))

	tt := map[string]struct {
		input string
		want  string
	}{
		"Found": {
			apiPath + "/test/testdata/html/partial.cms",
			"<h1>This is a partial file.</h1>",
		},
		"Not Found": {
			"wrong path",
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := tplFileContents(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}