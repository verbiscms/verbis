// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *FieldTestSuite) TestService_GetLayout() {
	tt := map[string]struct {
		id     interface{}
		name   string
		layout domain.FieldGroups
		args   []interface{}
		want   interface{}
		err    bool
	}{
		"Success": {
			id:   1,
			name: "key3",
			layout: domain.FieldGroups{
				{
					Title:  "test1",
					Fields: domain.Fields{{Name: "key1"}, {Name: "key2"}},
				},
				{
					Title:  "test2",
					Fields: domain.Fields{{Name: "key3"}, {Name: "key4"}},
				},
			},
			args: nil,
			want: domain.Field{Name: "key3"},
		},
		"Error": {
			id:     1,
			name:   "key3",
			layout: nil,
			args:   nil,
			want:   "no groups exist",
			err:    true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(nil)
			s.layout = test.layout

			got := s.GetLayout(test.name, test.args...)
			if test.err {
				t.Contains(t.logWriter.String(), test.want)
				t.Reset()
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *FieldTestSuite) TestService_GetLayouts() {
	var f domain.FieldGroups

	fg := domain.FieldGroups{
		{
			Title:  "test1",
			Fields: domain.Fields{{Name: "key1"}, {Name: "key2"}},
		},
		{
			Title:  "test2",
			Fields: domain.Fields{{Name: "key3"}, {Name: "key4"}},
		},
	}

	tt := map[string]struct {
		id     interface{}
		layout domain.FieldGroups
		args   []interface{}
		want   interface{}
	}{
		"Success": {
			id:     1,
			layout: fg,
			args:   nil,
			want:   fg,
		},
		"Error": {
			id:     1,
			layout: nil,
			args:   nil,
			want:   f,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(nil)
			s.layout = test.layout

			t.Equal(test.want, s.GetLayouts(test.args...))
		})
	}
}

func (t *FieldTestSuite) TestService_HandleLayoutArgs() {
	var f domain.FieldGroups

	tt := map[string]struct {
		layout domain.FieldGroups
		args   []interface{}
		mock   func(p *mocks.PostsRepository)
		want   interface{}
	}{
		"Default": {
			layout: domain.FieldGroups{
				{Title: "test1", Fields: domain.Fields{{Name: "key1"}, {Name: "key2"}}},
			},
			args: nil,
			want: domain.FieldGroups{
				{Title: "test1", Fields: domain.Fields{{Name: "key1"}, {Name: "key2"}}},
			},
		},
		"1 Args (Post)": {
			layout: nil,
			args:   []interface{}{1},
			mock: func(p *mocks.PostsRepository) {
				p.On("GetByID", 1, true).Return(domain.PostDatum{
					Post: domain.Post{Id: 1, Title: "post"},
					Layout: domain.FieldGroups{
						{Title: "test1", Fields: domain.Fields{{Name: "key1"}, {Name: "key2"}}},
					},
				}, nil)
			},
			want: domain.FieldGroups{
				{Title: "test1", Fields: domain.Fields{{Name: "key1"}, {Name: "key2"}}},
			},
		},
		"1 Args (Post Error)": {
			layout: nil,
			args:   []interface{}{1},
			mock: func(p *mocks.PostsRepository) {
				p.On("GetByID", 1, true).Return(domain.PostDatum{}, fmt.Errorf("error"))
			},
			want: f,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetPostsMockService(nil, test.mock)
			s.layout = test.layout
			t.Equal(test.want, s.handleLayoutArgs(test.args))
		})
	}
}

func (t *FieldTestSuite) TestService_GetLayoutsByPost() {
	tt := map[string]struct {
		id   interface{}
		mock func(p *mocks.PostsRepository)
		want domain.FieldGroups
	}{
		"Success": {
			id: 1,
			mock: func(p *mocks.PostsRepository) {
				p.On("GetByID", 1, true).Return(domain.PostDatum{
					Post:   domain.Post{Id: 1, Title: "post"},
					Layout: domain.FieldGroups{{Title: "test"}},
				}, nil)
			},
			want: domain.FieldGroups{{Title: "test"}},
		},
		"Cast Error": {
			id:   noStringer{},
			want: nil,
		},
		"Not Found": {
			id: 1,
			mock: func(p *mocks.PostsRepository) {
				p.On("GetByID", 1, true).Return(domain.PostDatum{}, fmt.Errorf("error"))
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetPostsMockService(nil, test.mock).getLayoutByPost(test.id))
		})
	}
}
