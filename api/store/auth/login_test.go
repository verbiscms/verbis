// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"github.com/verbiscms/verbis/api/common/encryption"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/mocks/store/users"
)

func (t *AuthTestSuite) TestStore_Login() {
	hashed, err := encryption.HashPassword("password")
	t.NoError(err)
	user.Password = hashed

	tt := map[string]struct {
		want     interface{}
		password string
		mock     func(m *mocks.Repository)
	}{
		"Success": {
			user,
			"password",
			func(m *mocks.Repository) {
				m.On("FindByEmail", user.Email).Return(user, nil)
				m.On("UpdateToken", "token").Return(nil)
			},
		},
		"No User": {
			ErrLoginMsg,
			user.Password,
			func(m *mocks.Repository) {
				m.On("FindByEmail", user.Email).Return(domain.User{}, fmt.Errorf("error"))
			},
		},
		"Wrong Password": {
			ErrLoginMsg,
			"wrong",
			func(m *mocks.Repository) {
				m.On("FindByEmail", user.Email).Return(user, nil)
			},
		},
		"Update Token Error": {
			"error",
			"password",
			func(m *mocks.Repository) {
				m.On("FindByEmail", user.Email).Return(user, nil)
				m.On("UpdateToken", "token").Return(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			u := &mocks.Repository{}
			test.mock(u)
			s := &Store{
				userStore: u,
			}
			user, err := s.Login(user.Email, test.password)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, user)
		})
	}
}
