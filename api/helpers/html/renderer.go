package html

import (
	"bytes"
	"html/template"
)

// Render template executes the html and returns a string
func RenderTemplate(layout string, data interface{}, files ...string) (string, error) {
	t, err := template.New("").ParseFiles(files...)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, layout, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}