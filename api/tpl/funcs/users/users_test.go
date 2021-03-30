// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/http/pagination"
	"github.com/ainsleyclark/verbis/api/store"
	"github.com/ainsleyclark/verbis/api/test/dummy"
	"github.com/ainsleyclark/verbis/api/tpl/params"
	"github.com/stretchr/testify/assert"
	"testing"

	mocks "github.com/ainsleyclark/verbis/api/mocks/store/users"
)

type noStringer struct{}

func Setup() (*Namespace, *mocks.Repository) {
	mock := &mocks.Repository{}
	return &Namespace{deps: &deps.Deps{
		Store: &store.Repository{
			User: mock,
		},
	}}, mock
}

func TestNamespace_Find(t *testing.T) {
	user := domain.User{
		UserPart: domain.UserPart{Id: 1, FirstName: "verbis"},
	}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Success": {
			1,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(user, nil).Once()
			},
			user.HideCredentials(),
		},
		"Not Found": {
			1,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(domain.User{}, fmt.Errorf("error")).Once()
			},
			nil,
		},
		"No Stringer": {
			noStringer{},
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(user, nil).Once()
			},
			nil,
		},
		"Nil": {
			nil,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(user, nil).Once()
			},
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup()
			test.mock(mock)
			got := ns.Find(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_List(t *testing.T) {
	users := domain.Users{
		domain.User{UserPart: domain.UserPart{Id: 1, FirstName: "verbis"}},
		domain.User{UserPart: domain.UserPart{Id: 1, FirstName: "cms"}},
	}

	tt := map[string]struct {
		input params.Query
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Success": {
			params.Query{"limit": 15},
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams, "").Return(users, 2, nil).Once()
			},
			Users{
				Users: users.HideCredentials(),
				Pagination: &pagination.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 2,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Nil": {
			nil,
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams, "").Return(users, 2, nil).Once()
			},
			Users{
				Users: users.HideCredentials(),
				Pagination: &pagination.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 2,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Not Found": {
			params.Query{"limit": 15},
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams, "").Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no users found"}).Once()
			},
			nil,
		},
		"Internal Error": {
			params.Query{"limit": 15},
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams, "").Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "config error"}).Once()
			},
			"config error",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup()
			test.mock(mock)
			got, err := ns.List(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
