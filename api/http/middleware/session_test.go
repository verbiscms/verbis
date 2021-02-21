// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"net/http"
)

const (
	// The default token used for testing.
	Token = "token"
)

func (t *MiddlewareTestSuite) Test_SessionCheck() {

	tt := map[string]struct {
		want interface{}
		status      int
		message string
		cookie *http.Cookie
		mock        func(m *mocks.UserRepository)
	}{
		"Expired": {
			`{"errors":{"session":"expired"}}`,
			401,
			"Session expired, please login again",
			&http.Cookie{
				Name:     "verbis-session",
				Value:    "",
				Path:     "/",
				Raw:      "verbis-session=; Path=/; Max-Age=0; HttpOnly",
				Domain:   "",
				MaxAge:   -1,
				Secure:   false,
				HttpOnly: true,
			},
			func(m *mocks.UserRepository) {
				m.On("CheckSession", Token).Return(fmt.Errorf("error"))
			},
		},
		"Continue": {
			"",
			200,
			"",
			nil,
			func(m *mocks.UserRepository) {
				m.On("CheckSession", Token).Return(nil)
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			mock := &mocks.UserRepository{}
			if test.mock != nil {
				test.mock(mock)
			}

			t.Engine.Use(SessionCheck(&deps.Deps{
				Store:   &models.Store{User: mock},
			}))

			t.NewRequest(http.MethodGet, "/test", nil)
			t.Context.Request.Header.Set("token", Token)
			t.Engine.GET("/test", t.DefaultHandler)
			t.ServeHTTP()

			if test.cookie != nil {
				t.Equal(test.cookie, t.Recorder.Result().Cookies()[0])
				t.RunT(test.want, test.status, test.message)
			} else {
				t.Equal(test.status, t.Status())
			}

			t.Reset()
		})
	}
}