package templates

import (
	"bytes"
	"github.com/yosssi/gohtml"
	"html/template"
)

// getFooter obtains all of the site and post wide Code Injection
func (t *TemplateFunctions) getFooter() template.HTML {
	var b bytes.Buffer

	// Get Global Foot (Code Injection)
	if t.options.CodeInjectionFoot != "" {
		b.WriteString(t.options.CodeInjectionFoot)
	}

	// Get Code Injection for the Post
	if *t.post.CodeInjectFoot != "" {
		b.WriteString(*t.post.CodeInjectFoot)
	}

	return template.HTML(gohtml.Format(b.String()))
}