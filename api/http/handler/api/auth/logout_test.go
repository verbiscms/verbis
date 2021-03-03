// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/ainsleyclark/verbis/api/errors"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"net/http"
	"time"
)

func (t *AuthTestSuite) TestAuth_Logout() {
	token := "test"

	tt := map[string]struct {
		want    interface{}
		status  int
		message string
		input   string
		cookie  bool
		mock    func(m *mocks.AuthRepository)
	}{
		"Success": {
			nil,
			200,
			"Successfully logged out",
			"test",
			true,
			func(m *mocks.AuthRepository) {
				m.On("Logout", token).Return(-1, nil)
			},
		},
		"Not Found": {
			nil,
			400,
			"not found",
			token,
			false,
			func(m *mocks.AuthRepository) {
				m.On("Logout", token).Return(-1, &errors.Error{Code: errors.NOTFOUND, Message: "not found"})
			},
		},
		"Internal Error": {
			nil,
			500,
			"internal",
			token,
			false,
			func(m *mocks.AuthRepository) {
				m.On("Logout", token).Return(-1, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.NewRequest(http.MethodPost, "/logout", nil)
			t.Context.Request.Header.Set("token", test.input)

			t.Setup(test.mock).Logout(t.Context)

			if test.cookie {
				cookie := http.Cookie{
					Name:     "verbis-session",
					Expires:  time.Time{},
					MaxAge:   -1,
					Path:     "/",
					Raw:      "verbis-session=; Path=/; Max-Age=0; HttpOnly",
					HttpOnly: true,
				}
				t.Equal(t.Recorder.Result().Cookies()[0], &cookie)
			}

			t.RunT(test.want, test.status, test.message)
		})
	}
}
