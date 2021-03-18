// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *ResolverTestSuite) TestValue_User() {
	tt := map[string]struct {
		value domain.FieldValue
		mock  func(u *mocks.UserRepository)
		want  interface{}
	}{
		"User": {
			value: domain.FieldValue("1"),
			mock: func(u *mocks.UserRepository) {
				u.On("GetByID", 1).Return(domain.User{
					UserPart: domain.UserPart{FirstName: "user"},
				}, nil)
			},
			want: domain.UserPart{FirstName: "user"},
		},
		"User Error": {
			value: domain.FieldValue("1"),
			mock: func(u *mocks.UserRepository) {
				u.On("GetByID", 1).Return(domain.User{}, fmt.Errorf("not found"))
			},
			want: "not found",
		},
		"Cast Error": {
			value: domain.FieldValue("wrongval"),
			mock:  func(u *mocks.UserRepository) {},
			want:  `strconv.Atoi: parsing "wrongval": invalid syntax`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()
			userMock := &mocks.UserRepository{}

			test.mock(userMock)
			v.deps.Store.User = userMock

			got, err := v.user(test.value)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *ResolverTestSuite) TestValue_UserResolve() {
	tt := map[string]struct {
		field domain.PostField
		mock  func(u *mocks.UserRepository)
		want  interface{}
	}{
		"Success": {
			field: domain.PostField{OriginalValue: "1,2,3", Type: "user"},
			mock: func(u *mocks.UserRepository) {
				u.On("GetByID", 1).Return(domain.User{UserPart: domain.UserPart{FirstName: "user1"}}, nil)
				u.On("GetByID", 2).Return(domain.User{UserPart: domain.UserPart{FirstName: "user2"}}, nil)
				u.On("GetByID", 3).Return(domain.User{UserPart: domain.UserPart{FirstName: "user3"}}, nil)
			},
			want: domain.PostField{OriginalValue: "1,2,3", Type: "user", Value: []interface{}{
				domain.UserPart{FirstName: "user1"},
				domain.UserPart{FirstName: "user2"},
				domain.UserPart{FirstName: "user3"},
			}},
		},
		"Trailing Comma": {
			field: domain.PostField{OriginalValue: "1,2,3,", Type: "user"},
			mock: func(u *mocks.UserRepository) {
				u.On("GetByID", 1).Return(domain.User{UserPart: domain.UserPart{FirstName: "user1"}}, nil)
				u.On("GetByID", 2).Return(domain.User{UserPart: domain.UserPart{FirstName: "user2"}}, nil)
				u.On("GetByID", 3).Return(domain.User{UserPart: domain.UserPart{FirstName: "user3"}}, nil)
			},
			want: domain.PostField{OriginalValue: "1,2,3,", Type: "user", Value: []interface{}{
				domain.UserPart{FirstName: "user1"},
				domain.UserPart{FirstName: "user2"},
				domain.UserPart{FirstName: "user3"},
			}},
		},
		"Leading Comma": {
			field: domain.PostField{OriginalValue: ",1,2,3", Type: "user"},
			mock: func(u *mocks.UserRepository) {
				u.On("GetByID", 1).Return(domain.User{UserPart: domain.UserPart{FirstName: "user1"}}, nil)
				u.On("GetByID", 2).Return(domain.User{UserPart: domain.UserPart{FirstName: "user2"}}, nil)
				u.On("GetByID", 3).Return(domain.User{UserPart: domain.UserPart{FirstName: "user3"}}, nil)
			},
			want: domain.PostField{OriginalValue: ",1,2,3", Type: "user", Value: []interface{}{
				domain.UserPart{FirstName: "user1"},
				domain.UserPart{FirstName: "user2"},
				domain.UserPart{FirstName: "user3"},
			}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()
			userMock := &mocks.UserRepository{}

			test.mock(userMock)
			v.deps.Store.User = userMock

			got := v.resolve(test.field)

			t.Equal(test.want, got)
		})
	}
}
