package engine

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	// Delimiter left for the template
	DelimitersLeft = "{{"
	// Delimiter right for the template
	DelimitersRight = "}}"
)

type TplEngine struct {
	deps          *deps.Deps
}

// Config configuration options
type Config struct {
	Root         string           //view root
	Extension    string           //template extension
	Master       string           //template master
}

// FileHandler file handler interface
type FileHandler func(config Config, template string) (content string, err error)


type Execute struct {
	config     Config
	deps *deps.Deps
	funcMap       template.FuncMap
	tplMap      map[string]*template.Template
	tplMutex    sync.RWMutex
	fileHandler   FileHandler
}

func (t *TplEngine) Execute(c Config) *Execute {
	return &Execute{
		config:      c,
		funcMap:     nil,
		tplMap:      make(map[string]*template.Template),
		tplMutex:    sync.RWMutex{},
		fileHandler: DefaultFileHandler(),
	}
}

// RenderWriter
//
// Render a template with io.Writer
func (e *Execute) Execute(w io.Writer, name string, data interface{}) error {
	return e.executeRender(w, name, data)
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

func (e *Execute) executeRender(out io.Writer, name string, data interface{}) error {
	useMaster := true
	if filepath.Ext(name) == e.config.Extension {
		useMaster = false
		name = strings.TrimSuffix(name, e.config.Extension)
	}
	return e.executeTemplate(out, name, data, useMaster)
}

func (e *Execute) executeTemplate(out io.Writer, name string, data interface{}, useMaster bool) error {
	var tpl *template.Template
	var err error
	var ok bool


	e.tplMutex.RLock()
	tpl, ok = e.tplMap[name]
	e.tplMutex.RUnlock()

	exeName := name
	if useMaster && e.config.Master != "" {
		exeName = e.config.Master
	}

	if !ok {
		tplList := make([]string, 0)
		if useMaster {
			//render()
			if e.config.Master != "" {
				tplList = append(tplList, e.config.Master)
			}
		}
		tplList = append(tplList, name)

		// Loop through each template and test the full path
		tpl = template.New(name).Funcs(e.funcMap).Delims(DelimitersLeft, DelimitersRight)
		for _, v := range tplList {
			var data string
			data, err = e.fileHandler(e.config, v)
			if err != nil {
				return err
			}
			var tmpl *template.Template
			if v == name {
				tmpl = tpl
			} else {
				tmpl = tpl.New(v)
			}
			_, err = tmpl.Parse(data)
			if err != nil {
				return fmt.Errorf("ViewEngine render parser name:%v, error: %v", v, err)
			}
		}
		e.tplMutex.Lock()
		e.tplMap[name] = tpl
		e.tplMutex.Unlock()
	}

	// Display the content to the screen
	err = tpl.Funcs(e.funcMap).ExecuteTemplate(out, exeName, data)
	if err != nil {
		return fmt.Errorf("ViewEngine execute template error: %v", err)
	}

	return nil
}

// DefaultFileHandler
//
// Accepts a template path and looks up the page template by the
// template path and file extension set in the engine.
// Returns
func DefaultFileHandler() FileHandler {
	const op = "TemplateEngine.defaultFileHandler"

	return func(config Config, template string) (content string, err error) {
		// Get the absolute path of the root template
		path, err := filepath.Abs(config.Root + string(os.PathSeparator) + template + config.Extension)
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
