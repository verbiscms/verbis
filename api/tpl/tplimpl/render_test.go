package tplimpl

import (
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
		"Success": {
			func(config *mocks.TemplateConfig) {
				config.On("GetRoot").Return("wrongval")
				config.On("GetExtension").Return("wrongval")
			},
			"",
			"no such file or directory",
		},
		"Bad Path": {
			func(config *mocks.TemplateConfig) {
				config.On("GetRoot").Return("wrongval")
				config.On("GetExtension").Return("wrongval")
			},
			"",
			"no such file or directory",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := &mocks.TemplateConfig{}
			test.mock(m)
			fn := DefaultFileHandler()
			got, err := fn(m, test.template)

			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.Equal(t, test.want, got)
		})
	}
}


func TestExecute_ExecuteRender(t *testing.T) {

	tt := map[string]struct {
		mock     func(config *mocks.TemplateConfig)
		template string
		want     interface{}
	}{
		"Success": {
			func(config *mocks.TemplateConfig) {
				config.On("GetRoot").Return("wrongval")
				config.On("GetExtension").Return("wrongval")
			},
			"",
			"no such file or directory",
		},
		"Bad Path": {
			func(config *mocks.TemplateConfig) {
				config.On("GetRoot").Return("wrongval")
				config.On("GetExtension").Return("wrongval")
			},
			"",
			"no such file or directory",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := &mocks.TemplateConfig{}
			test.mock(m)
			fn := DefaultFileHandler()
			got, err := fn(m, test.template)

			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.Equal(t, test.want, got)
		})
	}
}