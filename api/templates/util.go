package templates

import (
	"html/template"
)

// escape HTML
func (t *TemplateFunctions) escape(text string) template.HTML {
	return template.HTML(text)
}