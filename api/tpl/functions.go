package tpl

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/fields"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
)

type TemplateManager struct {
	gin          *gin.Context
	post         *domain.PostData
	site         domain.Site
	store        *models.Store
	options      domain.Options
	config       config.Configuration
	themeConfig  domain.ThemeConfig
	fieldService fields.FieldService
	token        string
}

type TypeOfPage struct {
	PageType string
	Data     interface{}
}

// Obtain all of the paths in to variables
// for use with testing
var (
	themePath   = paths.Theme()
	basePath    = paths.Base()
	adminPath   = paths.Admin()
	apiPath     = paths.Api()
	storagePath = paths.Storage()
	uploadsPath = paths.Uploads()
)

// NewManager - Construct
func NewManager(g *gin.Context, s *models.Store, p *domain.PostData, c config.Configuration) *TemplateManager {
	return &TemplateManager{
		gin:          g,
		post:         p,
		site:         s.Site.GetGlobalConfig(),
		store:        s,
		options:      s.Options.GetStruct(),
		themeConfig:  s.Site.GetThemeConfig(),
		fieldService: fields.NewService(s, *p),
		config:       c,
	}
}

// Get all template functions
func (t *TemplateManager) GetFunctions() template.FuncMap {

	funcMap := template.FuncMap{
		"test": t.dd,
		// Attributes
		"body": t.body,
		"lang": t.lang,
		// Auth
		"auth":  t.auth,
		"admin": t.admin,
		// Fields
		"field":       t.fieldService.GetField,
		"fieldObject": t.fieldService.GetFieldObject,
		"fields":      t.fieldService.GetFields,
		"layout":      t.fieldService.GetLayout,
		"layouts":     t.fieldService.GetLayouts,
		"repeater":    t.fieldService.GetRepeater,
		"flexible":    t.fieldService.GetFlexible,
		// Header & Footer
		"verbisHead": t.header,
		"verbisFoot": t.footer,
		"metaTitle":  t.metaTitle,
		// Media
		"media": t.getMedia,
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
		// URL
		"baseUrl": t.getBaseURL,
		"scheme":  t.getScheme,
		"host":    t.getHost,
		"fullUrl": t.getFullURL,
		"url":     t.getURL,
		"query":   t.getQueryParams,
	}

	return funcMap
}

// GetData - Returns all the necessary data for template usage.
func (t *TemplateManager) GetData() map[string]interface{} {

	theme := t.store.Site.GetThemeConfig()

	data := map[string]interface{}{
		"Type":  t.orderOfSearch(),
		"Site":  t.site,
		"Theme": theme.Theme,
		//"Token": csrf.GetToken(t.gin),
		"Post": t.post,
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

func (t *TemplateManager) orderOfSearch() TypeOfPage {
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
