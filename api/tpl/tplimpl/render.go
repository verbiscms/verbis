package tplimpl

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/tpl"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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

func (e *Execute) executeRender(out io.Writer, name string, data interface{}) error {
	useMaster := true
	if filepath.Ext(name) == e.config.GetExtension() {
		useMaster = false
		name = strings.TrimSuffix(name, e.config.GetExtension())
	}
	return e.executeTemplate(out, name, data, useMaster)
}

func (e *Execute) executeTemplate(out io.Writer, name string, data interface{}, useMaster bool) error {
	const op = "TemplateEngine.executeTemplate"

	var tpl *template.Template
	var err error
	var ok bool

	e.tplMutex.RLock()
	tpl, ok = e.tplMap[name]
	e.tplMutex.RUnlock()

	exeName := name
	if useMaster && e.config.GetMaster() != "" {
		exeName = e.config.GetMaster()
	}

	if !ok {
		tplList := make([]string, 0)
		if useMaster {
			if e.config.GetMaster() != "" {
				tplList = append(tplList, e.config.GetMaster())
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
				return &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Template engine error parsing template with name: %s", v), Operation: op, Err: err}
			}
		}
		e.tplMutex.Lock()
		e.tplMap[name] = tpl
		e.tplMutex.Unlock()
	}

	// Display the content to the screen
	err = tpl.Funcs(e.funcMap).ExecuteTemplate(out, exeName, data)
	if err != nil {
		return &errors.Error{Code: errors.TEMPLATE, Message: "Template engine execute template error", Operation: op, Err: err}
	}

	return nil
}