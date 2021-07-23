// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tplimpl

import (
	"fmt"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/tpl"
	"github.com/verbiscms/verbis/api/tpl/funcs/partial"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// FileHandler function describing the process for obtaining template files.
type FileHandler func(config tpl.TemplateConfig, template string) (content string, err error)

// fpAbs is an alias for filepath.Abs
var fpAbs = filepath.Abs

// DefaultFileHandler
//
// Accepts a template path and looks up the page template
// by the template path and file extension set in the
// engine.
//
// Returns errors.TEMPLATE if thee file does not exist or filepath.Abs failed.
func DefaultFileHandler() FileHandler {
	const op = "TemplateEngine.defaultFileHandler"

	return func(config tpl.TemplateConfig, template string) (content string, err error) {
		path := config.GetRoot() + string(os.PathSeparator) + template + config.GetExtension()

		fs := config.GetFS()
		if fs != nil {
			data, err := fs.ReadFile(strings.TrimPrefix(path, string(os.PathSeparator)))
			if err != nil {
				return "", err
			}
			return string(data), nil
		}

		// Get the absolute path of the root template
		abs, err := fpAbs(config.GetRoot() + string(os.PathSeparator) + template + config.GetExtension())
		if err != nil {
			return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Error obtaining absolute file of template:%v", path), Operation: op, Err: err}
		}

		data, err := ioutil.ReadFile(abs)
		if err != nil {
			return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Render read name:%v, path:%v", template, path), Operation: op, Err: err}
		}

		return string(data), nil
	}
}

// executeRender
//
// Checks whether to use the master layout and returns
// the executeTemplate function.
func (e *Execute) executeRender(out io.Writer, name string, data interface{}) (string, error) {
	useMaster := true
	if filepath.Ext(name) == e.config.GetExtension() {
		useMaster = false
		name = strings.TrimSuffix(name, e.config.GetExtension())
	}
	return e.executeTemplate(out, name, data, useMaster)
}

// executeTemplate
//
// Returns the path of the executed file, and and error
// if the was an issue writing to the io.Writer.
// Partial functions are included.
//
// Returns errors.TEMPLATE if the template file failed to execute or was unable to be parsed.
func (e *Execute) executeTemplate(out io.Writer, name string, data interface{}, useMaster bool) (string, error) {
	const op = "TemplateEngine.Execute"

	var (
		tpl *template.Template //nolint
		err error
		ok  bool
	)

	e.tplMutex.RLock()
	tpl, ok = e.tplMap[name]
	e.tplMutex.RUnlock()

	exeName := name
	if useMaster && e.config.GetMaster() != "" {
		exeName = e.config.GetMaster()
	}

	pfn := partial.Partial(e.funcMap, e)
	e.funcMap["partial"] = pfn
	e.funcMap["include"] = pfn

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
				return v, err
			}
			var tmpl *template.Template
			if v == name {
				tmpl = tpl
			} else {
				tmpl = tpl.New(v)
			}
			_, err = tmpl.Parse(data)
			if err != nil {
				return v, &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to parse template with the name: %s", v), Operation: op, Err: err}
			}
		}
		e.tplMutex.Lock()
		e.tplMap[name] = tpl
		e.tplMutex.Unlock()
	}

	// Display the content to the screen
	err = tpl.Funcs(e.funcMap).ExecuteTemplate(out, exeName, data)
	if err != nil {
		return name, &errors.Error{Code: errors.TEMPLATE, Message: "template engine execute template error", Operation: op, Err: err}
	}

	return name, nil
}
