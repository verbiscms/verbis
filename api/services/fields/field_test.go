// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/stretchr/testify/mock"
	"github.com/verbiscms/verbis/api/domain"
	cache "github.com/verbiscms/verbis/api/mocks/cache"
	categories "github.com/verbiscms/verbis/api/mocks/store/categories"
	fields "github.com/verbiscms/verbis/api/mocks/store/fields"
)

func (t *FieldTestSuite) TestService_GetField() {
	tt := map[string]struct {
		fields domain.PostFields
		key    string
		mock   func(f *fields.Repository, c *categories.Repository, ca *cache.Store)
		args   []interface{}
		want   interface{}
		err    bool
	}{
		"Success": {
			fields: domain.PostFields{
				{Type: "text", Name: "key1", OriginalValue: "test"},
			},
			key: "key1",
			mock: func(f *fields.Repository, c *categories.Repository, ca *cache.Store) {
				CacheFieldError(ca)
			},
			args: nil,
			want: "test",
		},
		"From Cache": {
			fields: domain.PostFields{
				{Type: "text", Name: "key1", OriginalValue: "test"},
			},
			key: "key1",
			mock: func(f *fields.Repository, c *categories.Repository, ca *cache.Store) {
				ca.On("Get", mock.Anything, "field-0-0-key1-"+standardCacheKey).
					Return("test", nil)
			},
			args: nil,
			want: "test",
		},
		"No Field": {
			fields: nil,
			key:    "wrongval",
			mock: func(f *fields.Repository, c *categories.Repository, ca *cache.Store) {
				CacheFieldError(ca)
			},
			args: nil,
			want: "",
			err:  true,
		},
		"Post": {
			fields: domain.PostFields{
				{Type: "text", Name: "key1", OriginalValue: "test"},
			},
			key: "key2",
			mock: func(f *fields.Repository, c *categories.Repository, ca *cache.Store) {
				CacheFieldError(ca)
				f.On("Find", 2).Return(domain.PostFields{{Type: "text", Name: "key2", OriginalValue: "test"}}, nil)
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
				t.Contains(t.LogWriter.String(), test.want)
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
		mock   func(f *fields.Repository, c *categories.Repository, ca *cache.Store)
		args   []interface{}
		want   interface{}
		err    bool
	}{
		"Success": {
			fields: domain.PostFields{
				{Type: "text", Name: "key1", OriginalValue: "test"},
			},
			key:  "key1",
			mock: nil,
			args: nil,
			want: domain.PostField{Type: "text", Name: "key1", OriginalValue: "test", Value: "test"},
			err:  false,
		},
		"No Field": {
			fields: nil,
			key:    "wrongval",
			mock:   nil,
			args:   nil,
			want:   "no field exists with the name: wrongval",
			err:    true,
		},
		"post": {
			fields: domain.PostFields{
				{Type: "text", Name: "key1"},
			},
			key: "key2",
			mock: func(f *fields.Repository, c *categories.Repository, ca *cache.Store) {
				f.On("Find", 2).Return(domain.PostFields{{Type: "text", Name: "key2", OriginalValue: "test"}}, nil)
			},
			args: []interface{}{2},
			want: domain.PostField{Type: "text", Name: "key2", OriginalValue: "test", Value: "test"},
			err:  false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetMockService(test.fields, test.mock)

			got := s.GetFieldObject(test.key, test.args...)
			if test.err {
				t.Contains(t.LogWriter.String(), test.want)
				t.Reset()
				return
			}

			t.Equal(test.want, got)
		})
	}
}
