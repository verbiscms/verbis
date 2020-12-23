package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/spf13/cast"
	"html/template"
)

// safeHTML
//
// Returns a given string as html/template HTML content
func (t *TemplateFunctions) safeHTML(i interface{}) (template.HTML, error) {
	const op = "Templates.safeHTML"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe HTML to string"), Operation: op, Err: err}
	}
	return template.HTML(s), nil
}

// safeHTMLAttr
//
// Returns a given string as html/template HTMLAttr content
func (t *TemplateFunctions) safeHTMLAttr(i interface{}) (template.HTMLAttr, error) {
	const op = "Templates.safeHTMLAttr"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe HTMLAttr to string"), Operation: op, Err: err}
	}
	return template.HTMLAttr(s), nil
}

// safeCSS
//
// Returns a given string as html/template HTML content
func (t *TemplateFunctions) safeCSS(i interface{}) (template.CSS, error) {
	const op = "Templates.safeCSS"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe CSS to string"), Operation: op, Err: err}
	}
	return template.CSS(s), nil
}

// safeJS
//
// Returns a given string as html/template HTML content
func (t *TemplateFunctions) safeJS(i interface{}) (template.JS, error) {
	const op = "Templates.safeJS"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe JS to string"), Operation: op, Err: err}
	}
	return template.JS(s), nil
}

// safeJSStr
//
// Returns the given string as a html/template JSStr content
func (t *TemplateFunctions) safeJSStr(i interface{}) (template.JSStr, error) {
	const op = "Templates.safeJSStr"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe JSStr to string"), Operation: op, Err: err}
	}
	return template.JSStr(s), nil
}

// safeUrl
//
// Returns a given string as html/template URL content
func (t *TemplateFunctions) safeUrl(i interface{}) (template.URL, error) {
	const op = "Templates.safeUrl"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe URL to string"), Operation: op, Err: err}
	}
	return template.URL(s), nil
}
