// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package posts

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/domain"
	mocks "github.com/verbiscms/verbis/api/mocks/store/options"
)

func (t *PostsTestSuite) TestStore_Format() {
	key := uuid.New()

	tt := map[string]struct {
		layout bool
		raw    []postsRaw
		want   interface{}
	}{
		"Simple": {
			false,
			[]postsRaw{
				{Post: postData[0].Post}, {Post: postData[1].Post},
			},
			postData,
		},
		"Duplicate": {
			false,
			[]postsRaw{
				{Post: domain.Post{Id: 1, Slug: "slug"}}, {Post: domain.Post{Id: 1, Slug: "slug"}},
			},
			domain.PostData{
				{
					Post:   domain.Post{Id: 1, Slug: "slug", Permalink: "/slug"},
					Fields: domain.PostFields{},
				},
			},
		},
		"With Fields": {
			false,
			[]postsRaw{
				{
					Post:  domain.Post{Id: 1, Slug: "slug", Title: "post", Permalink: "/slug"},
					Field: postsRawFields{Name: "name", PostId: 1, UUID: &key},
				},
			},
			domain.PostData{
				{
					Post:   domain.Post{Id: 1, Slug: "slug", Title: "post", Permalink: "/slug"},
					Fields: domain.PostFields{domain.PostField{Name: "name", PostId: 1, UUID: key}},
				},
			},
		},
		"With Multiple Fields": {
			false,
			[]postsRaw{
				{
					Post:  domain.Post{Id: 1, Slug: "slug", Title: "post", Permalink: "/slug"},
					Field: postsRawFields{Name: "name", PostId: 1, UUID: &key},
				},
				{
					Post:  domain.Post{Id: 2, Slug: "slug", Title: "post", Permalink: "/slug"},
					Field: postsRawFields{Name: "name", PostId: 2, UUID: &key},
				},
			},
			domain.PostData{
				{
					Post:   domain.Post{Id: 1, Slug: "slug", Title: "post", Permalink: "/slug"},
					Fields: domain.PostFields{domain.PostField{Name: "name", PostId: 1, UUID: key}},
				},
				{
					Post:   domain.Post{Id: 2, Slug: "slug", Title: "post", Permalink: "/slug"},
					Fields: domain.PostFields{domain.PostField{Name: "name", PostId: 2, UUID: key}},
				},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.Setup(nil)
			got := s.format(test.raw, test.layout)
			t.Equal(test.want, got)
		})
	}
}

func (t *PostsTestSuite) TestStore_Format_Panic() {
	s := t.Setup(nil)
	m := &mocks.Repository{}
	m.On("GetTheme", mock.Anything).Return(nil, fmt.Errorf("error"))
	s.options = m
	assert.Panics(t.T(), func() {
		s.format([]postsRaw{
			{Post: postData[0].Post}, {Post: postData[1].Post},
		}, true)
	})
}
