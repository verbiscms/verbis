package tplimpl

import (
	"bytes"
	"fmt"
	mocks "github.com/verbiscms/verbis/api/mocks/tpl"
	mockfs "github.com/verbiscms/verbis/api/mocks/verbisfs"
	"html/template"
	"sync"
)

func (t *TplTestSuite) TestDefaultFileHandler() {
	tt := map[string]struct {
		mock     func(config *mocks.TemplateConfig)
		template string
		want     interface{}
	}{
		"Success": {
			func(config *mocks.TemplateConfig) {
				config.On("GetFS").Return(nil)
				config.On("GetRoot").Return("testdata")
				config.On("GetExtension").Return(".html")
				config.On("GetMaster").Return("")
			},
			"standard",
			"<h1>Verbis</h1>",
		},
		"Bad Path": {
			func(config *mocks.TemplateConfig) {
				config.On("GetFS").Return(nil)
				config.On("GetRoot").Return("wrongval")
				config.On("GetExtension").Return("wrongval")
			},
			"",
			"no such file or directory",
		},
		"With FS - Success": {
			func(config *mocks.TemplateConfig) {
				fs := &mockfs.FS{}
				fs.On("ReadFile", "root/extension").Return([]byte("test"), nil)

				config.On("GetFS").Return(fs)
				config.On("GetRoot").Return("root")
				config.On("GetExtension").Return("extension")
			},
			"",
			"test",
		},
		"With FS - Error": {
			func(config *mocks.TemplateConfig) {
				fs := &mockfs.FS{}
				fs.On("ReadFile", "root/extension").Return(nil, fmt.Errorf("error"))

				config.On("GetFS").Return(fs)
				config.On("GetRoot").Return("root")
				config.On("GetExtension").Return("extension")
			},
			"",
			"error",
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

func (t *TplTestSuite) TestDefaultFileHandler_AbsError() {
	orig := fpAbs
	defer func() { fpAbs = orig }()
	fpAbs = func(path string) (string, error) {
		return "", fmt.Errorf("error")
	}

	m := &mocks.TemplateConfig{}
	m.On("GetRoot").Return("wrongval")
	m.On("GetExtension").Return("wrongval")
	m.On("GetFS").Return(nil)
	fn := DefaultFileHandler()

	_, err := fn(m, "test")
	if err == nil {
		t.Fail("expecting error")
		return
	}

	t.Contains(err.Error(), "error")
}

func (t *TplTestSuite) TestExecute_ExecuteRender() {
	tt := map[string]struct {
		mock     func(config *mocks.TemplateConfig)
		template string
		writeErr bool
		want     interface{}
	}{
		"Success": {
			func(config *mocks.TemplateConfig) {
				config.On("GetFS").Return(nil)
				config.On("GetRoot").Return("testdata")
				config.On("GetExtension").Return(".html")
				config.On("GetMaster").Return("")
			},
			"standard",
			false,
			"<h1>Verbis</h1>",
		},
		"Bad Path": {
			func(config *mocks.TemplateConfig) {
				config.On("GetFS").Return(nil)
				config.On("GetRoot").Return("wrongval")
				config.On("GetExtension").Return("wrongval")
				config.On("GetMaster").Return("")
			},
			"",
			false,
			"no such file or directory",
		},
		"Write Error": {
			func(config *mocks.TemplateConfig) {
				config.On("GetFS").Return(nil)
				config.On("GetRoot").Return("testdata")
				config.On("GetExtension").Return(".html")
				config.On("GetMaster").Return("")
			},
			"standard",
			true,
			"",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := &mocks.TemplateConfig{}
			test.mock(m)

			e := &Execute{
				config:      m,
				tplMutex:    sync.RWMutex{},
				fileHandler: DefaultFileHandler(),
				funcMap: map[string]interface{}{
					"partial": func() {},
				},
				tplMap: make(map[string]*template.Template),
			}

			var (
				err  error
				name string
				buf  = bytes.Buffer{}
			)

			if test.writeErr {
				name, err = e.executeRender(mockWriterErr(""), test.template, nil)
			} else {
				name, err = e.executeRender(&buf, test.template, nil)
			}

			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, buf.String())
			t.Equal(test.template, name)
		})
	}
}
