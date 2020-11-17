package templates

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"html/template"
	"strings"
)

// partial - Takes in the name of a template relative to the theme
// as well as any data to be passed. The template is executed and
// panics if no file was found or the template could not be
// executed.


// TODO have paths of the tempalte functions struct.

func (t *TemplateFunctions) partial(name string, data ...interface{}) template.HTML {
	path := paths.Theme() + "/" + name

	var context interface{}
	if len(data) > 0 {
		context = data[0]
	}

	if !files.Exists(path) {
		panic(fmt.Errorf("No file exists with the path: %s", name))
	}

	pathArr := strings.Split(path, "/")
	file, err := template.New(pathArr[len(pathArr) - 1]).Funcs(t.GetFunctions()).ParseFiles(path)
	if err != nil {
		panic(fmt.Errorf("Unable to create a new partial file: %v", err))
	}

	var tpl bytes.Buffer
	err = file.Execute(&tpl, context)
	if err != nil {
		panic(err)
	}

	return template.HTML(tpl.String())
}
