// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *FieldTestSuite) TestService_HandleArgs() {
	tt := map[string]struct {
		fields domain.PostFields
		args   []interface{}
		mock   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		want   domain.PostFields
	}{
		"Default": {
			fields: domain.PostFields{{Name: "test"}},
			args:   nil,
			want:   domain.PostFields{{Name: "test"}},
		},
		"1 Args (Post Fields)": {
			args: []interface{}{domain.PostDatum{
				Post:   domain.Post{Id: 1, Title: "post"},
				Fields: domain.PostFields{{Id: 1, Type: "text", Name: "post"}},
			}},
			want: domain.PostFields{{Id: 1, Type: "text", Name: "post"}},
		},
		"1 Args (Post)": {
			fields: nil,
			args:   []interface{}{1},
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return(domain.PostFields{
					{Id: 1, Type: "text", Name: "post"},
				}, nil)
			},
			want: domain.PostFields{{Id: 1, Type: "text", Name: "post"}},
		},
		"1 Args (Post Template)": {
			fields: nil,
			args: []interface{}{domain.PostTemplate{
				Post:   domain.Post{Id: 1, Title: "post"},
				Fields: domain.PostFields{{Id: 1, Type: "text", Name: "post"}},
			}},
			mock: nil,
			want: domain.PostFields{{Id: 1, Type: "text", Name: "post"}},
		},
		"1 Args (Fields)": {
			fields: nil,
			args:   []interface{}{domain.PostFields{{Id: 1, Type: "text", Name: "post"}}},
			mock:   nil,
			want:   domain.PostFields{{Id: 1, Type: "text", Name: "post"}},
		},
		"1 Args (Post Error)": {
			fields: domain.PostFields{{Name: "test"}},
			args:   []interface{}{1},
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return(nil, fmt.Errorf("error"))
			},
			want: nil,
		},
		"Cast Error": {
			fields: nil,
			args:   []interface{}{noStringer{}},
			mock:   nil,
			want:   nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetMockService(test.fields, test.mock).handleArgs(test.args))
		})
	}
}

func (t *FieldTestSuite) TestService_GetFieldsByPost() {
	tt := map[string]struct {
		id   int
		mock func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		want domain.PostFields
	}{
		"Success": {
			id: 1,
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return(domain.PostFields{
					{Id: 1, Type: "text", Name: "post"},
				}, nil)
			},
			want: domain.PostFields{{Id: 1, Type: "text", Name: "post"}},
		},
		"Get Error": {
			id: 1,
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 1).Return(nil, fmt.Errorf("error"))
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetMockService(nil, test.mock).getFieldsByPost(test.id))
		})
	}
}

func (t *FieldTestSuite) TestService_FindFieldByName() {

	tt := map[string]struct {
		name   string
		fields domain.PostFields
		want   interface{}
	}{
		"Success": {
			name:   "test",
			fields: domain.PostFields{{Id: 1, Type: "text", Name: "test"}},
			want:   domain.PostField{Id: 1, Type: "text", Name: "test"},
		},
		"Fail": {
			name:   "test",
			fields: nil,
			want:   "no field exists with the name: test",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(test.fields)

			got, err := s.findFieldByName(test.name, test.fields)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

//func (t *FieldTestSuite) TestResolve_Walker() {
//
//	tt := map[string]struct {
//		resolver resolve
//		want     interface{}
//	}{
//		"No Prefix": {
//			resolver: resolve{
//				Key:   "",
//				Index: 0,
//				Field: domain.PostField{Type: "repeater", Name: "repeater", OriginalValue: "1"},
//				Fields: domain.PostFields{
//					{Type: "text", Name: "text", OriginalValue: "text1", Key: ""},
//				},
//			},
//			want: domain.PostField{Type: "text", Name: "text", OriginalValue: "text1", Value: "text1", Key: ""},
//		},
//		"Success": {
//			resolver: resolve{
//				Key:   "",
//				Index: 0,
//				Field: domain.PostField{Type: "repeater", Name: "repeater", OriginalValue: "1"},
//				Fields: domain.PostFields{
//					{Type: "text", Name: "text", OriginalValue: "text1", Key: "repeater|0|text"},
//				},
//			},
//			want: domain.PostField{Type: "text", Name: "text", OriginalValue: "text1", Value: "text1", Key: "repeater|0|text"},
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			test.resolver.Service = t.GetService(nil)
//			test.resolver.Walker(func(field domain.PostField) {
//				t.Equal(test.want, field)
//			})
//		})
//	}
//}
