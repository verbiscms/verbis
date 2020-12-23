package templates

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
)

type TemplateFunctions struct {
	gin         *gin.Context
	post        *domain.PostData
	fields      map[string]interface{}
	site        *domain.Site
	store       *models.Store
	options     domain.Options
	config      config.Configuration
	themeConfig domain.ThemeConfig
	token       string
}

type TypeOfPage struct {
	PageType string
	Data     interface{}
}

// NewFunctions - Construct
func NewFunctions(g *gin.Context, s *models.Store, p *domain.PostData, c config.Configuration) *TemplateFunctions {

	return &TemplateFunctions{
		gin:         g,
		post:        p,
		fields:      p.Fields,
		site:        s.Site.GetGlobalConfig(),
		store:       s,
		options:     s.Options.GetStruct(),
		themeConfig: s.Site.GetThemeConfig(),
		config:      c,
	}
}

// Get all template functions
func (t *TemplateFunctions) GetFunctions() template.FuncMap {

	funcMap := template.FuncMap{
		// Auth
		"auth":  t.auth,
		"admin": t.admin,
		// Body
		"body": t.body,
		// Categories
		// Date & Time
		"date":           t.date,
		"dateInZone":     t.dateInZone,
		"ago":            t.ago,
		"htmlDate":       t.htmlDate,
		"htmlDateInZone": t.htmlDateInZone,
		"duration":       t.duration,
		// Dict
		"dict": t.dict,
		// Fields
		"field":    t.getField,
		"fields":   t.getFields,
		"hasField": t.hasField,
		"repeater": t.getRepeater,
		"flexible": t.getFlexible,
		"subfield": t.getSubField,
		// Header & Footer
		"verbisHead": t.header,
		"verbisFoot": t.footer,
		"metaTitle":  t.metaTitle,
		// Media
		"media": t.getMedia,
		// OS
		"env":       t.env,
		"expandEnv": t.expandEnv,
		// Partials
		"partial": t.partial,
		// Posts
		"post":           t.getPost,
		"posts":          t.getPosts,
		"paginationPage": t.getPaginationPage,
		// Paths
		"basePath":      t.basePath,
		"adminPath":     t.adminPath,
		"apiPath":       t.apiPath,
		"themePath":     t.themePath,
		"uploadsPath":   t.uploadsPath,
		"assetsPath":    t.assetsPath,
		"storagePath":   t.storagePath,
		"templatesPath": t.templatesPath,
		"layoutsPath":   t.layoutsPath,
		// Safe
		"safeHTML":     t.safeHTML,
		"safeHTMLAttr": t.safeHTMLAttr,
		"safeCSS":      t.safeCSS,
		"safeJS":       t.safeJS,
		"safeJSStr":    t.safeJSStr,
		"safeURL":      t.safeUrl,
		// Helpers
		//"fullUrl": t.getFullUrl,
	}

	return funcMap
}

// GetData - Returns all the necessary data for template usage.
func (t *TemplateFunctions) GetData() map[string]interface{} {

	theme := t.store.Site.GetThemeConfig()

	data := map[string]interface{}{
		"Type":  t.orderOfSearch(),
		"Site":  t.store.Site.GetGlobalConfig(),
		"Theme": theme.Theme,
		//"Token": csrf.GetToken(t.gin),
		"Post": map[string]interface{}{
			"Id":           t.post.Id,
			"UUID":         t.post.UUID,
			"Slug":         t.post.Slug,
			"Title":        t.post.Title,
			"Status":       t.post.Status,
			"Resource":     t.post.Resource,
			"PageTemplate": t.post.PageTemplate,
			"PageLayout":   t.post.PageLayout,
			"PublishedAt":  t.post.PublishedAt,
			"UpdatedAt":    t.post.UpdatedAt,
			"CreatedAt":    t.post.CreatedAt,
			"Author":       t.post.Author,
			"Category":     t.post.Category,
		},
		"Options": map[string]interface{}{
			"Social": map[string]interface{}{
				"Facebook":  t.options.SocialFacebook,
				"Twitter":   t.options.SocialTwitter,
				"Youtube":   t.options.SocialYoutube,
				"LinkedIn":  t.options.SocialLinkedIn,
				"Instagram": t.options.SocialInstagram,
				"Pintrest":  t.options.SocialPinterest,
			},
			"Contact": map[string]interface{}{
				"Email":     t.options.ContactEmail,
				"Telephone": t.options.ContactTelephone,
				"Address":   t.options.ContactAddress,
			},
		},
	}

	return data
}

func (t *TemplateFunctions) orderOfSearch() TypeOfPage {
	const op = "Templates.orderOfSearch"

	data := TypeOfPage{
		PageType: "page",
		Data:     nil,
	}

	slug := t.post.Slug
	slugArr := strings.Split(slug, "/")
	last := slugArr[len(slugArr)-1]

	theme := t.store.Site.GetThemeConfig()

	if _, ok := theme.Resources[last]; ok {
		data.PageType = "archive"
		data.Data = t.post.Post.Resource
		return data
	}

	if t.store.Categories.ExistsBySlug(last) {

		cat, err := t.store.Categories.GetBySlug(last)
		if err != nil {
			return data
		}

		parentCat, err := t.store.Categories.GetById(cat.Id)
		if err != nil {
			data.PageType = "category_child_archive"
			data.Data = cat
			return data
		} else {
			data.PageType = "category_archive"
			data.Data = parentCat
			return data
		}
	}

	return data
}
