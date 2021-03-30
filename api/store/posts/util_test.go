// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	users "github.com/ainsleyclark/verbis/api/mocks/store/users"
)

func (t *PostsTestSuite) TestStore_CheckOwner() {
	tt := map[string]struct {
		input int
		mock  func(m *users.Repository)
		want  interface{}
	}{
		"Default": {
			1,
			func(m *users.Repository) {
				m.On("Exists", 1).Return(false)
			},
			1,
		},
		"Exists": {
			2,
			func(m *users.Repository) {
				m.On("Exists", 2).Return(true)
			},
			2,
		},
		"Zero Value": {
			0,
			func(m *users.Repository) {
				m.On("Exists", 0).Return(true)
			},
			1,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil)

			u := &users.Repository{}
			test.mock(u)
			s.users = u

			got := s.checkOwner(test.input)
			t.Equal(test.want, got)
		})
	}
}
