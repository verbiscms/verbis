// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/http/handler/api"
	mocks "github.com/verbiscms/verbis/api/mocks/store/users"
	"github.com/verbiscms/verbis/api/store/users"
	"net/http"
)

func (t *MiddlewareTestSuite) TestTokenCheck() {
	token := "token"
	user := domain.User{UserPart: domain.UserPart{Role: domain.Role{ID: domain.OwnerRoleID}}}

	tt := map[string]struct {
		mock func(m *mocks.Repository)
		status int
		want   string
	}{
		"No User": {
			func(m *mocks.Repository) {
				m.On("FindByToken", token).
					Return(domain.User{}, fmt.Errorf("error"))
			},
			http.StatusUnauthorized,
			"Invalid token in the request header",
		},
		"Session Expired": {
			func(m *mocks.Repository) {
				m.On("FindByToken", token).
					Return(domain.User{}, nil)
				m.On("CheckSession", token).
					Return(users.ErrSessionExpired)
			},
			http.StatusUnauthorized,
			"Session expired, please login again",
		},
		"Invalid Token": {
			func(m *mocks.Repository) {
				m.On("FindByToken", token).
					Return(domain.User{}, nil)
				m.On("CheckSession", token).
					Return(fmt.Errorf("error"))
			},
			http.StatusUnauthorized,
			"Invalid token in the request header",
		},
		"Banned": {
			func(m *mocks.Repository) {
				m.On("FindByToken", token).
					Return(domain.User{UserPart: domain.UserPart{Role: domain.Role{ID: domain.BannedRoleID}}}, nil)
				m.On("CheckSession", token).
					Return(nil)
			},
			http.StatusForbidden,
			"Your account has been suspended",
		},
		"Authorised": {
			func(m *mocks.Repository) {
				m.On("FindByToken", token).
					Return(user, nil)
				m.On("CheckSession", token).
					Return(nil)
			},
			http.StatusOK,
			"success",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			defer t.Reset()

			m := &mocks.Repository{}
			if test.mock != nil {
				test.mock(m)
			}
			t.Engine.Use(TokenCheck(m))

			t.Engine.GET("/test", func(ctx *gin.Context) {
				ctxUser, ok := ctx.Get("user")
				t.True(ok)
				t.Equal(ctxUser, user)
				api.Respond(ctx, http.StatusOK, "success", nil)
			})

			t.NewRequest(http.MethodGet, "/test", nil)
			t.Context.Request.Header.Set("token", token)
			t.ServeHTTP()

			respond, _ := t.RespondData()
			t.Equal(test.status, t.Status())
			t.Contains(respond.Message, test.want)
		})
	}
}
