package tplimpl

import (
	"github.com/ainsleyclark/verbis/api/deps"
)

func (t *TplTestSuite) TestNew() {
	d := &deps.Deps{}
	tm := TemplateManager{deps: d}
	t.Equal(tm, *New(d))
}
