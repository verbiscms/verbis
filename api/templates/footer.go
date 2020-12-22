package templates

import (
	"bytes"
	"github.com/yosssi/gohtml"
	"html/template"
)

// getFooter obtains all of the site and post wide Code Injection
func (t *TemplateFunctions) footer() template.HTML {
	var b bytes.Buffer

	// Get Global Foot (Code Injection)
	if t.options.CodeInjectionFoot != "" {
		b.WriteString(t.options.CodeInjectionFoot)
	}

	// Get Code Injection for the Post
	if *t.post.CodeInjectionFoot != "" {
		b.WriteString(*t.post.CodeInjectionFoot)
	}

	return template.HTML(gohtml.Format(b.String()))
}
