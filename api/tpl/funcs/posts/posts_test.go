// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/http/pagination"
	mocks "github.com/verbiscms/verbis/api/mocks/store/posts"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/store/posts"
	"github.com/verbiscms/verbis/api/test/dummy"
	"github.com/verbiscms/verbis/api/tpl/params"
	"testing"
)

var (
	cat  = &domain.Category{}
	post = domain.Post{
		ID:     1,
		Title:  "test title",
		UserID: 1,
	}
	postData = domain.PostDatum{
		Post:     post,
		Author:   domain.UserPart{},
		Category: cat,
	}
	postDataSlice = domain.PostData{
		postData, postData,
	}
	tplPost = domain.PostTemplate{
		Author:   domain.UserPart{},
		Category: cat,
		Post:     post,
	}
	tplPostSlice = []domain.PostTemplate{
		tplPost, tplPost,
	}
)

type noStringer struct{}

func Setup() (*Namespace, *mocks.Repository) {
	mock := &mocks.Repository{}
	return &Namespace{deps: &deps.Deps{
		Store: &store.Repository{
			Posts: mock,
		},
	}}, mock
}

func TestNamespace_Find(t *testing.T) {
	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Success": {
			1,
			func(m *mocks.Repository) {
				m.On("Find", 1, false).Return(postData, nil)
			},
			tplPost,
		},
		"Not Found": {
			1,
			func(m *mocks.Repository) {
				m.On("Find", 1, false).Return(domain.PostDatum{}, fmt.Errorf("error"))
			},
			nil,
		},
		"No Stringer": {
			noStringer{},
			func(m *mocks.Repository) {
				m.On("Find", 1, false).Return(postData, nil)
			},
			nil,
		},
		"Nil": {
			nil,
			func(m *mocks.Repository) {
				m.On("Find", 1, false).Return(postData, nil)
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
	cfg := posts.ListConfig{
		Resource: "",
		Status:   "published",
	}

	tt := map[string]struct {
		input params.Query
		mock  func(m *mocks.Repository)
		want  interface{}
	}{
		"Success": {
			params.Query{"limit": 15},
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams, false, cfg).Return(postDataSlice, 5, nil)
			},
			Posts{
				Posts: tplPostSlice,
				Pagination: &pagination.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 5,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Nil": {
			nil,
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams, false, cfg).Return(postDataSlice, 5, nil)
			},
			Posts{
				Posts: tplPostSlice,
				Pagination: &pagination.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 5,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Not Found": {
			params.Query{"limit": 15},
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams, false, cfg).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no posts found"})
			},
			nil,
		},
		"Internal Error": {
			params.Query{"limit": 15},
			func(m *mocks.Repository) {
				m.On("List", dummy.DefaultParams, false, cfg).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "config error"})
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
