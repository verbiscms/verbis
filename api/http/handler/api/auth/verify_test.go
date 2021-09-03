// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	mocks "github.com/verbiscms/verbis/api/mocks/store/auth"
	"net/http"
)

func (t *AuthTestSuite) TestAuth_VerifyPasswordToken() {
	token := "test"

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   string
		mock    func(m *mocks.Repository, c *cache.Store)
		url     string
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully verified token",
			token,
			func(m *mocks.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, token, mock.Anything).Return(nil)
			},
			"/verify/" + token,
		},
		"Not Found": {
			nil,
			http.StatusNotFound,
			"No user exists with the token: " + token,
			token,
			func(m *mocks.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, token, mock.Anything).Return(fmt.Errorf("error"))
			},
			"/verify/" + token,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, test.url, "/verify/:token", nil, func(ctx *gin.Context) {
				t.SetupCache(test.mock).VerifyPasswordToken(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
