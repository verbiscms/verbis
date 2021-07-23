// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package categories

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/pagination"
	mocks "github.com/verbiscms/verbis/api/mocks/store/categories"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/store/categories"
	"github.com/verbiscms/verbis/api/test/dummy"
	"github.com/verbiscms/verbis/api/tpl/params"
	"testing"
)

type noStringer struct{}

func Setup() (*Namespace, *mocks.Repository) {
	mock := &mocks.Repository{}
	return &Namespace{deps: &deps.Deps{
		Store: &store.Repository{
			Categories: mock,
		},
	}}, mock
}

func TestNamespace_Find(t *testing.T) {
	category := domain.Category{Id: 1, Name: "cat"}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Success": {
			1,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(category, nil)
			},
			category,
		},
		"Not Found": {
			1,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
		},
		"No Stringer": {
			noStringer{},
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(category, nil)
			},
			nil,
		},
		"Nil": {
			nil,
			func(m *mocks.Repository) {
				m.On("Find", 1).Return(category, nil)
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

func TestNamespace_ByName(t *testing.T) {
	category := domain.Category{Id: 1, Name: "cat"}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Success": {
			"cat",
			func(m *mocks.Repository) {
				m.On("FindByName", "cat").Return(category, nil)
			},
			category,
		},
		"Not Found": {
			"cat",
			func(m *mocks.Repository) {
				m.On("FindByName", "cat").Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
		},
		"No Stringer": {
			noStringer{},
			func(m *mocks.Repository) {
				m.On("FindByName", "cat").Return(category, nil)
			},
			nil,
		},
		"Nil": {
			nil,
			func(m *mocks.Repository) {
				m.On("FindByName", "cat").Return(category, nil)
			},
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup()
			test.mock(mock)
			got := ns.ByName(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Parent(t *testing.T) {
	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Success": {
			1,
			func(m *mocks.Repository) {
				m.On("FindParent", 1).Return(domain.Category{Id: 1, Name: "cat"}, nil)
			},
			domain.Category{Id: 1, Name: "cat"},
		},
		"Not Found": {
			1,
			func(m *mocks.Repository) {
				m.On("FindParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
		},
		"Nil Parent": {
			1,
			func(m *mocks.Repository) {
				m.On("FindParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
		},
		"No Stringer": {
			noStringer{},
			func(m *mocks.Repository) {
				m.On("FindParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
		},
		"Nil": {
			nil,
			func(m *mocks.Repository) {
				m.On("FindParent", 1).Return(domain.Category{}, fmt.Errorf("error"))
			},
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			t.Run(name, func(t *testing.T) {
				ns, mock := Setup()
				test.mock(mock)
				got := ns.Parent(test.input)
				assert.Equal(t, test.want, got)
			})
		})
	}
}

func TestNamespace_List(t *testing.T) {
	c := domain.Categories{
		{Id: 1, Name: "cat1"},
		{Id: 1, Name: "cat2"},
	}

	cfg := categories.ListConfig{
		Resource: "",
	}

	tt := map[string]struct {
		input params.Query
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Success": {
			params.Query{"limit": 15},
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams, cfg).Return(c, 2, nil)
			},
			Categories{
				Categories: c,
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
				m.On("List", dummy.DefaultParams, cfg).Return(c, 2, nil)
			},
			Categories{
				Categories: c,
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
				m.On("List", dummy.DefaultParams, cfg).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no categories found"})
			},
			nil,
		},
		"Internal Error": {
			params.Query{"limit": 15},
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams, cfg).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "config error"})
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
