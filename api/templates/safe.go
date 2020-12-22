package templates

import (
	"github.com/spf13/cast"
	"html/template"
)

// safeHTML
//
// Returns a given string as html/template HTML content
func (t *TemplateFunctions) safeHTML(i interface{}) (template.HTML, error) {
	s, err := cast.ToStringE(i)
	return template.HTML(s), err
}

// HTMLAttr
//
// Returns a given string as html/template HTMLAttr content
func (t *TemplateFunctions) safeHTMLAttr(i interface{}) (template.HTMLAttr, error) {
	s, err := cast.ToStringE(i)
	return template.HTMLAttr(s), err
}

// safeCSS
//
// Returns a given string as html/template HTML content
func (t *TemplateFunctions) safeCSS(i interface{}) (template.CSS, error) {
	s, err := cast.ToStringE(i)
	return template.CSS(s), err
}

// safeJS
//
// Returns a given string as html/template HTML content
func (t *TemplateFunctions) safeJS(i interface{}) (template.JS, error) {
	s, err := cast.ToStringE(i)
	return template.JS(s), err
}

// safeJSStr
//
// Returns the given string as a html/template JSStr content
func (t *TemplateFunctions) safeJSStr(i interface{}) (template.JSStr, error) {
	s, err := cast.ToStringE(i)
	return template.JSStr(s), err
}

// safeUrl
//
// Returns a given string as html/template URL content
func (t *TemplateFunctions) safeUrl(i interface{}) (template.URL, error) {
	s, err := cast.ToStringE(i)
	return template.URL(s), err
}
