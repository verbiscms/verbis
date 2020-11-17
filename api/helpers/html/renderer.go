package html

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/errors"
	"html/template"
)

// RenderTemplate executes the html and returns a string
// Returns errors.INTERNAL if the template failed to be created
// or be executed.
func RenderTemplate(layout string, data interface{}, files ...string) (string, error) {
	const op = "html.RenderTemplate"

	t, err := template.New("").ParseFiles(files...)
	if err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Unable to create a new template", Operation: op, Err: err}
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, layout, data); err != nil {
		return "", &errors.Error{Code: errors.INTERNAL, Message: "Unable to render the template", Operation: op, Err: err}
	}

	return tpl.String(), nil
}
