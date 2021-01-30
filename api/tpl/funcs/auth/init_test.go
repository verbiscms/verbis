package auth

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
	"testing"
)

func TestNamespace_Init(t *testing.T) {
	//var found bool
	var ns *internal.FuncsNamespace

	for _, nsf := range internal.FuncsNamespaceRegistry {
		ns = nsf(&deps.Deps{})
		if ns.Name == name {
			//found = true
			break
		}
	}

	//assert.True(t, found)
	//assert.Equal(t, &Namespace{&deps.Deps{}}, ns.Context())
}
