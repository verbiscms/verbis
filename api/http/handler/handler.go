package handler

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/http/handler/frontend"
	"github.com/ainsleyclark/verbis/api/http/handler/spa"
	"github.com/ainsleyclark/verbis/api/models"
)

type Handler struct {
	Auth       api.AuthHandler
	Cache      api.CacheHandler
	Categories api.CategoryHandler
	Fields     api.FieldHandler
	Frontend   frontend.PublicHandler
	Media      api.MediaHandler
	Options    api.OptionsHandler
	Posts      api.PostHandler
	SPA        spa.SPAHandler
	SEO        frontend.SEOHandler
	Site       api.SiteHandler
	User       api.UserHandler
}

// Construct
func New(m *models.Store, config config.Configuration) (*Handler, error) {

	c := Handler{
		Auth:       api.NewAuth(m, config),
		Cache:      api.NewCache(),
		Categories: api.NewCategories(m, config),
		Fields:     api.NewFields(m, config),
		Frontend:   frontend.NewPublic(m, config),
		Media:      api.NewMedia(m, config),
		Options:    api.NewwOptions(m, config),
		Posts:      api.NewPosts(m, config),
		SPA:        spa.NewSpa(config),
		SEO:        frontend.NewSEO(m, config),
		Site:       api.NewSite(m, config),
		User:       api.NewUser(m, config),
	}

	return &c, nil
}
