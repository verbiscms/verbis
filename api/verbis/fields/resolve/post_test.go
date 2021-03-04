// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *ResolverTestSuite) TestValue_Post() {
	tt := map[string]struct {
		value domain.FieldValue
		mock  func(p *mocks.PostsRepository)
		want  interface{}
	}{
		"Post": {
			value: domain.FieldValue("1"),
			mock: func(p *mocks.PostsRepository) {
				p.On("GetByID", 1, false).Return(domain.PostDatum{Post: domain.Post{Title: "post"}}, nil)
			},
			want: domain.PostDatum{
				Post: domain.Post{Title: "post"},
			},
		},
		"Post Error": {
			value: domain.FieldValue("1"),
			mock: func(p *mocks.PostsRepository) {
				p.On("GetByID", 1, false).Return(domain.PostDatum{}, fmt.Errorf("not found"))
			},
			want: "not found",
		},
		"Cast Error": {
			value: domain.FieldValue("wrongval"),
			mock:  func(p *mocks.PostsRepository) {},
			want:  `strconv.Atoi: parsing "wrongval": invalid syntax`,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()
			postMock := &mocks.PostsRepository{}

			test.mock(postMock)
			v.deps.Store.Posts = postMock

			got, err := v.post(test.value)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *ResolverTestSuite) TestValue_PostResolve() {
	tt := map[string]struct {
		field domain.PostField
		mock  func(p *mocks.PostsRepository)
		want  domain.PostField
	}{
		"Post": {
			field: domain.PostField{OriginalValue: "1,2,3", Type: "post"},
			mock: func(p *mocks.PostsRepository) {
				p.On("GetByID", 1, false).Return(domain.PostDatum{Post: domain.Post{Title: "post1"}}, nil)
				p.On("GetByID", 2, false).Return(domain.PostDatum{Post: domain.Post{Title: "post2"}}, nil)
				p.On("GetByID", 3, false).Return(domain.PostDatum{Post: domain.Post{Title: "post3"}}, nil)
			},
			want: domain.PostField{OriginalValue: "1,2,3", Type: "post", Value: []interface{}{
				domain.PostDatum{Post: domain.Post{Title: "post1"}},
				domain.PostDatum{Post: domain.Post{Title: "post2"}},
				domain.PostDatum{Post: domain.Post{Title: "post3"}},
			}},
		},
		"Trailing Comma": {
			field: domain.PostField{OriginalValue: "1,2,3,", Type: "post"},
			mock: func(p *mocks.PostsRepository) {
				p.On("GetByID", 1, false).Return(domain.PostDatum{Post: domain.Post{Title: "post1"}}, nil)
				p.On("GetByID", 2, false).Return(domain.PostDatum{Post: domain.Post{Title: "post2"}}, nil)
				p.On("GetByID", 3, false).Return(domain.PostDatum{Post: domain.Post{Title: "post3"}}, nil)
			},
			want: domain.PostField{OriginalValue: "1,2,3,", Type: "post", Value: []interface{}{
				domain.PostDatum{Post: domain.Post{Title: "post1"}},
				domain.PostDatum{Post: domain.Post{Title: "post2"}},
				domain.PostDatum{Post: domain.Post{Title: "post3"}},
			}},
		},
		"Leading Comma": {
			field: domain.PostField{OriginalValue: ",1,2,3", Type: "post"},
			mock: func(p *mocks.PostsRepository) {
				p.On("GetByID", 1, false).Return(domain.PostDatum{Post: domain.Post{Title: "post1"}}, nil)
				p.On("GetByID", 2, false).Return(domain.PostDatum{Post: domain.Post{Title: "post2"}}, nil)
				p.On("GetByID", 3, false).Return(domain.PostDatum{Post: domain.Post{Title: "post3"}}, nil)
			},
			want: domain.PostField{OriginalValue: ",1,2,3", Type: "post", Value: []interface{}{
				domain.PostDatum{Post: domain.Post{Title: "post1"}},
				domain.PostDatum{Post: domain.Post{Title: "post2"}},
				domain.PostDatum{Post: domain.Post{Title: "post3"}},
			}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()
			postMock := &mocks.PostsRepository{}

			test.mock(postMock)
			v.deps.Store.Posts = postMock

			got := v.resolve(test.field)

			t.Equal(test.want, got)
		})
	}
}
