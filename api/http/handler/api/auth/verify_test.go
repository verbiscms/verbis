// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/store/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *AuthTestSuite) TestAuth_VerifyPasswordToken() {
	token := "test"

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   string
		mock    func(m *mocks.Repository)
		url     string
	}{
		"Success": {
			nil,
			http.StatusOK,
			"Successfully verified token",
			token,
			func(m *mocks.Repository) {
				m.On("VerifyPasswordToken", token).Return(domain.PasswordReset{}, nil)
			},
			"/verify/" + token,
		},
		"Not Found": {
			nil,
			http.StatusNotFound,
			"not found",
			token,
			func(m *mocks.Repository) {
				m.On("VerifyPasswordToken", token).Return(domain.PasswordReset{}, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
			"/verify/" + token,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.RequestAndServe(http.MethodGet, test.url, "/verify/:token", nil, func(ctx *gin.Context) {
				t.Setup(test.mock).VerifyPasswordToken(ctx)
			})
			t.RunT(test.want, test.status, test.message)
		})
	}
}
