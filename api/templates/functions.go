package templates

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html/template"
)

type TemplateFunctions struct {
	gin *gin.Context
	post *domain.Post
	fields map[string]interface{}
	store *models.Store
}

// Construct
func NewFunctions(g *gin.Context, s *models.Store, p *domain.Post) *TemplateFunctions {

	f := make(map[string]interface{})
	if p.Fields != nil {
		if err := json.Unmarshal(*p.Fields, &f); err != nil {
			log.Error(err)
		}
	}

	tf := &TemplateFunctions{
		gin: g,
		post: p,
		fields: f,
		store: s,
	}

	return tf
}

// Get all template functions
func (t *TemplateFunctions) GetFunctions() template.FuncMap {
	return template.FuncMap{
		// Env
		//"appEnv": t.appEnv,
		"isProduction": t.isProduction,
		"isDebug": t.isDebug,
		// Header & Footer
		"verbis_head": t.getHeader,
		"verbis_foot": t.getFooter,
		// Posts
		"getResource": t.getResource,
		// Fields
		"getField": t.getField,
		"getFields": t.getFields,
		"hasField": t.hasField,
		"getRepeater": t.getRepeater,
		"getFlexible": t.getFlexible,
		"getSubField": t.getSubField,
		// Auth
		"auth": t.isAuth,
		"admin": t.isAdmin,
		// Paths
		"assets": t.assetsPath,
		"storage": t.storagePath,
		// Helpers
		"fullUrl": t.GetFullUrl,
		"escape": t.escape,
	}
}

/*
 * Environment
 * Functions for templates for the environment
 */

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

/*
 * Posts
 * Functions for templates for the Posts and modifications
*/

// Get the post resource
func (t *TemplateFunctions) getResource() string {
	resource := t.post.Resource
	if resource == nil {
		return ""
	}
	return *resource
}

func (t *TemplateFunctions) isResource(resource string) bool {

	return false
}

// Get the post resource
func (t *TemplateFunctions) getLayout() string {
	resource := t.post.Resource
	if resource == nil {
		return ""
	}
	return *resource
}

func (t *TemplateFunctions) isLayout(resource string) bool {

	return false
}

// Get the post template
func (t *TemplateFunctions) getTemplate() string {
	resource := t.post.Resource
	if resource == nil {
		return ""
	}
	return *resource
}

func (t *TemplateFunctions) isTemplate(resource string) bool {

	return false
}

// Get author
// Is draft
// Is published
func (t *TemplateFunctions) getResources(query map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{}
}


/*
 * Auth
 * Functions for templates for anything else
*/

// If the user is authenticated
func (t *TemplateFunctions) isAuth() bool {
	cookie, err := t.gin.Cookie("verbis-session")
	if err != nil {
		return false
	}

	_, err = t.store.Session.GetByKey(cookie)
	if err != nil {
		return false
	}

	return true
}

// If the user is admin
func (t *TemplateFunctions) isAdmin() bool {
	cookie, err := t.gin.Cookie("verbis-session")
	if err != nil {
		return false
	}

	us, err := t.store.Session.GetByKey(cookie)
	if err != nil {
		return false
	}

	_, err = t.store.User.GetById(us.UserId)
	if err != nil {
		return false
	}

	//if u.AccessLevel != 2 {
	//	return false
	//}

	return true
}

/*
 * Paths
 * Functions for templates for paths
 */

// Retrieve the assets path for the theme
func (t *TemplateFunctions) assetsPath() string {
	//return config.Theme.AssetsPath
	return ""
}

// Retrieve the assets path for the theme
func (t *TemplateFunctions) storagePath() string {
	// TODO: Make dynamic?
	return "/storage"
}

/*
 * Helpers
 * Functions for templates for miscellaneous
*/

// Get all fields for template
func (t *TemplateFunctions) GetFullUrl() string {
	return t.gin.Request.Host + t.gin.Request.URL.String()
}

// escape HTML
func (t *TemplateFunctions) escape(text string) template.HTML {
	return template.HTML(text)
}