// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"fmt"
	app "github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/redirects"
	"github.com/ainsleyclark/verbis/api/store"
	"net/http"
)

const (
	// The default redirect path used for testing.
	RedirectPath = "/page"
)

func (t *MiddlewareTestSuite) Test_Redirects() {
	tt := map[string]struct {
		status      int
		url         string
		redirectURL string
		mock        func(m *mocks.Repository)
	}{
		"Admin Path": {
			http.StatusOK,
			"/admin",
			"",
			nil,
		},
		"API Path": {
			http.StatusOK,
			app.HTTPAPIRoute,
			"",
			nil,
		},
		"No Redirects": {
			http.StatusOK,
			RedirectPath,
			"",
			func(m *mocks.Repository) {
				m.On("FindByFrom", RedirectPath).Return(domain.Redirect{}, fmt.Errorf("err"))
			},
		},
		"300": {
			http.StatusMultipleChoices,
			"/page/test",
			RedirectPath,
			func(m *mocks.Repository) {
				m.On("FindByFrom", "/page/test").Return(domain.Redirect{
					To: RedirectPath, From: "/page/test", Code: 300,
				}, nil)
			},
		},
		"301": {
			http.StatusMovedPermanently,
			"/page/test",
			RedirectPath,
			func(m *mocks.Repository) {
				m.On("FindByFrom", "/page/test").Return(domain.Redirect{
					To: RedirectPath, From: "/page/test", Code: 301,
				}, nil)
			},
		},
		"302": {
			http.StatusFound,
			"/page/test",
			RedirectPath,
			func(m *mocks.Repository) {
				m.On("FindByFrom", "/page/test").Return(domain.Redirect{
					To: RedirectPath, From: "/page/test", Code: 302,
				}, nil)
			},
		},
		"303": {
			http.StatusSeeOther,
			"/page/test",
			RedirectPath,
			func(m *mocks.Repository) {
				m.On("FindByFrom", "/page/test").Return(domain.Redirect{
					To: RedirectPath, From: "/page/test", Code: 303,
				}, nil)
			},
		},
		"304": {
			http.StatusNotModified,
			"/page/test",
			RedirectPath,
			func(m *mocks.Repository) {
				m.On("FindByFrom", "/page/test").Return(domain.Redirect{
					To: RedirectPath, From: "/page/test", Code: 304,
				}, nil)
			},
		},
		"307": {
			http.StatusTemporaryRedirect,
			"/page/test",
			RedirectPath,
			func(m *mocks.Repository) {
				m.On("FindByFrom", "/page/test").Return(domain.Redirect{
					To: RedirectPath, From: "/page/test", Code: 307,
				}, nil)
			},
		},
		"308": {
			http.StatusPermanentRedirect,
			"/page/test",
			RedirectPath,
			func(m *mocks.Repository) {
				m.On("FindByFrom", "/page/test").Return(domain.Redirect{
					To: RedirectPath, From: "/page/test", Code: 308,
				}, nil)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			mock := &mocks.Repository{}
			if test.mock != nil {
				test.mock(mock)
			}

			t.Engine.Use(Redirects(&deps.Deps{
				Store: &store.Repository{Redirects: mock},
				Config: &domain.ThemeConfig{
					Admin: domain.AdminConfig{
						Path: "admin",
					},
				},
			}))

			t.RequestAndServe(http.MethodGet, test.url, test.url, nil, t.DefaultHandler)

			t.Equal(test.status, t.Status())
			if test.redirectURL != "" {
				loc, err := t.Recorder.Result().Location()
				t.NoError(err)
				t.Equal(test.redirectURL, loc.Path)
			}

			t.Context.Request.Body.Close()
			t.Reset()
		})
	}
}
