package templates

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html/template"
	"strings"
)

type TemplateFunctions struct {
	gin *gin.Context
	post *domain.Post
	author *domain.User
	category *domain.Category
	fields map[string]interface{}
	site *domain.Site
	store *models.Store
	options domain.Options
}

type TypeOfPage struct {
	PageType string
	Data interface{}
}

// NewFunctions - Construct
func NewFunctions(g *gin.Context, s *models.Store, p *domain.Post) *TemplateFunctions {
	const op = "Templates.NewFunctions"

	// TODO - This needs to be in the posts
	fields := make(map[string]interface{})
	if p.Fields != nil {
		if err := json.Unmarshal(*p.Fields, &fields); err != nil {
			log.WithFields(log.Fields{
				"error": errors.Error{Code: errors.INTERNAL, Message: "Could not unmarshal the post fields", Operation: op, Err: err},
			}).Error()
		}
	}

	options, err := s.Options.GetStruct()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: err},
		}).Fatal()
	}

	author, _ := s.User.GetById(p.UserId)

	category, _ := s.Categories.GetByPost(p.Id)

	site := s.Site.GetGlobalConfig()

	return &TemplateFunctions{
		gin: g,
		post: p,
		author: &author,
		category: category,
		fields: fields,
		site: site,
		store: s,
		options: options,
	}
}

// Get all template functions
func (t *TemplateFunctions) GetFunctions() template.FuncMap {

	funcMap := template.FuncMap{
		// Env
		//"appEnv": t.appEnv,
		"isProduction": t.isProduction,
		"isDebug": t.isDebug,
		// Header & Footer
		"verbisHead": t.getHeader,
		"verbisFoot": t.getFooter,
		"getMetaTitle": t.getMetaTitle,
		// Fields
		"getField": t.getField,
		"getFields": t.getFields,
		"hasField": t.hasField,
		"getRepeater": t.getRepeater,
		"getFlexible": t.getFlexible,
		"getSubField": t.getSubField,
		// Auth
		"isAuth": t.isAuth,
		"isAdmin": t.isAdmin,
		// Posts
		"getPost": t.getPost,
		"getPosts": t.getPosts,
		"getPaginationPage": t.getPaginationPage,
		// Media
		"getMedia": t.getMedia,
		// Paths
		"assets": t.assetsPath,
		"storage": t.storagePath,
		// Partials
		"partial": t.partial,
		// Dict
		"dict": t.dict,
		// Helpers
		"fullUrl": t.getFullUrl,
		"escape": t.escape,
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
	 	"Type": t.orderOfSearch(),
		"Site": t.store.Site.GetGlobalConfig(),
		"Theme": theme.Theme,
		"Post": map[string]interface{}{
			"Id": t.post.Id,
			"UUID": t.post.UUID,
			"Slug": t.post.Slug,
			"Title": t.post.Title,
			"Status": t.post.Status,
			"Resource": t.post.Resource,
			"PageTemplate": t.post.PageTemplate,
			"Layout": t.post.Layout,
			"PublishedAt": t.post.PublishedAt,
			"UpdatedAt": t.post.UpdatedAt,
			"CreatedAt": t.post.CreatedAt,
			"Author": t.author,
			"Category": t.category,
		},
		"Options": map[string]interface{}{
			"Social": map[string]interface{}{
				"Facebook": t.options.SocialFacebook,
				"Twitter": t.options.SocialTwitter,
				"Youtube": t.options.SocialYoutube,
				"LinkedIn": t.options.SocialLinkedIn,
				"Instagram": t.options.SocialInstagram,
				"Pintrest": t.options.SocialPinterest,
			},
			"Contact": map[string]interface{}{
				"Email": t.options.ContactEmail,
				"Telephone": t.options.ContactTelephone,
				"Address": t.options.ContactAddress,
			},
		},
	 }

	return data, nil
}

func (t *TemplateFunctions) orderOfSearch() TypeOfPage {
	const op = "Templates.orderOfSearch"

	data := TypeOfPage{
		PageType: "page",
		Data: nil,
	}

	slug := t.post.Slug
	slugArr := strings.Split(slug, "/")
	last := slugArr[len(slugArr) - 1]

	theme, err := t.store.Site.GetThemeConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Could not get the theme config ", Operation: op, Err: err},
		}).Error()
	}

	if _, ok := theme.Resources[last]; ok {
		data.PageType = "archive"
		data.Data = t.post.Resource
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


