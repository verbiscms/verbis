// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package attributes

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/store/users"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/tpl/funcs/auth"
	"github.com/verbiscms/verbis/api/tpl/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNamespace_Body(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tt := map[string]struct {
		post   domain.Post
		cookie bool
		want   string
	}{
		"ID": {
			domain.Post{
				ID:           123,
				Title:        "title",
				Resource:     "",
				PageTemplate: "template",
				PageLayout:   "layout",
			},
			false,
			"page page-id-123 page-title-title page-template-template page-layout-layout",
		},
		"Resource": {
			domain.Post{
				ID:           1,
				Title:        "title",
				Resource:     "resource",
				PageTemplate: "template",
				PageLayout:   "layout",
			},
			false,
			"resource page-id-1 page-title-title page-template-template page-layout-layout",
		},
		"template": {
			domain.Post{
				ID:           1,
				Title:        "title",
				Resource:     "resource",
				PageTemplate: "%$££@template*&",
				PageLayout:   "layout",
			},
			false,
			"resource page-id-1 page-title-title page-template-template page-layout-layout",
		},
		"Logged In": {
			domain.Post{
				ID:           1,
				Title:        "title",
				Resource:     "",
				PageTemplate: "template",
				PageLayout:   "layout",
			},
			true,
			"page page-id-1 page-title-title page-template-template page-layout-layout logged-in",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			post := &domain.PostDatum{Post: test.post}

			ns := Namespace{
				deps: nil,
				tpld: &internal.TemplateDeps{Post: post},
			}

			mock := &mocks.Repository{}
			mock.On("FindByToken", "token").Return(domain.User{}, nil)

			rr := httptest.NewRecorder()
			g, _ := gin.CreateTestContext(rr)
			g.Request, _ = http.NewRequest("GET", "/get", nil)
			ns.auth = auth.New(
				&deps.Deps{Store: &store.Repository{User: mock}},
				&internal.TemplateDeps{Context: g, Post: post},
			)

			if test.cookie {
				g.Request.Header.Set("Cookie", "verbis-session=token")
			}

			got := ns.Body()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Homepage(t *testing.T) {
	tt := map[string]struct {
		input int
		want  bool
	}{
		"Is Homepage": {
			1,
			true,
		},
		"Not Homepage": {
			2,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			post := &domain.PostDatum{Post: domain.Post{ID: test.input}}
			ns := Namespace{
				deps: &deps.Deps{Options: &domain.Options{Homepage: 1}},
				tpld: &internal.TemplateDeps{Post: post},
			}
			got := ns.Homepage()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_CSSValidString(t *testing.T) {
	tt := map[string]struct {
		input string
		want  string
	}{
		"Regex 1": {
			"£$verbis$£$",
			"verbis",
		},
		"Regex 2": {
			"£@$@£$$verbis{}|%$@£%",
			"verbis",
		},
		"Spaces": {
			"verbis cms",
			"verbis-cms",
		},
		"Forward Slash": {
			"verbis/cms/is/the/best",
			"verbis-cms-is-the-best",
		},
		"Capital Letters": {
			"Verbis CMS",
			"verbis-cms",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := cssValidString(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_Lang(t *testing.T) {
	ns := New(&deps.Deps{
		Options: &domain.Options{
			GeneralLocale: "en-gb",
		},
	}, &internal.TemplateDeps{})
	got := ns.Lang()
	assert.Equal(t, "en-gb", got)
}
