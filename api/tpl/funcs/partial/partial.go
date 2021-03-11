// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package partial

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/tpl"
	"html/template"
	"strings"
)

// Func describes the function for includes or rendering
// partials within the template.
type Func func(name string, data ...interface{}) (template.HTML, error)

// Partial
//
// Takes in the name of a template relative to the theme as well
// as any data to be passed. The template is executed and
// returns an error if no

// Returns errors.TEMPLATE if no file was found or the template
// could not be executed.
//
// Example: {{ partial "partials/circle.svg" (dict "radius" 50 "fill" "red") }}
func Partial(tplFuncs template.FuncMap, exec tpl.TemplateExecutor) Func {
	const op = "Templates.Partial"

	return func(name string, data ...interface{}) (template.HTML, error) {
		path := exec.Config().GetRoot() + "/" + name

		var context interface{}
		if len(data) == 1 {
			context = data[0]
		} else {
			context = data
		}

		pathArr := strings.Split(path, "/")

		file, err := template.New(pathArr[len(pathArr)-1]).Funcs(tplFuncs).ParseFiles(path)
		if err != nil {
			return "", &errors.Error{Code: errors.TEMPLATE, Message: "Partial file does not exist", Operation: op, Err: fmt.Errorf("no file exists with the path: %s", name)}
		}

		var b bytes.Buffer
		err = file.Execute(&b, context)
		if err != nil {
			return "", &errors.Error{Code: errors.TEMPLATE, Message: "Unable to execute partial file", Operation: op, Err: err}
		}

		return template.HTML(b.String()), nil
	}
}
