package templates

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/files"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html/template"
	"strings"
)

type TemplateFunctions struct {
	gin *gin.Context
	post *domain.Post
	fields map[string]interface{}
	store *models.Store
	functions template.FuncMap
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

	fields := newFields(t.fields)

	funcMap := template.FuncMap{
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
		"getField": fields.getField,
		"getFields": fields.getFields,
		"hasField": fields.hasField,
		"getRepeater": fields.getRepeater,
		"getFlexible": fields.getFlexible,
		"getSubField": fields.getSubField,
		// Auth
		"auth": t.isAuth,
		"admin": t.isAdmin,
		// Paths
		"assets": t.assetsPath,
		"storage": t.storagePath,
		// Helpers
		"fullUrl": t.GetFullUrl,
		"escape": t.escape,
		"partial": t.partial,
	}

	t.functions = funcMap

	return funcMap
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

// partial
func (t *TemplateFunctions) partial(name string, data ...interface{}) template.HTML {
	path := paths.Theme() + "/" + name

	fmt.Println(data)

	if !files.Exists(path) {
		panic(fmt.Errorf("No file exists with the path: %s", name))
	}

	pathArr := strings.Split(path, "/")
	file, err := template.New(pathArr[len(pathArr) - 1]).Funcs(t.functions).ParseFiles(path)
	if err != nil {
		panic(fmt.Errorf("Unable to create a new partial file: %v", err))
	}

	var tpl bytes.Buffer
	err = file.Execute(&tpl, data)
	if err != nil {
		panic(err)
	}

	return template.HTML(tpl.String())
}