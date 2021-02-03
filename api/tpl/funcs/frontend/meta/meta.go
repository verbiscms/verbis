package meta

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"

	//"github.com/ainsleyclark/verbis/api/tpl/templates"
	"html/template"
)

type TemplateMeta struct {
	Site          domain.Site
	Post          *domain.PostData
	Options       domain.Options
	FacebookImage string
	TwitterImage  string
	deps          *deps.Deps
}

const (
	// The path of the emmbedded files to execute.
	EmbeddedPath = "/api/tpl/embedded/"
)

func (tm *TemplateMeta) GetImage(id int) string {
	img, err := tm.deps.Store.Media.GetById(id)
	if err != nil {
		return ""
	}
	return img.Url
}

// Header
//
// Header obtains all of the site and post wide Code Injection
// as well as any meta information from the page.
//
// Example: {{ verbisHead }}
func (ns *Namespace) Header() template.HTML {
	const op = "Templates.Header"

	tm := &TemplateMeta{
		Site:    ns.deps.Site,
		Post:    ns.post,
		Options: ns.deps.Options,
		deps:    ns.deps,
	}

	head := ns.executeTemplates(tm, []string{"meta.cms", "opengraph.cms", "twitter.cms"})

	return template.HTML(head)
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
	tm := &TemplateMeta{
		Post:    ns.post,
		Options: ns.deps.Options,
	}

	foot := ns.executeTemplates(tm, []string{"footer.cms"})

	return template.HTML(foot)
}

func (ns *Namespace) executeTemplates(tm *TemplateMeta, tpls []string) string {
	//head := ""
	//for _, name := range tpls {

		//var b bytes.Buffer
		//err := ns.deps.Tpl.Prepare(templates.Config{
		//	Root:      ns.deps.Paths.Base + EmbeddedPath,
		//	Extension: ".html",
		//}).Execute(&b, name, tm)

		//if err != nil {
		//	fmt.Println(err)
		//}

		//head += fmt.Sprintf("%s\n", b.String())
	//}
	return "head"
}
