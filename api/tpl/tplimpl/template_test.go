package tplimpl

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplateManager_Prepare(t *testing.T) {
	//tm := TemplateManager{deps: &deps.Deps{}}
	//config := tpl.Config{}

	//got := tm.Prepare(config)
	//want := &Execute{
	//	&tm,
	//	config,
	//	make(map[string]*template.Template),
	//	sync.RWMutex{},
	//	DefaultFileHandler(),
	//	template.FuncMap{},
	//}
//	assert.Equal(t, want, got)
}

func TestExecute_Execute(t *testing.T) {
	tm := TemplateManager{deps: &deps.Deps{}}
	got, want := tm.Prepare(tpl.Config{}).Execute(&bytes.Buffer{}, "test", "")

	fmt.Println(got, want)
}

func TestExecute_Exists(t *testing.T) {

	tt := map[string]struct {
		handler fileHandler
		want bool
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
		t.Run(name, func(t *testing.T) {
			e := Execute{fileHandler:     test.handler}
			got := e.Exists("test")
			assert.Equal(t, test.want, got)
		})
	}
}

func TestExecute_Config(t *testing.T) {
	cfg := tpl.Config{
		Root: "test",
	}
	e := Execute{config: cfg}
	assert.Equal(t, cfg, e.Config())
}
//
//func TestExecute_Executor(t *testing.T) {
//	e := Execute{}
//	assert.EqualValues(t, e, e.Executor())
//}