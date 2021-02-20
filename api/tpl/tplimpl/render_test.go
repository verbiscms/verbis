package tplimpl

import (
	mocks "github.com/ainsleyclark/verbis/api/mocks/tpl"
)

func (t *TplTestSuite) TestDefaultFileHandler() {

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
		t.Run(name, func() {
			m := &mocks.TemplateConfig{}
			test.mock(m)
			fn := DefaultFileHandler()
			got, err := fn(m, test.template)

			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *TplTestSuite) TestExecute_ExecuteRender() {

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
		t.Run(name, func() {
			m := &mocks.TemplateConfig{}
			test.mock(m)
			fn := DefaultFileHandler()
			got, err := fn(m, test.template)

			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}
