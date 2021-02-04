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
	const op = "Templates.Partial"

	path := ns.tpld.Cfg.GetRoot() + "/" + name

	if !files.Exists(path) {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: "Partial file does not exist", Operation: op, Err: fmt.Errorf("no file exists with the path: %s", name)}
	}

	var context interface{}
	if len(data) == 1 {
		context = data[0]
	} else {
		context = data
	}

	pathArr := strings.Split(path, "/")
	funcs := ns.deps.Tpl.FuncMap(ns.tpld.Context, ns.tpld.Post, ns.tpld.Cfg)

	file, err := template.New(pathArr[len(pathArr)-1]).Funcs(funcs).ParseFiles(path)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: "Unable to parse partial file", Operation: op, Err: err}
	}

	var b bytes.Buffer
	err = file.Execute(&b, context)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: "Unable to execute partial file", Operation: op, Err: err}
	}

	return template.HTML(b.String()), nil
}
