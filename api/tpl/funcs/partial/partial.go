// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package partial

import (
	"bytes"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/tpl"
	"html/template"
	"path/filepath"
	"strings"
)

// Func describes the function for includes or rendering
// partials within the template.
type Func func(name string, data ...interface{}) (template.HTML, error)

// newTpl is an alias for template.New
var newTpl = template.New

// Partial
//
// Takes in the name of a template relative to the theme as well
// as any data to be passed. The template is executed and
// returns an error if no
//
// Returns errors.TEMPLATE if no file was found or the template
// could not be executed.
//
// Example: {{ partial "partials/circle.svg" (dict "radius" 50 "fill" "red") }}
func Partial(tplFuncs template.FuncMap, exec tpl.TemplateExecutor) Func {
	const op = "Templates.Partial"

	return func(name string, data ...interface{}) (template.HTML, error) {
		path := filepath.Join(exec.Config().GetRoot(), name)

		var context interface{}
		if len(data) == 1 {
			context = data[0]
		} else {
			context = data
		}

		pathArr := strings.Split(path, "/")

		file, err := newTpl(pathArr[len(pathArr)-1]).Funcs(tplFuncs).ParseFiles(path)
		if err != nil {
			return "", &errors.Error{Code: errors.TEMPLATE, Message: "Error parsing partial", Operation: op, Err: err}
		}

		var b bytes.Buffer
		err = file.Execute(&b, context)
		if err != nil {
			return "", &errors.Error{Code: errors.TEMPLATE, Message: "Error executing partial", Operation: op, Err: err}
		}

		return template.HTML(b.String()), nil
	}
}
