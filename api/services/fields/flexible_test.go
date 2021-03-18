// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

func (t *FieldTestSuite) TestService_GetFlexible() {
	tt := map[string]struct {
		fields domain.PostFields
		input  interface{}
		want   interface{}
		err    bool
	}{
		"Cast to Flexible": {
			fields: nil,
			input: Flexible{
				{
					Name: "layout1",
					SubFields: SubFields{
						{Type: "text", Name: "text1", OriginalValue: "text1", Value: "text1", Key: "flex|0|text1"},
						{Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "flex|0|text2"},
					},
				},
			},
			want: Flexible{
				{
					Name: "layout1",
					SubFields: SubFields{
						{Type: "text", Name: "text1", OriginalValue: "text1", Value: "text1", Key: "flex|0|text1"},
						{Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "flex|0|text2"},
					},
				},
			},
			err: false,
		},
		"No Stringer": {
			fields: nil,
			input:  noStringer{},
			want:   "unable to cast fields.noStringer{} of type fields.noStringer to string",
			err:    true,
		},
		"No Field": {
			fields: nil,
			input:  "test",
			want:   "",
			err:    true,
		},
		"Wrong Field Type": {
			fields: domain.PostFields{
				{Id: 1, Type: "text", Name: "test", OriginalValue: "text", Key: ""},
			},
			input: "test",
			want:  "field with the name: test, is not flexible content",
			err:   true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(test.fields)

			got := s.GetFlexible(test.input)
			if test.err {
				t.Contains(t.logWriter.String(), test.want)
				t.Reset()
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *FieldTestSuite) TestService_ResolveFlexible() {
	tt := map[string]struct {
		flexible domain.PostField
		fields   domain.PostFields
		key      string
		want     interface{}
	}{
		"One Layout": {
			flexible: domain.PostField{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "layout1"},
			fields: domain.PostFields{
				{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "layout1,layout2,layout3"},
				{Type: "text", Name: "text1", OriginalValue: "text1", Key: "flex|0|text1"},
				{Type: "text", Name: "text2", OriginalValue: "text2", Key: "flex|0|text2"},
			},
			key: "flex",
			want: Flexible{
				{
					Name: "layout1",
					SubFields: SubFields{
						{Type: "text", Name: "text1", OriginalValue: "text1", Value: "text1", Key: "flex|0|text1"},
						{Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "flex|0|text2"},
					},
				},
			},
		},
		"Simple": {
			flexible: domain.PostField{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "layout1,layout2,layout3"},
			fields: domain.PostFields{
				{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "layout1,layout2,layout3"},
				{Type: "text", Name: "text1", OriginalValue: "text1", Key: "flex|0|text1"},
				{Type: "text", Name: "text2", OriginalValue: "text2", Key: "flex|0|text2"},
				{Type: "text", Name: "text3", OriginalValue: "text3", Key: "flex|1|text3"},
				{Type: "text", Name: "text4", OriginalValue: "text4", Key: "flex|1|text4"},
				{Type: "text", Name: "text5", OriginalValue: "text5", Key: "flex|2|text5"},
				{Type: "text", Name: "text6", OriginalValue: "text6", Key: "flex|2|text6"},
			},
			key: "flex",
			want: Flexible{
				{
					Name: "layout1",
					SubFields: SubFields{
						{Type: "text", Name: "text1", OriginalValue: "text1", Value: "text1", Key: "flex|0|text1"},
						{Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "flex|0|text2"},
					},
				},
				{
					Name: "layout2",
					SubFields: SubFields{
						{Type: "text", Name: "text3", OriginalValue: "text3", Value: "text3", Key: "flex|1|text3"},
						{Type: "text", Name: "text4", OriginalValue: "text4", Value: "text4", Key: "flex|1|text4"},
					},
				},
				{
					Name: "layout3",
					SubFields: SubFields{
						{Type: "text", Name: "text5", OriginalValue: "text5", Value: "text5", Key: "flex|2|text5"},
						{Type: "text", Name: "text6", OriginalValue: "text6", Value: "text6", Key: "flex|2|text6"},
					},
				},
			},
		},
		"Nested": {
			flexible: domain.PostField{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "layout1"},
			fields: domain.PostFields{
				{Type: "text", Name: "text1", OriginalValue: "text1", Key: "flex|0|text1"},
				{Type: "flexible", Name: "nested", OriginalValue: "nestedlayout", Key: "flex|0|nested"},
				{Type: "text", Name: "text2", OriginalValue: "text2", Key: "flex|0|nested|0|text1"},
			},
			key: "flex",
			want: Flexible{
				{
					Name: "layout1",
					SubFields: SubFields{
						{Type: "text", Name: "text1", OriginalValue: "text1", Value: "text1", Key: "flex|0|text1"},
						{Type: "flexible", Name: "nested", OriginalValue: "nestedlayout", Key: "flex|0|nested", Value: Flexible{
							{
								Name: "nestedlayout",
								SubFields: SubFields{
									{Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "flex|0|nested|0|text1"},
								},
							},
						}},
					},
				},
			},
		},
		"Repeater": {
			flexible: domain.PostField{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "layout1,layout1"},
			fields: domain.PostFields{
				{Id: 1, Type: "flexible", Name: "flex", OriginalValue: "layout1,layout2"},
				{Type: "text", Name: "text1", OriginalValue: "text1", Key: "flex|0|text1"},
				{Type: "text", Name: "text2", OriginalValue: "text2", Key: "flex|0|text2"},
				{Type: "text", Name: "text3", OriginalValue: "text3", Key: "flex|1|text3"},
				{Type: "text", Name: "text4", OriginalValue: "text4", Key: "flex|1|text4"},

				{Type: "repeater", Name: "repeater", OriginalValue: "1", Key: "flex|0|repeater"},
				{Type: "text", Name: "text", OriginalValue: "text1", Key: "flex|0|repeater|0|text"},
				{Type: "text", Name: "text2", OriginalValue: "text2", Key: "flex|0|repeater|0|text2"},

				{Type: "repeater", Name: "repeater", OriginalValue: "1", Key: "flex|1|repeater"},
				{Type: "text", Name: "text", OriginalValue: "text1", Key: "flex|1|repeater|0|text"},
				{Type: "text", Name: "text2", OriginalValue: "text2", Key: "flex|1|repeater|0|text2"},
			},
			key: "flex",
			want: Flexible{
				{
					Name: "layout1",
					SubFields: SubFields{
						{Type: "text", Name: "text1", OriginalValue: "text1", Value: "text1", Key: "flex|0|text1"},
						{Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "flex|0|text2"},
						{Type: "repeater", Name: "repeater", OriginalValue: "1", Key: "flex|0|repeater", Value: Repeater{
							Row{
								{Type: "text", Name: "text", OriginalValue: "text1", Value: "text1", Key: "flex|0|repeater|0|text"},
								{Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "flex|0|repeater|0|text2"},
							}},
						},
					},
				},
				{
					Name: "layout1",
					SubFields: SubFields{
						{Type: "text", Name: "text3", OriginalValue: "text3", Value: "text3", Key: "flex|1|text3"},
						{Type: "text", Name: "text4", OriginalValue: "text4", Value: "text4", Key: "flex|1|text4"},
						{Type: "repeater", Name: "repeater", OriginalValue: "1", Key: "flex|1|repeater", Value: Repeater{
							Row{
								{Type: "text", Name: "text", OriginalValue: "text1", Value: "text1", Key: "flex|1|repeater|0|text"},
								{Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "flex|1|repeater|0|text2"},
							}},
						},
					},
				},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, t.GetService(test.fields).resolveFlexible("", test.flexible, test.fields))
		})
	}
}

func (t *FieldTestSuite) TestFlexible_HasRows() {
	tt := map[string]struct {
		flexible Flexible
		want     interface{}
	}{
		"With Rows": {
			flexible: Flexible{
				{Name: "layout", SubFields: SubFields{domain.PostField{Id: 1, Name: "test"}}},
			},
			want: true,
		},
		"Without Rows": {
			flexible: Flexible{},
			want:     false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, test.flexible.HasRows())
		})
	}
}

func (t *FieldTestSuite) TestSubFields_SubField() {
	subfield := SubFields{
		{Id: 1, Name: "test1", Value: 1},
		{Id: 2, Name: "test2", Value: 2},
		{Id: 3, Name: "test3", Value: 3},
	}

	tt := map[string]struct {
		key  string
		want interface{}
	}{
		"Found": {
			key:  "test1",
			want: 1,
		},
		"Not Found": {
			key:  "wrongval",
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, subfield.SubField(test.key))
		})
	}
}

func (t *FieldTestSuite) TestSubFields_First() {
	tt := map[string]struct {
		subfields SubFields
		want      interface{}
	}{
		"Found": {
			subfields: SubFields{
				{Id: 1, Name: "test1", Value: 1},
				{Id: 2, Name: "test2", Value: 2},
				{Id: 3, Name: "test3", Value: 3},
			},
			want: domain.PostField{Id: 1, Name: "test1", Value: 1},
		},
		"Not Found": {
			subfields: SubFields{},
			want:      nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, test.subfields.First())
		})
	}
}

func (t *FieldTestSuite) TestSubFields_Last() {
	tt := map[string]struct {
		subfields SubFields
		want      interface{}
	}{
		"Found": {
			subfields: SubFields{
				{Id: 1, Name: "test1", Value: 1},
				{Id: 2, Name: "test2", Value: 2},
				{Id: 3, Name: "test3", Value: 3},
			},
			want: domain.PostField{Id: 3, Name: "test3", Value: 3},
		},
		"Not Found": {
			subfields: SubFields{},
			want:      nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, test.subfields.Last())
		})
	}
}
