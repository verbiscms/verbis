package meta

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/yosssi/gohtml"
	"html/template"
)

// Header
//
// Header obtains all of the site and post wide Code Injection
// as well as any meta information from the page.
//
// Example: {{ verbisHead }}
func (ns *Namespace) Header() template.HTML {
	var b bytes.Buffer

	// Get Code Injection from the Post
	if ns.post.CodeInjectionHead != nil {
		b.WriteString(*ns.post.CodeInjectionHead)
	}

	// Get Code Injection from the Options (globally)
	if ns.deps.Options.CodeInjectionHead != "" {
		b.WriteString(ns.deps.Options.CodeInjectionHead)
	}

	// Obtain SEO & set post public
	seo := domain.PostSeo{
		Public:         false,
		ExcludeSitemap: true,
		Canonical:      nil,
	}

	if ns.post.SeoMeta.Seo != nil {
		seo = *ns.post.SeoMeta.Seo
	}

	postPublic := true
	if !seo.Public {
		postPublic = false
	}

	// Check if the site is public or page is public
	if !ns.deps.Options.SeoPublic || !postPublic {
		b.WriteString(`<meta name="robots" content="noindex">`)
	}

	// Check if there are trailing slashes
	slash := ""
	if ns.deps.Options.SeoEnforceSlash && ns.post.Slug != "/" {
		slash = "/"
	}

	// Write the Canonical
	if seo.Canonical != nil && *seo.Canonical != "" {
		b.WriteString(fmt.Sprintf(`<link rel="canonical" href="%s%s" />`, *seo.Canonical, slash))
	} else {
		b.WriteString(fmt.Sprintf(`<link rel="canonical" href="%s%s" />`, ns.deps.Site.Url+ns.post.Slug, slash))
	}

	// Obtain Meta
	meta := ns.post.SeoMeta.Meta
	if meta != nil {

		if meta.Description != "" {
			ns.writeMeta(&b, meta.Description)
		} else {
			ns.writeMeta(&b, ns.deps.Options.MetaDescription)
		}

		if meta.Facebook.Title != "" || meta.Facebook.Description != "" {
			ns.writeFacebook(&b, meta.Facebook.Title, meta.Facebook.Title, meta.Facebook.ImageId)
		} else {
			ns.writeFacebook(&b, ns.deps.Options.MetaFacebookTitle, ns.deps.Options.MetaFacebookDescription, ns.deps.Options.MetaFacebookImageId)
		}

		if meta.Twitter.Title != "" || meta.Twitter.Description != "" {
			ns.writeTwitter(&b, meta.Twitter.Title, meta.Twitter.Description, meta.Twitter.ImageId)
		} else {
			ns.writeTwitter(&b, ns.deps.Options.MetaTwitterTitle, ns.deps.Options.MetaTwitterDescription, ns.deps.Options.MetaTwitterImageId)
		}

	} else {
		ns.writeMeta(&b, ns.deps.Options.MetaDescription)
		ns.writeFacebook(&b, ns.deps.Options.MetaFacebookTitle, ns.deps.Options.MetaFacebookDescription, ns.deps.Options.MetaFacebookImageId)
		ns.writeTwitter(&b, ns.deps.Options.MetaTwitterTitle, ns.deps.Options.MetaTwitterDescription, ns.deps.Options.MetaTwitterImageId)
	}

	return template.HTML(gohtml.Format(b.String()))
}

// writeMeta
//
// Writes to the given *bytes.Buffer with meta description
// and article published time if they are not nil.
func (ns *Namespace) writeMeta(bytes *bytes.Buffer, description string) {
	if description != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"description\" content=\"%s\">", description))
	}
	if ns.post.PublishedAt != nil {
		bytes.WriteString(fmt.Sprintf("<meta property=\"article:modified_time\" content=\"%s\" />", ns.post.PublishedAt))
	}
}

// writeFacebook
//
// Opengraph writing to the given *bytes.Bufffer, this function
// will write website, site name, locale from options, title,
// description & post image if there is one.
func (ns *Namespace) writeFacebook(bytes *bytes.Buffer, title string, description string, imageId int) {

	if title != "" || description != "" {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:type\" content=\"website\">"))
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:site_name\" content=\"%s\">", ns.deps.Options.SiteTitle))
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:locale\" content=\"%s\">", ns.deps.Options.GeneralLocale))
	}

	if title != "" {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:title\" content=\"%s\">", title))
	}

	if description != "" {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:description\" content=\"%s\">", description))
	}

	image, foundImage := ns.deps.Store.Media.GetById(imageId)
	if foundImage == nil {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:image\" content=\"%s\">", ns.deps.Options.SiteUrl+image.Url))
	}
}

// writeTwitter
//
// Twitter card writing to the given *bytes.Bufffer, this function
// will write the title, description & post image if there is
// one.
func (ns *Namespace) writeTwitter(bytes *bytes.Buffer, title string, description string, imageId int) {
	if title != "" || description != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:card\" content=\"summary\">"))
	}

	if title != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:title\" content=\"%s\">", title))
	}

	if description != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:description\" content=\"%s\">", title))
	}

	image, foundImage := ns.deps.Store.Media.GetById(imageId)
	if foundImage == nil {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:image\" content=\"%s\">", ns.deps.Options.SiteUrl+image.Url))
	}
}

// MetaTitle
//
// metaTitle obtains the meta title from the post, if there is no
// title set on the post, it will look for the global title, if
// none, return empty string.
//
// Example: <title>Verbis - {{ metaTitle }}</title>
func (ns *Namespace) MetaTitle() string {
	postMeta := ns.post.SeoMeta.Meta

	if postMeta == nil {
		return ""
	}

	if postMeta.Title != "" {
		return postMeta.Title
	}

	if ns.deps.Options.MetaTitle != "" {
		return ns.deps.Options.MetaTitle
	}

	return ""
}

// Footer
//
// Obtains all of the site and post wide Code Injection
// Returns formatted HTML template for use after the
// closing `</body>`.
//
// Example: {{ verbisFoot }}
func (ns *Namespace) Footer() template.HTML {
	var b bytes.Buffer

	// Get Global Foot (Code Injection)
	if ns.deps.Options.CodeInjectionFoot != "" {
		b.WriteString(ns.deps.Options.CodeInjectionFoot)
	}

	// Get Code Injection for the Post
	if ns.post.CodeInjectionFoot != nil {
		b.WriteString(*ns.post.CodeInjectionFoot)
	}

	return template.HTML(gohtml.Format(b.String()))
}
