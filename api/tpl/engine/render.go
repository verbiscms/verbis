package engine

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

func (t *TplEngine) executeRender(out io.Writer, name string, data interface{}) error {
	useMaster := true
	if filepath.Ext(name) == t.fileExtension {
		useMaster = false
		name = strings.TrimSuffix(name, t.fileExtension)
	}
	return t.executeTemplate(out, name, data, useMaster)
}

func (t *TplEngine) executeTemplate(out io.Writer, name string, data interface{}, useMaster bool) error {
	tpl := template.New(name).Funcs(t.funcMap).Delims(DelimitersLeft, DelimitersRight)

	tpl.ParseFiles()

	err := tpl.Funcs(t.funcMap).ExecuteTemplate(out, exeName, data)
	if err != nil {
		return fmt.Errorf("TemplateEngine execute template error: %v", err)
	}

	return nil
}
