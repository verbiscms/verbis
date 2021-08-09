// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/store/users"
	"github.com/verbiscms/verbis/api/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Setup(cookie string) (*Namespace, *mocks.Repository) {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	g, _ := gin.CreateTestContext(rr)
	g.Request, _ = http.NewRequest("GET", "/get", nil)
	g.Request.Header.Set("Cookie", cookie)

	mock := &mocks.Repository{}
	return &Namespace{
		deps: &deps.Deps{
			Store: &store.Repository{
				User: mock,
			},
		},
		ctx: g,
	}, mock
}

func Test_Auth(t *testing.T) {
	tt := map[string]struct {
		want   interface{}
		cookie string
		mock   func(m *mocks.Repository)
	}{
		"Logged In": {
			want:   true,
			cookie: "verbis-session=token",
			mock: func(m *mocks.Repository) {
				m.On("FindByToken", "token").Return(domain.User{}, nil)
			},
		},
		"No Cookie": {
			want:   false,
			cookie: "",
			mock: func(m *mocks.Repository) {
				m.On("FindByToken", "token").Return(domain.User{}, nil)
			},
		},
		"No User": {
			want:   false,
			cookie: "verbis-session=token",
			mock: func(m *mocks.Repository) {
				m.On("FindByToken", "token").Return(domain.User{}, fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup(test.cookie)
			test.mock(mock)
			got := ns.Auth()
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_Admin(t *testing.T) {
	tt := map[string]struct {
		want   interface{}
		cookie string
		mock   func(m *mocks.Repository)
	}{
		"Is Admin": {
			want:   true,
			cookie: "verbis-session=token",
			mock: func(m *mocks.Repository) {
				m.On("FindByToken", "token").Return(domain.User{
					UserPart: domain.UserPart{ID: 0, Role: domain.Role{ID: 6}},
				}, nil)
			},
		},
		"Not Admin": {
			want:   false,
			cookie: "verbis-session=token",
			mock: func(m *mocks.Repository) {
				m.On("FindByToken", "token").Return(domain.User{
					UserPart: domain.UserPart{ID: 0, Role: domain.Role{ID: 1}},
				}, nil)
			},
		},
		"No Cookie": {
			want:   false,
			cookie: "",
			mock: func(m *mocks.Repository) {
				m.On("FindByToken", "token").Return(domain.User{}, nil)
			},
		},
		"No User": {
			want:   false,
			cookie: "verbis-session=token",
			mock: func(m *mocks.Repository) {
				m.On("FindByToken", "token").Return(domain.User{}, fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup(test.cookie)
			test.mock(mock)
			got := ns.Admin()
			assert.Equal(t, test.want, got)
		})
	}
}
