// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	params2 "github.com/ainsleyclark/verbis/api/helpers/params"
	vhttp "github.com/ainsleyclark/verbis/api/http"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/tpl/params"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	cat  = &domain.Category{}
	post = domain.Post{
		Id:     1,
		Title:  "test title",
		UserId: 1,
	}
	postData = domain.PostData{
		Post:     post,
		Author:   domain.UserPart{},
		Category: cat,
	}
	postDataSlice = []domain.PostData{
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

func Setup() (*Namespace, *mocks.PostsRepository) {
	mock := &mocks.PostsRepository{}
	return &Namespace{deps: &deps.Deps{
		Store: &models.Store{
			Posts: mock,
		},
	}}, mock
}

func TestNamespace_Find(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.PostsRepository)
		want  interface{}
	}{
		"Success": {
			1,
			func(m *mocks.PostsRepository) {
				m.On("GetById", 1, false).Return(postData, nil)
			},
			tplPost,
		},
		"Not Found": {
			1,
			func(m *mocks.PostsRepository) {
				m.On("GetById", 1, false).Return(domain.PostData{}, fmt.Errorf("error"))
			},
			nil,
		},
		"No Stringer": {
			noStringer{},
			func(m *mocks.PostsRepository) {
				m.On("GetById", 1, false).Return(postData, nil)
			},
			nil,
		},
		"Nil": {
			nil,
			func(m *mocks.PostsRepository) {
				m.On("GetById", 1, false).Return(postData, nil)
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

	p := params2.Params{
		Page:           1,
		Limit:          15,
		LimitAll:       false,
		OrderBy:        OrderBy,
		OrderDirection: OrderDirection,
	}

	tt := map[string]struct {
		input params.Query
		mock  func(m *mocks.PostsRepository)
		want  interface{}
	}{
		"Success": {
			params.Query{"limit": 15},
			func(m *mocks.PostsRepository) {
				m.On("Get", p, false, "", "published").Return(postDataSlice, 5, nil)
			},
			Posts{
				Posts: tplPostSlice,
				Pagination: &vhttp.Pagination{
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
			func(m *mocks.PostsRepository) {
				m.On("Get", p, false, "", "published").Return(postDataSlice, 5, nil)
			},
			Posts{
				Posts: tplPostSlice,
				Pagination: &vhttp.Pagination{
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
			func(m *mocks.PostsRepository) {
				m.On("Get", p, false, "", "published").Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no posts found"})
			},
			nil,
		},
		"Internal Error": {
			params.Query{"limit": 15},
			func(m *mocks.PostsRepository) {
				m.On("Get", p, false, "", "published").Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal error"})
			},
			"internal error",
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
