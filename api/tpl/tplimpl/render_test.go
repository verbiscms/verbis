package tplimpl

import (
	"fmt"
	mocks "github.com/ainsleyclark/verbis/api/mocks/tpl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultFileHandler(t *testing.T) {

	tt := map[string]struct {
		mock     func(config *mocks.TemplateConfig)
		template string
		want     interface{}
	}{
		"Valid": {
			func(config *mocks.TemplateConfig) {
				config.On("GetRoot").Return("wrongval")
				config.On("GetExtension").Return("wrongval")
			},
			"",
			"dd",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := &mocks.TemplateConfig{}
			test.mock(m)
			fn := DefaultFileHandler()
			got, err := fn(m, test.template)

			if err != nil {
				fmt.Print(err)
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.Equal(t, test.want, got)
		})
	}
}
