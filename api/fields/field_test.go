// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *FieldTestSuite) TestService_GetField() {

	tt := map[string]struct {
		fields domain.PostFields
		key    string
		mock   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		args   []interface{}
		want   interface{}
		err    bool
	}{
		"Success": {
			fields: domain.PostFields{
				{Id: 1, Type: "text", Name: "key1", OriginalValue: "test"},
			},
			key:  "key1",
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args: nil,
			want: "test",
		},
		"No Field": {
			fields: nil,
			key:    "wrongval",
			mock:   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args:   nil,
			want:   "",
			err:    true,
		},
		"Post": {
			fields: domain.PostFields{
				{Id: 1, Type: "text", Name: "key1", OriginalValue: "test"},
			},
			key: "key2",
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 2).Return(domain.PostFields{{Id: 2, Type: "text", Name: "key2", OriginalValue: "test"}}, nil)
			},
			args: []interface{}{2},
			want: "test",
			err:  false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetMockService(test.fields, test.mock)

			got := s.GetField(test.key, test.args...)
			if test.err {
				t.Contains(t.logWriter.String(), test.want)
				t.Reset()
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *FieldTestSuite) TestService_GetFieldObject() {

	tt := map[string]struct {
		fields domain.PostFields
		key    string
		mock   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		args   []interface{}
		want   interface{}
		err    bool
	}{
		"Success": {
			fields: domain.PostFields{
				{Id: 1, Type: "text", Name: "key1", OriginalValue: "test"},
			},
			key:  "key1",
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args: nil,
			want: domain.PostField{Id: 1, Type: "text", Name: "key1", OriginalValue: "test", Value: "test"},
			err:  false,
		},
		"No Field": {
			fields: nil,
			key:    "wrongval",
			mock:   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args:   nil,
			want:   "no field exists with the name: wrongval",
			err:    true,
		},
		"Post": {
			fields: domain.PostFields{
				{Id: 1, Type: "text", Name: "key1"},
			},
			key: "key2",
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 2).Return(domain.PostFields{{Id: 2, Type: "text", Name: "key2", OriginalValue: "test"}}, nil)
			},
			args: []interface{}{2},
			want: domain.PostField{Id: 2, Type: "text", Name: "key2", OriginalValue: "test", Value: "test"},
			err:  false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetMockService(test.fields, test.mock)

			got := s.GetFieldObject(test.key, test.args...)
			if test.err {
				t.Contains(t.logWriter.String(), test.want)
				t.Reset()
				return
			}

			t.Equal(test.want, got)

		})
	}
}
