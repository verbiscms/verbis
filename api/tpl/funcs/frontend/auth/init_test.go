package auth

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNamespace_Init(t *testing.T) {
	d := &deps.Deps{}
	ctx := &gin.Context{}
	td := &internal.TemplateDeps{Context: ctx}

	ns := Init(d, td)
	assert.Equal(t, ns.Name, name)
	assert.Equal(t, &Namespace{deps: d, ctx: ctx}, ns.Context())
}
