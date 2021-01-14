package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
)


func (t *FieldTestSuite) TestService_GetRepeater() {

	tt := map[string]struct {
		fields []domain.PostField
		input  interface{}
		want   interface{}
	}{
		"Cast to Repeater": {
			fields: nil,
			input: Repeater{
				Row{{Id: 1, Type: "text", Name: "text", OriginalValue: "text1", Value: "text1", Key: "repeater|0|text"}},
				Row{{Id: 2, Type: "text", Name: "text", OriginalValue: "text2", Value: "text2", Key: "repeater|1|text"}},
			},
			want: Repeater{
				Row{{Id: 1, Type: "text", Name: "text", OriginalValue: "text1", Value: "text1", Key: "repeater|0|text"}},
				Row{{Id: 2, Type: "text", Name: "text", OriginalValue: "text2", Value: "text2", Key: "repeater|1|text"}},
			},
		},
		"No Stringer": {
			fields: nil,
			input: noStringer{},
			want: "unable to cast fields.noStringer{} of type fields.noStringer to string",
		},
		"No Field": {
			fields: nil,
			input: "test",
			want: "no field exists with the name: test",
		},
		"Wrong Field Type": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "test", OriginalValue: "text", Key: ""},
			},
			input: "test",
			want: "field with the name: test, is not a repeater",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(test.fields)

			got, err := s.GetRepeater(test.input)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *FieldTestSuite) TestService_Repeater() {

	tt := map[string]struct {
		fields []domain.PostField
		key    string
		want   interface{}
	}{
		"Simple": {
			fields: []domain.PostField{
				{Id: 1, Type: "repeater", Name: "repeater", OriginalValue: "3"},
				{Id: 2, Type: "text", Name: "text", OriginalValue: "text1", Key: "repeater|0|text"},
				{Id: 3, Type: "text", Name: "text2", OriginalValue: "text2", Key: "repeater|0|text2"},
				{Id: 4, Type: "text", Name: "text", OriginalValue: "text3", Key: "repeater|1|text"},
				{Id: 5, Type: "text", Name: "text2", OriginalValue: "text4", Key: "repeater|1|text2"},
				{Id: 6, Type: "text", Name: "text", OriginalValue: "text5", Key: "repeater|2|text"},
				{Id: 7, Type: "text", Name: "text2", OriginalValue: "text6", Key: "repeater|2|text2"},
			},
			key: "repeater",
			want: Repeater{
				Row{
					{Id: 2, Type: "text", Name: "text", OriginalValue: "text1", Value: "text1", Key: "repeater|0|text"},
					{Id: 3, Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "repeater|0|text2"},
				},
				Row{
					{Id: 4, Type: "text", Name: "text", OriginalValue: "text3", Value: "text3", Key: "repeater|1|text"},
					{Id: 5, Type: "text", Name: "text2", OriginalValue: "text4", Value: "text4", Key: "repeater|1|text2"},
				},
				Row{
					{Id: 6, Type: "text", Name: "text", OriginalValue: "text5", Value: "text5", Key: "repeater|2|text"},
					{Id: 7, Type: "text", Name: "text2", OriginalValue: "text6", Value: "text6", Key: "repeater|2|text2"},
				},
			},
		},
		"Nested": {
			fields: []domain.PostField{
				{Id: 1, Type: "repeater", Name: "repeater", OriginalValue: "2"},
				{Id: 4, Type: "text", Name: "parent_text", OriginalValue: "R1", Key: "repeater|0|parent_text"},
				{Id: 5, Type: "text", Name: "parent_text", OriginalValue: "R2", Key: "repeater|1|parent_text"},
				{Id: 2, Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|0|nested"},
				{Id: 6, Type: "text", Name: "nested_text", OriginalValue: "N1", Key: "repeater|0|nested|0|nested_test"},
				{Id: 7, Type: "text", Name: "nested_text", OriginalValue: "N2", Key: "repeater|0|nested|1|nested_test"},
				{Id: 3, Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|1|nested"},
				{Id: 8, Type: "text", Name: "nested_text", OriginalValue: "N3", Key: "repeater|1|nested|0|nested_test"},
				{Id: 9, Type: "text", Name: "nested_text", OriginalValue: "N4", Key: "repeater|1|nested|1|nested_test"},
			},
			key: "repeater",
			want: Repeater{
				Row{
					{Id: 4, Type: "text", Name: "parent_text", OriginalValue: "R1", Value: "R1", Key: "repeater|0|parent_text"},
					{Id: 2, Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|0|nested", Value: Repeater{
						Row{{Id: 6, Type: "text", Name: "nested_text", OriginalValue: "N1", Value: "N1", Key: "repeater|0|nested|0|nested_test"}},
						Row{{Id: 7, Type: "text", Name: "nested_text", OriginalValue: "N2", Value: "N2", Key: "repeater|0|nested|1|nested_test"}},
					}},
				},
				Row{
					{Id: 5, Type: "text", Name: "parent_text", OriginalValue: "R2", Value: "R2", Key: "repeater|1|parent_text"},
					{Id: 3, Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|1|nested", Value: Repeater{
						Row{{Id: 8, Type: "text", Name: "nested_text", OriginalValue: "N3", Value: "N3", Key: "repeater|1|nested|0|nested_test"}},
						Row{{Id: 9, Type: "text", Name: "nested_text", OriginalValue: "N4", Value: "N4", Key: "repeater|1|nested|1|nested_test"}},
					}},
				},
			},
		},
		"Nested Nested :(": {
			fields: []domain.PostField{
				{Id: 1, Type: "repeater", Name: "repeater", OriginalValue: "2"},
				{Id: 4, Type: "text", Name: "parent_text", OriginalValue: "R1", Key: "repeater|0|parent_text"},
				{Id: 5, Type: "text", Name: "parent_text", OriginalValue: "R2", Key: "repeater|1|parent_text"},

				{Id: 2, Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|0|nested"},
				{Id: 6, Type: "text", Name: "nested_text", OriginalValue: "N1", Key: "repeater|0|nested|0|nested_test"},
				{Id: 7, Type: "text", Name: "nested_text", OriginalValue: "N2", Key: "repeater|0|nested|1|nested_test"},

				{Id: 3, Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|1|nested"},
				{Id: 8, Type: "text", Name: "nested_text", OriginalValue: "N3", Key: "repeater|1|nested|0|nested_test"},
				{Id: 9, Type: "text", Name: "nested_text", OriginalValue: "N4", Key: "repeater|1|nested|1|nested_test"},
			},
			key: "repeater",
			want: Repeater{
				Row{
					{Id: 4, Type: "text", Name: "parent_text", OriginalValue: "R1", Value: "R1", Key: "repeater|0|parent_text"},
					{Id: 2, Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|0|nested", Value: Repeater{
						Row{{Id: 6, Type: "text", Name: "nested_text", OriginalValue: "N1", Value: "N1", Key: "repeater|0|nested|0|nested_test"}},
						Row{{Id: 7, Type: "text", Name: "nested_text", OriginalValue: "N2", Value: "N2", Key: "repeater|0|nested|1|nested_test"}},
					}},
				},
				Row{
					{Id: 5, Type: "text", Name: "parent_text", OriginalValue: "R2", Value: "R2", Key: "repeater|1|parent_text"},
					{Id: 3, Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|1|nested", Value: Repeater{
						Row{{Id: 8, Type: "text", Name: "nested_text", OriginalValue: "N3", Value: "N3", Key: "repeater|1|nested|0|nested_test"}},
						Row{{Id: 9, Type: "text", Name: "nested_text", OriginalValue: "N4", Value: "N4", Key: "repeater|1|nested|1|nested_test"}},
					}},
				},
			},
		},
	}


	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(test.fields)

			got, err := s.GetRepeater(test.key)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}


func (t *FieldTestSuite) TestRepeater_HasRows() {

	tt := map[string]struct {
		repeater Repeater
		want     interface{}
	}{
		"With Rows": {
			repeater: Repeater{
				Row{
					{Id: 1}, {Id: 2}, {Id: 3},
				},
			},
			want: true,
		},
		"Without Rows": {
			repeater: Repeater{},
			want:     false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, test.repeater.HasRows())
		})
	}
}

func (t *FieldTestSuite) TestRepeater_SubField() {

	row := Row{
		{Id: 1, Name: "test1", Type: "text", OriginalValue: "1", Value: "1"},
		{Id: 2, Name: "test2", Type: "text", OriginalValue: "2", Value: "2"},
		{Id: 3, Name: "test3", Type: "text", OriginalValue: "3", Value: "3"},
	}

	tt := map[string]struct {
		key  string
		want interface{}
	}{
		"Found": {
			key:  "test1",
			want: "1",
		},
		"Found 2": {
			key:  "test2",
			want: "2",
		},
		"Found 3": {
			key:  "test3",
			want: "3",
		},
		"Not Found": {
			key:  "wrongval",
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, row.SubField(test.key))
		})
	}
}

func (t *FieldTestSuite) TestRow_First() {

	tt := map[string]struct {
		row Row
		want     interface{}
	}{
		"Found": {
			row: Row{
				{Id: 1, Name: "test1", Type: "text", OriginalValue: "1", Value: "1"},
				{Id: 2, Name: "test2", Type: "text", OriginalValue: "2", Value: "2"},
				{Id: 3, Name: "test3", Type: "text", OriginalValue: "3", Value: "3"},
			},
			want: domain.PostField{Id: 1, Name: "test1", Type: "text", OriginalValue: "1", Value: "1"},
		},
		"Not Found": {
			row: Row{},
			want:     nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, test.row.First())
		})
	}
}

func (t *FieldTestSuite) TestRow_Last() {

	tt := map[string]struct {
		row Row
		want     interface{}
	}{
		"Found": {
			row: Row{
				{Id: 1, Name: "test1", Type: "text", OriginalValue: "1", Value: "1"},
				{Id: 2, Name: "test2", Type: "text", OriginalValue: "2", Value: "2"},
				{Id: 3, Name: "test3", Type: "text", OriginalValue: "3", Value: "3"},
			},
			want: domain.PostField{Id: 3, Name: "test3", Type: "text", OriginalValue: "3", Value: "3"},
		},
		"Not Found": {
			row: Row{},
			want:     nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, test.row.Last())
		})
	}
}
