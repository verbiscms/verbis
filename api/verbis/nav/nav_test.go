// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nav

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	posts "github.com/verbiscms/verbis/api/mocks/store/posts"
	"github.com/verbiscms/verbis/api/store"
	"io/ioutil"
	"testing"
)

var (
	postID = 1
	post   = domain.PostDatum{
		Post: domain.Post{
			ID:        1,
			Title:     "title",
			Permalink: "/news/item",
		},
	}
)

func TestNew(t *testing.T) {
	tt := map[string]struct {
		input *domain.Options
		want  interface{}
	}{
		"Success": {
			&domain.Options{NavMenus: map[string]interface{}{}},
			nil,
		},
		"Marshal Error": {
			&domain.Options{NavMenus: map[string]interface{}{
				"test": make(chan bool),
			}},
			"Error marshalling navigation menus",
		},
		"Unmarshal Error": {
			&domain.Options{NavMenus: map[string]interface{}{
				"test": 2,
			}},
			"Error unmarshalling navigation menus",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			d := &deps.Deps{Options: test.input}
			_, got := New(d, &domain.PostDatum{})
			if got != nil {
				assert.Contains(t, errors.Message(got), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGet(t *testing.T) {
	logger.SetOutput(ioutil.Discard)

	tt := map[string]struct {
		input      Args
		navigation Nav
		post       *domain.PostDatum
		mock       func(m *posts.Repository, c *cache.Store)
		want       interface{}
	}{
		"Args Marshal Error": {
			Args{"menu": make(chan bool)},
			nil,
			nil,
			nil,
			"Error converting arguments to navigation options",
		},
		"Args Unmarshal Error": {
			Args{"menu": 2},
			nil,
			nil,
			nil,
			"Error converting arguments to navigation options",
		},
		"Validation Failed": {
			Args{},
			nil,
			nil,
			nil,
			"Error validating navigation options",
		},
		"Menu Not Found": {
			Args{"menu": "main-menu"},
			nil,
			nil,
			func(m *posts.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, "nav-menu-main-menu").Return(nil, fmt.Errorf("error"))
			},
			"Error obtaining navigation items",
		},
		"From Cache": {
			Args{"menu": "main-menu"},
			nil,
			nil,
			func(m *posts.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, "nav-menu-main-menu").Return(Items{{Href: "link-one", Title: "title"}}, nil)
			},
			Items{{Href: "link-one", Title: "title"}},
		},
		"Simple": {
			Args{"menu": "main-menu"},
			Nav{"main-menu": Items{{Href: "link-one", Title: "title"}}},
			nil,
			func(m *posts.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, "nav-menu-main-menu").Return(nil, fmt.Errorf("error"))
			},
			Items{{Href: "link-one", Title: "title"}},
		},
		"Children": {
			Args{"menu": "main-menu"},
			Nav{"main-menu": Items{
				{Href: "link-one", HasChildren: true, Children: Items{
					{Href: "link-two"},
				}},
			}},
			nil,
			func(m *posts.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, "nav-menu-main-menu").Return(nil, fmt.Errorf("error"))
			},
			Items{{Href: "link-one", HasChildren: true, Children: Items{
				{Href: "link-two"},
			}}},
		},
		"External": {
			Args{"menu": "main-menu"},
			Nav{"main-menu": Items{{Href: "link-one", External: true}}},
			nil,
			func(m *posts.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, "nav-menu-main-menu").Return(nil, fmt.Errorf("error"))
			},
			Items{{Href: "link-one", External: true}},
		},
		"Post": {
			Args{"menu": "main-menu"},
			Nav{"main-menu": Items{{Href: "link-one", PostID: &postID}}},
			&domain.PostDatum{},
			func(m *posts.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, "nav-menu-main-menu").Return(nil, fmt.Errorf("error"))
				m.On("Find", postID, false).Return(post, nil)
			},
			Items{{Href: post.Permalink, Title: post.Title, PostID: &postID, External: false}},
		},
		"Post Error": {
			Args{"menu": "main-menu"},
			Nav{"main-menu": Items{{Href: "link-one", PostID: &postID}}},
			&domain.PostDatum{},
			func(m *posts.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, "nav-menu-main-menu").Return(nil, fmt.Errorf("error"))
				m.On("Find", postID, false).Return(domain.PostDatum{}, fmt.Errorf("error"))
			},
			Items{{Href: "link-one", PostID: &postID}},
		},
		"Current": {
			Args{"menu": "main-menu"},
			Nav{"main-menu": Items{{Href: "link-one", PostID: &postID}}},
			&post,
			func(m *posts.Repository, c *cache.Store) {
				c.On("Get", mock.Anything, "nav-menu-main-menu").Return(nil, fmt.Errorf("error"))
				m.On("Find", postID, false).Return(post, nil)
			},
			Items{{Href: post.Permalink, PostID: &postID, Title: post.Title, External: false, IsActive: true}},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			m := &posts.Repository{}
			c := &cache.Store{}
			c.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			if test.mock != nil {
				test.mock(m, c)
			}
			s := Service{
				nav: test.navigation,
				deps: &deps.Deps{
					Store: &store.Repository{Posts: m},
					Cache: c,
				},
				post: test.post,
			}
			got, err := s.Get(test.input)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
