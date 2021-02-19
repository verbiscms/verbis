// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
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
		mock    func(m *mocks.AuthRepository)
		url     string
	}{
		"Success": {
			nil,
			200,
			"Successfully verified token",
			token,
			func(m *mocks.AuthRepository) {
				m.On("VerifyPasswordToken", token).Return(nil)
			},
			"/verify/" + token,
		},
		"Not Found": {
			nil,
			404,
			"not found",
			token,
			func(m *mocks.AuthRepository) {
				m.On("VerifyPasswordToken", token).Return(&errors.Error{Code: errors.NOTFOUND, Message: "not found"})
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