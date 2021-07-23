// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/verbiscms/verbis/api/domain"
)

func (t *FieldTestSuite) TestService_GetFields() {
	tt := map[string]struct {
		fields domain.PostFields
		args   []interface{}
		want   interface{}
	}{
		"None": {
			fields: domain.PostFields{},
			args:   nil,
			want:   Fields{},
		},
		"Simple": {
			fields: domain.PostFields{
				{Type: "text", Name: "key1", OriginalValue: "1"},
				{Type: "text", Name: "key2", OriginalValue: "2"},
				{Type: "text", Name: "key3", OriginalValue: "3"},
			},
			args: nil,
			want: Fields{
				"key1": "1",
				"key2": "2",
				"key3": "3",
			},
		},
	}
	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetMockService(test.fields, nil).GetFields(test.args...))
		})
	}
}

func (t *FieldTestSuite) TestService_Mapper() {
	tt := map[string]struct {
		fields domain.PostFields
		want   interface{}
	}{
		"Simple": {
			fields: domain.PostFields{{Type: "text", Name: "key1", Value: "1"}},
			want:   domain.PostField{Type: "text", Name: "key1", Value: "1"},
		},
		"No Separator": {
			fields: domain.PostFields{{Type: "text", Name: "key1", Key: "map", Value: 1}},
			want:   domain.PostField{Type: "text", Name: "key1", Key: "map", Value: 1},
		},
		"Repeater": {
			fields: domain.PostFields{
				{Type: "repeater", Name: "repeater", OriginalValue: "1"},
				{Type: "text", Name: "text", OriginalValue: "text1", Key: "repeater|0|text"},
				{Type: "text", Name: "text2", OriginalValue: "text2", Key: "repeater|0|text2"},
			},
			want: domain.PostField{Type: "repeater", Name: "repeater", OriginalValue: "1", Value: Repeater{
				Row{
					{Type: "text", Name: "text", OriginalValue: "text1", Value: "text1", Key: "repeater|0|text"},
					{Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "repeater|0|text2"},
				}},
			},
		},
		"Flexible": {
			fields: domain.PostFields{
				{Type: "flexible", Name: "flex", OriginalValue: "layout1"},
				{Type: "text", Name: "text1", OriginalValue: "text1", Key: "flex|0|text1"},
				{Type: "text", Name: "text2", OriginalValue: "text2", Key: "flex|0|text2"},
			},
			want: domain.PostField{Type: "flexible", Name: "flex", OriginalValue: "layout1", Value: Flexible{
				{
					Name: "layout1",
					SubFields: SubFields{
						{Type: "text", Name: "text1", OriginalValue: "text1", Value: "text1", Key: "flex|0|text1"},
						{Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "flex|0|text2"},
					},
				},
			}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(test.fields)

			s.mapper(test.fields, func(field domain.PostField) {
				t.Equal(test.want, field)
			})
		})
	}
}
