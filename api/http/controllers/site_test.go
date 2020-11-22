package controllers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"testing"
)

func TestSiteController_GetSite(t *testing.T) {

	test := newResponseRecorder(t)

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

		siteController.GetSite(test.gin)

		test.runSuccess(site)
	})
}

func TestSiteController_GetTheme(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		test := newResponseRecorder(t)

		theme := domain.ThemeConfig{
			Theme: domain.Theme{
				Title:       "Verbis",
				Description: "VerbisCMS",
				Version:     "0.1",
			},
		}
		siteMock := &mocks.SiteRepository{}
		siteMock.On("GetThemeConfig").Return(theme, nil)

		siteController := SiteController{
			store: &models.Store{
				Site: siteMock,
			},
		}

		siteController.GetTheme(test.gin)

		test.runSuccess(theme)
	})

	t.Run("Error", func(t *testing.T) {

		test := newResponseRecorder(t)

		siteMock := &mocks.SiteRepository{}
		siteMock.On("GetThemeConfig").Return(domain.ThemeConfig{}, fmt.Errorf("error"))
		siteController := SiteController{
			store: &models.Store{
				Site: siteMock,
			},
		}

		siteController.GetTheme(test.gin)

		test.runInternalError()
	})
}

func TestSiteController_GetTemplates(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		test := newResponseRecorder(t)

		templates := domain.Templates{
			Template: []map[string]interface{}{
				{
					"test": "testing",
				},
			},
		}
		siteMock := &mocks.SiteRepository{}
		siteMock.On("GetTemplates").Return(&templates, nil)

		siteController := SiteController{
			store: &models.Store{
				Site: siteMock,
			},
		}

		siteController.GetTemplates(test.gin)

		test.runSuccess(templates)
	})

	t.Run("Error", func(t *testing.T) {

		test := newResponseRecorder(t)

		siteMock := &mocks.SiteRepository{}
		siteMock.On("GetTemplates").Return(&domain.Templates{}, fmt.Errorf("error"))
		siteController := SiteController{
			store: &models.Store{
				Site: siteMock,
			},
		}

		siteController.GetTemplates(test.gin)

		test.runInternalError()
	})
}

func TestSiteController_GetLayouts(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		test := newResponseRecorder(t)

		layouts := domain.Layouts{
			Layout: []map[string]interface{}{
				{
					"test": "testing",
				},
			},
		}
		siteMock := &mocks.SiteRepository{}
		siteMock.On("GetLayouts").Return(&layouts, nil)

		siteController := SiteController{
			store: &models.Store{
				Site: siteMock,
			},
		}

		siteController.GetLayouts(test.gin)

		test.runSuccess(layouts)
	})

	t.Run("Error", func(t *testing.T) {

		test := newResponseRecorder(t)

		siteMock := &mocks.SiteRepository{}
		siteMock.On("GetLayouts").Return(&domain.Layouts{}, fmt.Errorf("error"))
		siteController := SiteController{
			store: &models.Store{
				Site: siteMock,
			},
		}

		siteController.GetLayouts(test.gin)

		test.runInternalError()
	})
}
