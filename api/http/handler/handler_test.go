package handler

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/http/handler/api"
	"github.com/ainsleyclark/verbis/api/http/handler/frontend"
	"github.com/ainsleyclark/verbis/api/http/handler/spa"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_New - Test construct
func Test_New(t *testing.T) {
	mockSite := mocks.SiteRepository{}
	mockSite.On("GetThemeConfig").Return(domain.ThemeConfig{}, nil)
	mockOptions := mocks.OptionsRepository{}
	mockOptions.On("GetStruct").Return(domain.Options{})
	m := &models.Store{
		Options: &mockOptions,
		Site:    &mockSite,
	}
	c := config.Configuration{}
	want := &Handler{
		Auth:       api.NewAuth(m, c),
		Cache:      api.NewCache(),
		Categories: api.NewCategories(m, c),
		Fields:     api.NewFields(m, c),
		Media:      api.NewMedia(m, c),
		Options:    api.NewwOptions(m, c),
		Posts:      api.NewPosts(m, c),
		Site:       api.NewSite(m, c),
		User:       api.NewUser(m, c),
		SPA:        spa.NewSpa(c),
		Frontend:   frontend.NewPublic(m, c),
		SEO:        frontend.NewSEO(m, c),
	}
	got := New(m, c)
	assert.ObjectsAreEqual(want, got)
}
