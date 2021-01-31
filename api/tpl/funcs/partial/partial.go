package partial

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"html/template"
	"strings"
)

// Partial
//
// Takes in the name of a template relative to the theme as well
// as any data to be passed. The template is executed and
// returns an error if no

// Returns errors.TEMPLATE if no file was found or the template
// could not be executed.
//
// Example: {{ partial "partials/circle.svg" (dict "radius" 50 "fill" "red") }}
func (ns *Namespace) Partial(name string, data ...interface{}) (template.HTML, error) {
	const op = "Templates.partial"

	path := ns.deps.Paths.Theme + "/" + name

	var context interface{}
	if len(data) == 1 {
		context = data[0]
	} else {
		context = data
	}

	if !files.Exists(path) {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: "Partial file does not exist", Operation: op, Err: fmt.Errorf("no file exists with the path: %s", name)}
	}

	pathArr := strings.Split(path, "/")
	//file, err := template.New(pathArr[len(pathArr)-1]).Funcs(core.GetFuncMap(ns.deps)).ParseFiles(path)
	file, err := template.New(pathArr[len(pathArr)-1]).ParseFiles(path)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: "Unable to parse partial file", Operation: op, Err: err}
	}

	var tpl bytes.Buffer
	err = file.Execute(&tpl, context)
	if err != nil {
		return "", fmt.Errorf("Unable to execute partial file: %v", err)
	}

	return template.HTML(tpl.String()), nil
}
