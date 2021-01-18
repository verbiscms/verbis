package tpl

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/yosssi/gohtml"
	"html/template"
)

// header
//
// Header obtains all of the site and post wide Code Injection
// as well as any meta information from the page.
//
// Example: {{ verbisHead }}
func (t *TemplateManager) header() template.HTML {
	var b bytes.Buffer

	// Get Code Injection from the Post
	if t.post.CodeInjectionHead != nil {
		b.WriteString(*t.post.CodeInjectionHead)
	}

	// Get Code Injection from the Options (globally)
	if t.options.CodeInjectionHead != "" {
		b.WriteString(t.options.CodeInjectionHead)
	}

	// Obtain SEO & set post public
	seo := domain.PostSeo{
		Public:         false,
		ExcludeSitemap: true,
		Canonical:      nil,
	}

	if t.post.SeoMeta.Seo != nil {
		seo = *t.post.SeoMeta.Seo
	}

	postPublic := true
	if !seo.Public {
		postPublic = false
	}

	// Check if the site is public or page is public
	if !t.options.SeoPublic || !postPublic {
		b.WriteString(`<meta name="robots" content="noindex">`)
	}

	// Write the Canonical
	if seo.Canonical != nil && *seo.Canonical != "" {
		b.WriteString(fmt.Sprintf(`<link rel="canonical" href="%s" />`, *seo.Canonical))
	} else {
		b.WriteString(fmt.Sprintf(`<link rel="canonical" href="%s" />`, t.site.Url+t.post.Slug))
	}

	// Obtain Meta
	meta := t.post.SeoMeta.Meta
	if meta != nil {

		if meta.Description != "" {
			t.writeMeta(&b, meta.Description)
		} else {
			t.writeMeta(&b, t.options.MetaDescription)
		}

		if meta.Facebook.Title != "" || meta.Facebook.Description != "" {
			t.writeFacebook(&b, meta.Facebook.Title, meta.Facebook.Title, meta.Facebook.ImageId)
		} else {
			t.writeFacebook(&b, t.options.MetaFacebookTitle, t.options.MetaFacebookDescription, t.options.MetaFacebookImageId)
		}

		if meta.Twitter.Title != "" || meta.Twitter.Description != "" {
			t.writeTwitter(&b, meta.Twitter.Title, meta.Twitter.Description, meta.Twitter.ImageId)
		} else {
			t.writeTwitter(&b, t.options.MetaTwitterTitle, t.options.MetaTwitterDescription, t.options.MetaTwitterImageId)
		}

	} else {
		t.writeMeta(&b, t.options.MetaDescription)
		t.writeFacebook(&b, t.options.MetaFacebookTitle, t.options.MetaFacebookDescription, t.options.MetaFacebookImageId)
		t.writeTwitter(&b, t.options.MetaTwitterTitle, t.options.MetaTwitterDescription, t.options.MetaTwitterImageId)
	}

	return template.HTML(gohtml.Format(b.String()))
}

// writeMeta
//
// Writes to the given *bytes.Buffer with meta description
// and article published time if they are not nil.
func (t *TemplateManager) writeMeta(bytes *bytes.Buffer, description string) {
	if description != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"description\" content=\"%s\">", description))
	}
	if t.post.PublishedAt != nil {
		bytes.WriteString(fmt.Sprintf("<meta property=\"article:modified_time\" content=\"%s\" />", t.post.PublishedAt))
	}
}

// writeFacebook
//
// Opengraph writing to the given *bytes.Bufffer, this function
// will write website, site name, locale from options, title,
// description & post image if there is one.
func (t *TemplateManager) writeFacebook(bytes *bytes.Buffer, title string, description string, imageId int) {

	if title != "" || description != "" {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:type\" content=\"website\">"))
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:site_name\" content=\"%s\">", t.options.SiteTitle))
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:locale\" content=\"%s\">", t.options.GeneralLocale))
	}

	if title != "" {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:title\" content=\"%s\">", title))
	}

	if description != "" {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:description\" content=\"%s\">", description))
	}

	image, foundImage := t.store.Media.GetById(imageId)
	if foundImage == nil {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:image\" content=\"%s\">", t.options.SiteUrl+image.Url))
	}
}

// writeTwitter
//
// Twitter card writing to the given *bytes.Bufffer, this function
// will write the title, description & post image if there is
// one.
func (t *TemplateManager) writeTwitter(bytes *bytes.Buffer, title string, description string, imageId int) {
	if title != "" || description != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:card\" content=\"summary\">"))
	}

	if title != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:title\" content=\"%s\">", title))
	}

	if description != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:description\" content=\"%s\">", title))
	}

	image, foundImage := t.store.Media.GetById(imageId)
	if foundImage == nil {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:image\" content=\"%s\">", t.options.SiteUrl+image.Url))
	}
}

// metaTitle
//
// metaTitle obtains the meta title from the post, if there is no
// title set on the post, it will look for the global title, if
// none, return empty string.
//
// Example: <title>Verbis - {{ metaTitle }}</title>
func (t *TemplateManager) metaTitle() string {
	postMeta := t.post.SeoMeta.Meta

	if postMeta.Title != "" {
		return postMeta.Title
	}

	if t.options.MetaTitle != "" {
		return t.options.MetaTitle
	}

	return ""
}

// footer
//
// Obtains all of the site and post wide Code Injection
// Returns formatted HTML template for use after the
// closing `</body>`.
//
// Example: {{ verbisFoot }}
func (t *TemplateManager) footer() template.HTML {
	var b bytes.Buffer

	// Get Global Foot (Code Injection)
	if t.options.CodeInjectionFoot != "" {
		b.WriteString(t.options.CodeInjectionFoot)
	}

	// Get Code Injection for the Post
	if t.post.CodeInjectionFoot != nil {
		b.WriteString(*t.post.CodeInjectionFoot)
	}

	return template.HTML(gohtml.Format(b.String()))
}
