// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/mocks/store/users"
)

func (t *AuthTestSuite) TestStore_Logout() {
	tt := map[string]struct {
		want interface{}
		mock func(m *mocks.Repository)
	}{
		"Success": {
			user.Id,
			func(m *mocks.Repository) {
				m.On("FindByToken", user.Token).Return(user, nil)
				m.On("UpdateToken", "newtoken").Return(nil)
			},
		},
		"Find Error": {
			"error",
			func(m *mocks.Repository) {
				m.On("FindByToken", user.Token).Return(domain.User{}, fmt.Errorf("error"))
			},
		},
		"Update Error": {
			"error",
			func(m *mocks.Repository) {
				m.On("FindByToken", user.Token).Return(user, nil)
				m.On("UpdateToken", "newtoken").Return(fmt.Errorf("error"))
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			u := &mocks.Repository{}
			test.mock(u)
			s := &Store{
				UserStore: u,
				generateTokeFunc: func(name, email string) string {
					return "newtoken"
				},
			}
			id, err := s.Logout(user.Token)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, id)
		})
	}
}
