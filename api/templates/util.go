package templates

import (
	"html/template"
)

// escape HTML
func (t *TemplateFunctions) escape(text string) template.HTML {
	return template.HTML(text)
}

// Get all fields for template
func (t *TemplateFunctions) getFullUrl() string {
	return t.gin.Request.Host + t.gin.Request.URL.String()
}
