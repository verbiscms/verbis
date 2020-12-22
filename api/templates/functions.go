package templates

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html/template"
	"strings"
)

type TemplateFunctions struct {
	gin     *gin.Context
	post    *domain.PostData
	fields  map[string]interface{}
	site    *domain.Site
	store   *models.Store
	options domain.Options
	token   string
}

type TypeOfPage struct {
	PageType string
	Data     interface{}
}

// NewFunctions - Construct
func NewFunctions(g *gin.Context, s *models.Store, p *domain.PostData) *TemplateFunctions {

	return &TemplateFunctions{
		gin:     g,
		post:    p,
		fields:  p.Fields,
		site:    s.Site.GetGlobalConfig(),
		store:   s,
		options: s.Options.GetStruct(),
	}
}

// Get all template functions
func (t *TemplateFunctions) GetFunctions() template.FuncMap {

	funcMap := template.FuncMap{
		// Env
		"env":       t.env,
		"expandEnv": t.expandEnv,
		// Header & Footer
		"verbisHead": t.header,
		"verbisFoot": t.footer,
		"metaTitle":  t.metaTitle,
		// Fields
		"field":    t.getField,
		"fields":   t.getFields,
		"hasField": t.hasField,
		"repeater": t.getRepeater,
		"flexible": t.getFlexible,
		"subfield": t.getSubField,
		// Auth
		"auth":  t.auth,
		"admin": t.admin,
		// Posts
		"post":           t.getPost,
		"posts":          t.getPosts,
		"paginationPage": t.getPaginationPage,
		// Media
		"media": t.getMedia,
		// Paths
		"assets":  t.assetsPath,
		"storage": t.storagePath,
		// Body
		"body": t.body,
		// Partials
		"partial": t.partial,
		// Dict
		"dict": t.dict,
		// Dates
		"date":           t.date,
		"dateInZone":     t.dateInZone,
		"ago":            t.ago,
		"htmlDate":       t.htmlDate,
		"htmlDateInZone": t.htmlDateInZone,
		"duration":       t.duration,
		// Helpers
		"fullUrl": t.getFullUrl,
		"escape":  t.escape,
	}

	return funcMap
}

// GetData - Returns all the necessary data for template usage.
func (t *TemplateFunctions) GetData() (map[string]interface{}, error) {

	theme, err := t.store.Site.GetThemeConfig()
	if err != nil {
		return nil, err
	}

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
			"Layout":       t.post.Layout,
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

	return data, nil
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

	theme, err := t.store.Site.GetThemeConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Could not get the theme config ", Operation: op, Err: err},
		}).Error()
	}

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
