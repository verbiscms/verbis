// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

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
func getSiteMock(m models.SiteRepository) *Site {
	return &Site{
		store: &models.Store{
			Site: m,
		},
	}
}

// Test_NewSite - Test construct
func Test_NewSite(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &Site{
		store:  &store,
		config: config,
	}
	got := NewSite(&store, config)
	assert.Equal(t, got, want)
}

// TestSite_GetSite - Test GetSite route
func TestSite_GetSite(t *testing.T) {

	rr := newTestSuite(t)

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

		Site := Site{
			store: &models.Store{
				Site: &siteMock,
			},
		}

		Site.GetSite(rr.gin)

		want, err := json.Marshal(site)
		if err != nil {
			t.Fatal(err)
		}

		rr.Run(string(want), 200, "Successfully obtained site config")
	})
}

// TestSite_GetTheme - Test GetTheme route
func TestSite_GetTheme(t *testing.T) {

	theme := domain.ThemeConfig{
		Theme: domain.Theme{
			Title:       "Verbis",
			Description: "VerbisCMS",
			Version:     "0.1",
		},
	}

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(u *mocks.SiteRepository)
	}{
		"Success": {
			want:    `{"assets_path":"","editor":{"modules":null,"options":null},"file_extension":"","layout_dir":"","resources":null,"template_dir":"","theme":{"description":"VerbisCMS","title":"Verbis","version":"0.1"}}`,
			status:  200,
			message: "Successfully obtained theme config",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetThemeConfig").Return(theme, nil)
			},
		},
		//"Internal Error": {
		//	want:    `{}`,
		//	status:  500,
		//	message: "internal",
		//	mock: func(m *mocks.SiteRepository) {
		//		m.On("GetThemeConfig").Return(domain.ThemeConfig{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
		//	},
		//},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.SiteRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/theme", "/theme", nil, func(g *gin.Context) {
				getSiteMock(mock).GetTheme(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestSite_GetTemplates - Test GetTemplates route
func TestSite_GetTemplates(t *testing.T) {

	templates := domain.Templates{
		Template: []map[string]interface{}{
			{
				"test": "testing",
			},
		},
	}

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(u *mocks.SiteRepository)
	}{
		"Success": {
			want:    `{"templates":[{"test":"testing"}]}`,
			status:  200,
			message: "Successfully obtained templates",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetTemplates").Return(templates, nil)
			},
		},
		"Not Found": {
			want:    `{}`,
			status:  200,
			message: "not found",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetTemplates").Return(domain.Templates{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetTemplates").Return(domain.Templates{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.SiteRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/theme", "/theme", nil, func(g *gin.Context) {
				getSiteMock(mock).GetTemplates(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}

// TestSite_GetLayouts - Test GetLayouts route
func TestSite_GetLayouts(t *testing.T) {

	layouts := domain.Layouts{
		Layout: []map[string]interface{}{
			{
				"test": "testing",
			},
		},
	}

	tt := map[string]struct {
		want    string
		status  int
		message string
		mock    func(u *mocks.SiteRepository)
	}{
		"Success": {
			want:    `{"layouts":[{"test":"testing"}]}`,
			status:  200,
			message: "Successfully obtained layouts",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetLayouts").Return(layouts, nil)
			},
		},
		"Not Found": {
			want:    `{}`,
			status:  200,
			message: "not found",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetLayouts").Return(domain.Layouts{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			want:    `{}`,
			status:  500,
			message: "internal",
			mock: func(m *mocks.SiteRepository) {
				m.On("GetLayouts").Return(domain.Layouts{}, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.SiteRepository{}
			test.mock(mock)

			rr.RequestAndServe("GET", "/theme", "/theme", nil, func(g *gin.Context) {
				getSiteMock(mock).GetLayouts(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}
