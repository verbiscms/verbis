package controllers

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// getSiteMock is a helper to obtain a mock site controller
// for testing.
func getSiteMock(m models.SiteRepository) *SiteController {
	return &SiteController{
		store: &models.Store{
			Site: m,
		},
	}
}

// Test_NewSite - Test construct
func Test_NewSite(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &SiteController{
		store:  &store,
		config: config,
	}
	got := newSite(&store, config)
	assert.Equal(t, got, want)
}

// TestSiteController_GetSite - Test GetSite route
func TestSiteController_GetSite(t *testing.T) {

	rr := newResponseRecorder(t)

	t.Run("Success", func(t *testing.T) {

		site := &domain.Site{
			Title:       "Verbis",
			Description: "VerbisCMS",
			Logo:        "/logo.svg",
			Url:         "verbiscms.com",
			Version:     "0.1",
		}
		siteMock := mocks.SiteRepository{}
		siteMock.On("GetGlobalConfig").Return(site)

		siteController := SiteController{
			store: &models.Store{
				Site: &siteMock,
			},
		}

		siteController.GetSite(rr.gin)

		want, err := json.Marshal(site)
		if err != nil {
			t.Fatal(err)
		}

		assert.JSONEq(t, string(want), rr.Data())
		assert.Equal(t,  200, rr.recorder.Code)
		assert.Equal(t, "Successfully obtained site config", rr.Message())
		assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
	})
}

// TestSiteController_GetTheme - Test GetTheme route
func TestSiteController_GetTheme(t *testing.T) {

	theme := domain.ThemeConfig{
		Theme: domain.Theme{
			Title:       "Verbis",
			Description: "VerbisCMS",
			Version:     "0.1",
		},
	}

	tt := []struct {
		name       string
		want       string
		status     int
		message    string
		mock func(u *mocks.SiteRepository)
	}{
		{
			name: "Success",
			want: `{"assets_path":"","editor":{"modules":null,"options":null},"resources":null,"theme":{"description":"VerbisCMS","title":"Verbis","version":"0.1"}}`,
			status: 200,
			message: "Successfully obtained theme config",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetThemeConfig").Return(theme, nil)
			},
		},
		{
			name:       "Internal Error",
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetThemeConfig").Return(domain.ThemeConfig{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for _, test := range tt {

		t.Run(test.name, func(t *testing.T) {
			rr := newResponseRecorder(t)
			mock := &mocks.SiteRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/theme", "/theme", nil, func(g *gin.Context) {
				getSiteMock(mock).GetTheme(g)
			})

			assert.JSONEq(t, test.want, rr.Data())
			assert.Equal(t, test.status, rr.recorder.Code)
			assert.Equal(t, test.message, rr.Message())
			assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
		})
	}
}

// TestSiteController_GetTemplates - Test GetTemplates route
func TestSiteController_GetTemplates(t *testing.T) {

	templates := domain.Templates{
		Template: []map[string]interface{}{
			{
				"test": "testing",
			},
		},
	}

	tt := []struct {
		name       string
		want       string
		status     int
		message    string
		mock func(u *mocks.SiteRepository)
	}{
		{
			name: "Success",
			want: `{"templates":[{"test":"testing"}]}`,
			status: 200,
			message: "Successfully obtained templates",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetTemplates").Return(templates, nil)
			},
		},
		{
			name:       "Not Found",
			want:       `{}`,
			status:     200,
			message:    "not found",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetTemplates").Return(domain.Templates{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		{
			name:       "Internal Error",
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetTemplates").Return(domain.Templates{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for _, test := range tt {

		t.Run(test.name, func(t *testing.T) {
			rr := newResponseRecorder(t)
			mock := &mocks.SiteRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/theme", "/theme", nil, func(g *gin.Context) {
				getSiteMock(mock).GetTemplates(g)
			})

			assert.JSONEq(t, test.want, rr.Data())
			assert.Equal(t, test.status, rr.recorder.Code)
			assert.Equal(t, test.message, rr.Message())
			assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
		})
	}
}

// TestSiteController_GetLayouts - Test GetLayouts route
func TestSiteController_GetLayouts(t *testing.T) {

	layouts := domain.Layouts{
		Layout: []map[string]interface{}{
			{
				"test": "testing",
			},
		},
	}

	tt := []struct {
		name       string
		want       string
		status     int
		message    string
		mock func(u *mocks.SiteRepository)
	}{
		{
			name: "Success",
			want: `{"layouts":[{"test":"testing"}]}`,
			status: 200,
			message: "Successfully obtained layouts",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetLayouts").Return(layouts, nil)
			},
		},
		{
			name:       "Not Found",
			want:       `{}`,
			status:     200,
			message:    "not found",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetLayouts").Return(domain.Layouts{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		{
			name:       "Internal Error",
			want:       `{}`,
			status:     500,
			message:    "internal",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetLayouts").Return(domain.Layouts{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for _, test := range tt {

		t.Run(test.name, func(t *testing.T) {
			rr := newResponseRecorder(t)
			mock := &mocks.SiteRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/theme", "/theme", nil, func(g *gin.Context) {
				getSiteMock(mock).GetLayouts(g)
			})

			assert.JSONEq(t, test.want, rr.Data())
			assert.Equal(t, test.status, rr.recorder.Code)
			assert.Equal(t, test.message, rr.Message())
			assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
		})
	}
}