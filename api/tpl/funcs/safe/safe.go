package safe

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/spf13/cast"
	"html/template"
)

// HTML
//
// Returns a given string as html/template HTML content
// Returns errors.TEMPLATE if the inputted interface failed to be cast.
//
// Example: {{ "<p>verbis&cms</p>" | safeHTML }}
// Returns: `verbis&cms`
func (ns *Namespace) HTML(i interface{}) (template.HTML, error) {
	const op = "Templates.HTML"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe HTML to string"), Operation: op, Err: err}
	}
	return template.HTML(s), nil
}

// HTMLAttr
//
// Returns a given string as html/template HTMLAttr content.
// Returns errors.TEMPLATE if the inputted interface failed to be cast.
func (ns *Namespace) HTMLAttr(i interface{}) (template.HTMLAttr, error) {
	const op = "Templates.HTMLAttr"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe HTMLAttr to string"), Operation: op, Err: err}
	}
	return template.HTMLAttr(s), nil
}

// CSS
//
// Returns a given string as html/template HTML content.
// Returns errors.TEMPLATE if the inputted interface failed to be cast.
//
// Example: {{ "<p>verbis&cms</p>" | safeHTML }}
// Returns: `verbis&amp;cms`
func (ns *Namespace) CSS(i interface{}) (template.CSS, error) {
	const op = "Templates.CSS"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe CSS to string"), Operation: op, Err: err}
	}
	return template.CSS(s), nil
}

// JS
//
// Returns a given string as html/template HTML content.
// Returns errors.TEMPLATE if the inputted interface failed to be cast.
//
// Example: {{ "(2*2)" | safeJS }}
// Returns: `(2*2)`
func (ns *Namespace) JS(i interface{}) (template.JS, error) {
	const op = "Templates.JS"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe JS to string"), Operation: op, Err: err}
	}
	return template.JS(s), nil
}

// safeJSStr
//
// Returns the given string as a html/template JSStr content.
// Returns errors.TEMPLATE if the inputted interface failed to be cast.
func (ns *Namespace) JSStr(i interface{}) (template.JSStr, error) {
	const op = "Templates.JSStr"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe JSStr to string"), Operation: op, Err: err}
	}
	return template.JSStr(s), nil
}

// safeUrl
//
// Returns a given string as html/template URL content.
// Returns errors.TEMPLATE if the inputted interface failed to be cast.
//
// Example: {{ "https://verbiscms.com" | safeUrl }}
// Returns: `https://verbiscms.com`
func (ns *Namespace) Url(i interface{}) (template.URL, error) {
	const op = "Templates.Url"
	s, err := cast.ToStringE(i)
	if err != nil {
		return "", &errors.Error{Code: errors.TEMPLATE, Message: fmt.Sprintf("Unable to cast to safe URL to string"), Operation: op, Err: err}
	}
	return template.URL(s), nil
}
