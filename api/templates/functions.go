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
)

type TemplateFunctions struct {
	gin *gin.Context
	post *domain.Post
	fields map[string]interface{}
	store *models.Store
	options domain.Options
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

	options, err := s.Options.GetStruct()
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: fmt.Errorf("could not get the options struct")},
		}).Fatal()
	}

	return &TemplateFunctions{
		gin: g,
		post: p,
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
		"isAuth": t.isAuth,
		"isAdmin": t.isAdmin,
		// Posts
		"getPost": t.getPost,
		// Media
		"getMedia": t.getMedia,
		// Paths
		"assets": t.assetsPath,
		"storage": t.storagePath,
		// Partials
		"partial": t.partial,
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
	 	"Site": t.store.Site.GetGlobalConfig(),
	 	"Post": t.post,
	 	"Theme": theme.Theme,
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

