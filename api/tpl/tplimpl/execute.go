package tplimpl

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

const (
	// Delimiter left for the template
	DelimitersLeft = "{{"
	// Delimiter right for the template
	DelimitersRight = "}}"
)

// Execute defines the data for rendering a template
// contains
type Execute struct {
	*TemplateManager
	config      tpl.TemplateConfig
	tplMap      map[string]*template.Template
	tplMutex    sync.RWMutex
	fileHandler fileHandler
	funcMap     template.FuncMap
}

func (t *TemplateManager) Prepare(c tpl.TemplateConfig) tpl.TemplateExecutor {
	return &Execute{
		t,
		c,
		make(map[string]*template.Template),
		sync.RWMutex{},
		DefaultFileHandler(),
		template.FuncMap{},
	}
}

// RenderWriter
//
// Render a template with io.Writer
func (e *Execute) Execute(w io.Writer, name string, data interface{}) error {
	e.funcMap = e.GenericFuncMap()
	return e.executeRender(w, name, data)
}

// Post
//
// Is a wrapper for execute Render,
func (e *Execute) ExecutePost(w io.Writer, name string, ctx *gin.Context, post *domain.PostData) error {
	//data := e.Data(ctx, post)
	e.funcMap = e.FuncMap(ctx, post)
	return e.executeRender(w, name, nil)
}

// Exists
//
// Determines if a template file exists with the given name.
func (e *Execute) Exists(template string) bool {
	_, err := e.fileHandler(e.config, template)
	if err != nil {
		return false
	}
	return true
}

// FileHandler file handler interface
type fileHandler func(config tpl.TemplateConfig, template string) (content string, err error)

// DefaultFileHandler
//
// Accepts a template path and looks up the page template by the
// template path and file extension set in the engine.
// Returns
func DefaultFileHandler() fileHandler {
	const op = "TemplateEngine.defaultFileHandler"

	return func(config tpl.TemplateConfig, template string) (content string, err error) {
		// Get the absolute path of the root template
		path, err := filepath.Abs(config.GetRoot() + string(os.PathSeparator) + template + config.GetExtension())
		if err != nil {
			return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("No page template exists with the path:%v", path), Operation: op, Err: err}
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Render read name:%v, path:%v", template, path), Operation: op, Err: err}
		}

		return string(data), nil
	}
}