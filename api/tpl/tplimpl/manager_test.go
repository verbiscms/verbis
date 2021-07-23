package tplimpl

import (
	"bytes"
	"fmt"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/errors"
)

func (t *TplTestSuite) TestNew() {
	d := &deps.Deps{}
	tm := TemplateManager{deps: d}
	t.Equal(tm, *New(d))
}

type mockWriterErr string

func (m mockWriterErr) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error")
}

func (t *TplTestSuite) TestTemplateManager_ExecuteTpl() {
	tt := map[string]struct {
		text string
		err  bool
		want interface{}
	}{
		"Success": {
			"test",
			false,
			"test",
		},
		"Bad template": {
			"{{ {{.wrong}} }}",
			false,
			"Error parsing template",
		},
		"Error parsing": {
			"test",
			true,
			"Error executing template",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			d := &deps.Deps{}
			tm := TemplateManager{deps: d}

			var (
				err error
				buf = bytes.Buffer{}
			)

			if test.err {
				err = tm.ExecuteTpl(mockWriterErr(""), test.text, nil)
			} else {
				err = tm.ExecuteTpl(&buf, test.text, nil)
			}

			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}

			t.Equal(test.want, buf.String())
		})
	}
}
