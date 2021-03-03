package tplimpl

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
)

func (t *TplTestSuite) TestNew() {
	d := &deps.Deps{}
	tm := TemplateManager{deps: d}
	t.Equal(tm, *New(d))
}

func (t *TplTestSuite) TestTemplateManager_ExecuteTpl() {
	tt := map[string]struct {
		text string
		data interface{}
		want interface{}
	}{
		"Success": {
			text: "test",
			data: nil,
			want: "test",
		},
		"Bad Template": {
			text: "{{ {{.wrong}} }}",
			data: nil,
			want: "Error parsing template",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			d := &deps.Deps{}
			tm := TemplateManager{deps: d}

			b := bytes.Buffer{}
			err := tm.ExecuteTpl(&b, test.text, test.data)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}

			t.Equal(test.want, b.String())
		})
	}
}
