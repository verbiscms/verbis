package templates

import (
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
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
	store *models.Store
	options domain.Options
}

type TypeOfPage struct {
	PageType string
	Data interface{}
}

// Construct
func NewFunctions(g *gin.Context, s *models.Store, p *domain.Post) *TemplateFunctions {
	const op = "Templates.NewFunctions"

	// Unmarshal the fields on the post
	f := make(map[string]interface{})
	if p.Fields != nil {
		if err := json.Unmarshal(*p.Fields, &f); err != nil {
			log.WithFields(log.Fields{
				"error": errors.Error{Code: errors.INTERNAL, Message: "Could not update the site logo", Operation: op, Err: err},
			}).Error()
		}
	}

	// Get the options struct
	options, err := s.Options.GetStruct()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: err},
		}).Fatal()
	}

	// Get the author associated with the post
	author, _ := s.User.GetById(p.UserId)

	// Get the categories associated with the post
	category, _ := s.Categories.GetByPost(p.Id)

	// New TemplateFunctions
	return &TemplateFunctions{
		gin: g,
		post: p,
		author: &author,
		category: category,
		fields: f,
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

	 fmt.Println(t.orderOfSearch())

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

// Get the app env
func (t *TemplateFunctions) appEnv() string {
	return environment.GetAppEnv()
}

// If the app is in production or development
func (t *TemplateFunctions) isProduction() bool {
	return environment.IsProduction()
}

// If the app is in debug mode
func (t *TemplateFunctions) isDebug() bool {
	return environment.IsDebug()
}


func (t *TemplateFunctions) orderOfSearch() TypeOfPage {
	const op = "Templates.orderOfSearch"

	data := TypeOfPage{
		PageType: "page",
		Data: nil,
	}

	if t.post.Resource == nil {
		return data
	}

	slug := t.post.Slug
	slugArr := strings.Split(slug, "/")
	last := slugArr[len(slugArr) - 1]


	if t.store.Categories.ExistsBySlug(last) {

		cat, err := t.store.Categories.GetBySlug(slug)
		if err != nil {
			return data
		}

		parentCat, err := t.store.Categories.GetById(cat.Id)
		if err != nil {
			data.PageType = "category_child_archive"
			data.Data = cat
			return data
		}

		data.PageType = "category_archive"
		data.Data = parentCat

	} else {
		data.PageType = "archive"
		data.Data = t.post.Resource
	}

	return data
}


