package templates

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"net/http/httptest"
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
		"getField": t.getField,
		"getFields": t.getFields,
		"hasField": t.hasField,
		"getRepeater": t.getRepeater,
		"getFlexible": t.getFlexible,
		"getSubField": t.getSubField,
		// Posts
		"getPost": t.getPost,
		// Media
		"getMedia": t.getMedia,
		// Auth
		"auth": t.isAuth,
		"admin": t.isAdmin,
		// Paths
		"assets": t.assetsPath,
		"storage": t.storagePath,
		// Partials
		"partial": t.partial,
		// Helpers
		"fullUrl": t.getFullUrl,
		"escape": t.escape,
	}

	t.functions = funcMap

	return funcMap
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

func newTestSuite() *TemplateFunctions {
	g, _ := gin.CreateTestContext(httptest.NewRecorder())
	g.Request, _ = http.NewRequest("GET", "/get", nil)

	return NewFunctions(g, &models.Store{}, &domain.Post{})
}