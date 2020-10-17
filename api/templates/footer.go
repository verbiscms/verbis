package templates

import (
	"bytes"
	"fmt"
	"github.com/yosssi/gohtml"
	"html/template"
)

// getFooter obtains all of the site and post wide Code Injection
func (t *TemplateFunctions) getFooter() template.HTML {

	var b bytes.Buffer

	// Get Global Foot (Code Injection)
	siteCodeFoot, err := t.store.Options.GetByName("codeinjection_foot")
	if siteCodeFoot != "" && err == nil {
		b.WriteString(fmt.Sprintf("%v", siteCodeFoot))
	}

	// Get Code Injection for the Post
	if *t.post.CodeInjectFoot != "" {
		b.WriteString(*t.post.CodeInjectFoot)
	}

	return template.HTML(gohtml.Format(b.String()))
}