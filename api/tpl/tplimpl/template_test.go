package tplimpl

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl"
	"html/template"
	"sync"
)

const (
	// The path of test templates to be executed.
	root = "/test/testdata/tpl"
)

func (t *TplTestSuite) TestTemplateManager_Prepare() {
	tm := TemplateManager{deps: &deps.Deps{}}
	config := tpl.Config{}

	got := tm.Prepare(config)
	want := &Execute{
		&tm,
		config,
		make(map[string]*template.Template),
		sync.RWMutex{},
		DefaultFileHandler(),
		template.FuncMap{},
	}

	t.Equal(want.config, got.Config())
}

func (t *TplTestSuite) TestExecute_Execute() {
	tt := map[string]struct {
		config      tpl.Config
		name        string
		data        interface{}
		fileHandler FileHandler
		want        interface{}
		wantName    string
	}{
		"Simple": {
			tpl.Config{Extension: ".html"},
			"standard",
			nil,
			nil,
			"<h1>Verbis</h1>",
			"standard",
		},
		"Extension": {
			tpl.Config{Extension: ".html"},
			"standard.html",
			nil,
			nil,
			"<h1>Verbis</h1>",
			"standard",
		},
		"Error": {
			tpl.Config{Extension: ".html"},
			"error",
			nil,
			nil,
			"TemplateEngine.Execute: template: error:1: function \"wrongfunc\" not defined",
			"error",
		},
		"Master": {
			tpl.Config{Extension: ".html", Master: "layout"},
			"child",
			nil,
			nil,
			"<h1>Verbis</h1>",
			"child",
		},
		"File Handler Error": {
			tpl.Config{Extension: ".html"},
			"standard",
			nil,
			func(config tpl.TemplateConfig, template string) (content string, err error) {
				return "", fmt.Errorf("error")
			},
			"error",
			"standard",
		},
		"With Data": {
			tpl.Config{Extension: ".html"},
			"data",
			"verbis",
			nil,
			"verbis",
			"data",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			tm, ctx, post := t.Setup()

			test.config.Root = t.TestPath
			execute := Execute{
				tm,
				test.config,
				make(map[string]*template.Template),
				sync.RWMutex{},
				DefaultFileHandler(),
				template.FuncMap{},
			}

			if test.fileHandler != nil {
				execute.fileHandler = test.fileHandler
			}

			// Normal
			t.Run("Normal", func() {
				normalBuf := &bytes.Buffer{}
				normalPath, err := execute.Execute(normalBuf, test.name, test.data)
				t.Equal(normalPath, test.wantName)
				if err != nil {
					t.Contains(err.Error(), test.want)
					return
				}
				t.Equal(test.want, normalBuf.String())
			})

			// post
			t.Run("post", func() {
				if name == "With Data" {
					return
				}

				postBuf := &bytes.Buffer{}
				postPath, err := execute.ExecutePost(postBuf, test.name, ctx, post)
				t.Equal(postPath, test.wantName)
				if err != nil {
					t.Contains(err.Error(), test.want)
					return
				}
				t.Equal(test.want, postBuf.String())
			})
		})
	}
}

func (t *TplTestSuite) TestExecute_Exists() {
	tt := map[string]struct {
		handler FileHandler
		want    bool
	}{
		"Exists": {
			func(config tpl.TemplateConfig, template string) (content string, err error) {
				return "test", nil
			},
			true,
		},
		"Not Found": {
			func(config tpl.TemplateConfig, template string) (content string, err error) {
				return "", fmt.Errorf("error")
			},
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			e := Execute{fileHandler: test.handler}
			got := e.Exists("test")
			t.Equal(test.want, got)
		})
	}
}

func (t *TplTestSuite) TestExecute_Config() {
	cfg := tpl.Config{
		Root: "test",
	}
	e := Execute{config: cfg}
	t.Equal(cfg, e.Config())
}

func (t *TplTestSuite) TestExecute_Executor() {
	e := &Execute{}
	t.Equal(e, e.Executor())
}
