package handler

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/http/handler/frontend"
	"github.com/ainsleyclark/verbis/api/http/handler/spa"
	"github.com/ainsleyclark/verbis/api/models"
)

// Handler defines all of handler funcs for the app.
type Handler struct {
	Auth       api.AuthHandler
	Cache      api.CacheHandler
	Categories api.CategoryHandler
	Media      api.MediaHandler
	Options    api.OptionsHandler
	Posts      api.PostHandler
	Site       api.SiteHandler
	User       api.UserHandler
	Forms      api.FormHandler
	Fields     api.FieldHandler
	Frontend   frontend.PublicHandler
	SEO        frontend.SEOHandler
	SPA        spa.SPAHandler
}

// Construct
func New(m *models.Store, config config.Configuration) *Handler {
	return &Handler{
		Auth:       api.NewAuth(m, config),
		Cache:      api.NewCache(),
		Categories: api.NewCategories(m, config),
		Fields:     api.NewFields(m, config),
		Forms:      api.NewForms(m, config),
		Media:      api.NewMedia(m, config),
		Options:    api.NewwOptions(m, config),
		Posts:      api.NewPosts(m, config),
		Site:       api.NewSite(m, config),
		User:       api.NewUser(m, config),
		SPA:        spa.NewSpa(config),
		Frontend:   frontend.NewPublic(m, config),
		SEO:        frontend.NewSEO(m, config),
	}
}
