package tplimpl

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	d := &deps.Deps{}
	tm := TemplateManager{deps: d}
	assert.Equal(t, tm, *New(d))
}
