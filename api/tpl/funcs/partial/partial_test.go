package partial

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/stretchr/testify/assert"
	"html/template"
	"os"
	"path/filepath"
	"testing"
)

const (
	TestPath = "/test/testdata"
)

func Setup(t *testing.T) *Namespace {
	wd, err := os.Getwd()
	assert.NoError(t, err)
	apiPath := filepath.Join(filepath.Dir(wd), "../../..")

	d := &deps.Deps{
		Paths: deps.Paths{
			Theme: apiPath + TestPath,
		},
	}

	return New(d, &internal.TemplateDeps{})
}

func TestNamespace_Partial(t *testing.T) {

	tt := map[string]struct {
		name string
		data interface{}
		want interface{}
	}{
		"Success": {
			`html/partial.cms`,
			nil,
			template.HTML(`<h1>This is a partial file.</h1>`),
		},
		"Wrong Path": {
			`html/wrongpath.cms`,
			nil,
			"Templates.Partial: no file exists with the path: html/wrongpath.cms",
		},
		"Bad Data": {
			`html/partial-baddata.cms`,
			nil,
			template.HTML(""),
		},
		"File Type": {
			`images/gopher.png`,
			nil,
			template.HTML(""),
		},
		"Dict": {
			`html/partial-dict.cms`,
			map[string]interface{}{"Text": "cms"},
			template.HTML("cms"),
		},
		"Single Input": {
			`html/partial-data.cms`,
			"verbis",
			template.HTML("verbis"),
		},
		"Multiple Inputs": {
			`html/partial-data.cms`,
			[]interface{}{"hello", "verbis"},
			template.HTML("[hello verbis]"),
		},
		"Multiple Inputs 2": {
			`html/partial-data.cms`,
			[]interface{}{"hello", "verbis", 1, 2, 3},
			template.HTML("[hello verbis 1 2 3]"),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns := Setup(t)
			got, err := ns.Partial(test.name, test.data)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
