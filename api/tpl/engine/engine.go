package engine

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var (
	// HTMLContentType for rendering the template header
	htmlContentType = []string{"text/html; charset=utf-8"}
)

const (
	// Delimiter left for the template
	DelimitersLeft = "{{"
	// Delimiter right for the template
	DelimitersRight = "}}"
)

type TplEngine struct {
	deps          *deps.Deps
	Post          *domain.PostData
	FuncMap       template.FuncMap
	fileHandler   FileHandler
	templatePath  string
	layoutPath    string
	fileExtension string
}

// FileHandler file handler interface
type FileHandler func(file string) (content string, err error)

// Render
//
// Render a template with http.ResponseWriter
func (t *TplEngine) Render(w http.ResponseWriter, statusCode int, name string, data interface{}) error {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = htmlContentType
	}
	w.WriteHeader(statusCode)
	return t.executeRender(w, name, data)
}

// RenderWriter
//
// Render a template with io.Writer
func (t *TplEngine) RenderWriter(w io.Writer, name string, data interface{}) error {
	return t.executeRender(w, name, data)
}

// Exists
//
// Determines if a template file exists with the given name.
func (t *TplEngine) Exists(template string) bool {
	_, err := t.fileHandler(template)
	if err != nil {
		return false
	}
	return true
}

// DefaultFileHandler
//
// Accepts a template path and looks up the page template by the
// template path and file extension set in the engine.
// Returns
func (t *TplEngine) DefaultFileHandler() FileHandler {
	const op = "TemplateEngine.defaultFileHandler"

	return func(template string) (content string, err error) {
		// Get the absolute path of the root template
		path, err := filepath.Abs(t.templatePath + string(os.PathSeparator) + template + t.fileExtension)
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
