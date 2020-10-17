package templates

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/yosssi/gohtml"
	"html/template"
)

// getHeader obtains all of the site and post wide Code Injection
// as well as any meta information from the page.
func (t *TemplateFunctions) getHeader() template.HTML {

	var b bytes.Buffer

	// Get Global Head (Code Injection)
	siteCodeHead, err := t.store.Options.GetByName("codeinjection_head")
	if siteCodeHead != "" && err == nil {
		b.WriteString(fmt.Sprintf("%v", siteCodeHead))
	}

	// Get Code Injection for the Post
	if *t.post.CodeInjectHead != "" {
		b.WriteString(*t.post.CodeInjectHead)
	}

	// Obtain Meta
	var meta domain.PostMea
	err = json.Unmarshal(*t.post.SeoMeta.Meta, &meta)
	if err != nil {
		fmt.Println(err)
	}

	// Get site options
	siteTitle, _ := t.store.Options.GetByName("site_title")

	// Normal Meta
	if meta.Title != "" {
		//TODO: Ask Kirk!
	//	b.WriteString(fmt.Sprintf("<meta name=\"description\" content=\"%s\">", meta.Description))
	}
	if meta.Description != "" {
		b.WriteString(fmt.Sprintf("<meta name=\"description\" content=\"%s\">", meta.Description))
	}

	// Open Graph
	if meta.Facebook.Title != "" || meta.Facebook.Description != "" || meta.Facebook.Image != "" {
		b.WriteString(fmt.Sprintf("<meta property=\"og:type\" content=\"website\">"))
		b.WriteString(fmt.Sprintf("<meta property=\"og:site_name\" content=\"%s\">", siteTitle))
	}
	if meta.Facebook.Title != "" {
		b.WriteString(fmt.Sprintf("<meta property=\"og:title\" content=\"%s\">", meta.Facebook.Title))
	}
	if meta.Facebook.Description != "" {
		b.WriteString(fmt.Sprintf("<meta property=\"og:description\" content=\"%s\">", meta.Facebook.Description))
	}

	// Facebook
	if meta.Facebook.Title != "" || meta.Facebook.Description != "" || meta.Facebook.Image != "" {
		b.WriteString(fmt.Sprintf("<meta property=\"og:type\" content=\"website\">"))
		b.WriteString(fmt.Sprintf("<meta property=\"og:site_name\" content=\"%s\">", siteTitle))
	}
	if meta.Facebook.Title != "" {
		b.WriteString(fmt.Sprintf("<meta property=\"og:title\" content=\"%s\">", meta.Facebook.Title))
	}
	if meta.Facebook.Description != "" {
		b.WriteString(fmt.Sprintf("<meta property=\"og:description\" content=\"%s\">", meta.Facebook.Description))
	}
	// TODO: Add image

	// Twitter
	if meta.Twitter.Title != "" || meta.Twitter.Description != "" || meta.Twitter.Image != "" {
		b.WriteString(fmt.Sprintf("<meta name=\"twitter:card\" content=\"summary\">"))
	}
	if meta.Twitter.Title != "" {
		b.WriteString(fmt.Sprintf("<meta name=\"twitter:title\" content=\"%s\">", meta.Twitter.Title))
	}
	if meta.Twitter.Description != "" {
		b.WriteString(fmt.Sprintf("<meta name=\"twitter:description\" content=\"%s\">", meta.Twitter.Title))
	}
	// TODO: Add image

	return template.HTML(gohtml.Format(b.String()))
}