package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

func (t *FieldTestSuite) TestService_GetRepeater() {

	tt := map[string]struct {
		fields []domain.PostField
		input  interface{}
		want   interface{}
		err    bool
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
			want:   "no field exists with the name: test",
			err:    true,
		},
		"Wrong Field Type": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "test", OriginalValue: "text", Key: ""},
			},
			input: "test",
			want:  "field with the name: test, is not a repeater",
			err:   true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(test.fields)

			got := s.GetRepeater(test.input)
			if test.err {
				t.Contains(t.logWriter.String(), test.want)
				t.Reset()
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *FieldTestSuite) TestService_ResolveRepeater() {

	tt := map[string]struct {
		fields []domain.PostField
		key    string
		want   interface{}
	}{
		"Bad Cast to Int": {
			fields: []domain.PostField{
				{Id: 1, Type: "repeater", Name: "repeater", OriginalValue: "@£$$%^&%$^&"},
				{Id: 2, Type: "text", Name: "text", OriginalValue: "text1", Key: "repeater|0|text"},
			},
			key:  "repeater",
			want: Repeater{},
		},
		"Simple": {
			fields: []domain.PostField{
				{Type: "repeater", Name: "repeater", OriginalValue: "3"},
				{Type: "text", Name: "text", OriginalValue: "text1", Key: "repeater|0|text"},
				{Type: "text", Name: "text2", OriginalValue: "text2", Key: "repeater|0|text2"},
				{Type: "text", Name: "text", OriginalValue: "text3", Key: "repeater|1|text"},
				{Type: "text", Name: "text2", OriginalValue: "text4", Key: "repeater|1|text2"},
				{Type: "text", Name: "text", OriginalValue: "text5", Key: "repeater|2|text"},
				{Type: "text", Name: "text2", OriginalValue: "text6", Key: "repeater|2|text2"},
			},
			key: "repeater",
			want: Repeater{
				Row{
					{Type: "text", Name: "text", OriginalValue: "text1", Value: "text1", Key: "repeater|0|text"},
					{Type: "text", Name: "text2", OriginalValue: "text2", Value: "text2", Key: "repeater|0|text2"},
				},
				Row{
					{Type: "text", Name: "text", OriginalValue: "text3", Value: "text3", Key: "repeater|1|text"},
					{Type: "text", Name: "text2", OriginalValue: "text4", Value: "text4", Key: "repeater|1|text2"},
				},
				Row{
					{Type: "text", Name: "text", OriginalValue: "text5", Value: "text5", Key: "repeater|2|text"},
					{Type: "text", Name: "text2", OriginalValue: "text6", Value: "text6", Key: "repeater|2|text2"},
				},
			},
		},
		"Nested": {
			fields: []domain.PostField{
				{Type: "repeater", Name: "repeater", OriginalValue: "2"},
				{Type: "text", Name: "parent_text", OriginalValue: "R1", Key: "repeater|0|parent_text"},
				{Type: "text", Name: "parent_text", OriginalValue: "R2", Key: "repeater|1|parent_text"},
				{Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|0|nested"},
				{Type: "text", Name: "nested_text", OriginalValue: "N1", Key: "repeater|0|nested|0|nested_test"},
				{Type: "text", Name: "nested_text", OriginalValue: "N2", Key: "repeater|0|nested|1|nested_test"},
				{Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|1|nested"},
				{Type: "text", Name: "nested_text", OriginalValue: "N3", Key: "repeater|1|nested|0|nested_test"},
				{Type: "text", Name: "nested_text", OriginalValue: "N4", Key: "repeater|1|nested|1|nested_test"},
			},
			key: "repeater",
			want: Repeater{
				Row{
					{Type: "text", Name: "parent_text", OriginalValue: "R1", Value: "R1", Key: "repeater|0|parent_text"},
					{Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|0|nested", Value: Repeater{
						Row{{Type: "text", Name: "nested_text", OriginalValue: "N1", Value: "N1", Key: "repeater|0|nested|0|nested_test"}},
						Row{{Type: "text", Name: "nested_text", OriginalValue: "N2", Value: "N2", Key: "repeater|0|nested|1|nested_test"}},
					}},
				},
				Row{
					{Type: "text", Name: "parent_text", OriginalValue: "R2", Value: "R2", Key: "repeater|1|parent_text"},
					{Type: "repeater", Name: "nested", OriginalValue: "2", Key: "repeater|1|nested", Value: Repeater{
						Row{{Type: "text", Name: "nested_text", OriginalValue: "N3", Value: "N3", Key: "repeater|1|nested|0|nested_test"}},
						Row{{Type: "text", Name: "nested_text", OriginalValue: "N4", Value: "N4", Key: "repeater|1|nested|1|nested_test"}},
					}},
				},
			},
		},
		"Nested Nested": {
			fields: []domain.PostField{
				{Type: "repeater", Name: "repeater", OriginalValue: "2"},
				{Type: "text", Name: "parent_text", OriginalValue: "R1", Key: "repeater|0|parent_text"},
				{Type: "text", Name: "parent_text", OriginalValue: "R2", Key: "repeater|1|parent_text"},
				{Type: "repeater", Name: "nested", OriginalValue: "1", Key: "repeater|0|nested"},
				{Type: "text", Name: "nested_text", OriginalValue: "N1", Key: "repeater|0|nested|0|nested_test"},

				{Type: "repeater", Name: "nested_nested", OriginalValue: "1", Key: "repeater|0|nested|0|nested_nested"},
				{Type: "text", Name: "nested_nested_text", OriginalValue: "NN1", Key: "repeater|0|nested|0|nested_nested|0|nested_nested_text"},

				{Type: "repeater", Name: "nested", OriginalValue: "1", Key: "repeater|1|nested"},
				{Type: "text", Name: "nested_text", OriginalValue: "N2", Key: "repeater|1|nested|0|nested_test"},

				{Type: "repeater", Name: "nested_nested", OriginalValue: "1", Key: "repeater|1|nested|0|nested_nested"},
				{Type: "text", Name: "nested_nested_text", OriginalValue: "NN1", Key: "repeater|1|nested|0|nested_nested|0|nested_nested_text"},
			},
			key: "repeater",
			want: Repeater{
				Row{
					{Type: "text", Name: "parent_text", OriginalValue: "R1", Value: "R1", Key: "repeater|0|parent_text"},
					{Type: "repeater", Name: "nested", OriginalValue: "1", Key: "repeater|0|nested", Value: Repeater{
						Row{
							{Type: "text", Name: "nested_text", OriginalValue: "N1", Value: "N1", Key: "repeater|0|nested|0|nested_test"},
							{Type: "repeater", Name: "nested_nested", OriginalValue: "1", Key: "repeater|0|nested|0|nested_nested", Value: Repeater{
								Row{
									{Type: "text", Name: "nested_nested_text", OriginalValue: "NN1", Value: "NN1", Key: "repeater|0|nested|0|nested_nested|0|nested_nested_text"},
								},
							}},
						},
					}},
				},
				Row{
					{Type: "text", Name: "parent_text", OriginalValue: "R2", Value: "R2", Key: "repeater|1|parent_text"},
					{Type: "repeater", Name: "nested", OriginalValue: "1", Key: "repeater|1|nested", Value: Repeater{
						Row{
							{Type: "text", Name: "nested_text", OriginalValue: "N2", Value: "N2", Key: "repeater|1|nested|0|nested_test"},
							{Type: "repeater", Name: "nested_nested", OriginalValue: "1", Key: "repeater|1|nested|0|nested_nested", Value: Repeater{
								Row{
									{Type: "text", Name: "nested_nested_text", OriginalValue: "NN1", Value: "NN1", Key: "repeater|1|nested|0|nested_nested|0|nested_nested_text"},
								},
							}},
						},
					}},
				},
			},
		},

		// TODO FLEXIBLE
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(test.fields)

			got := s.GetRepeater(test.key)
			//if err != nil {
			//	t.Contains(err.Error(), test.want)
			//	return
			//}

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

func (t *FieldTestSuite) TestRow_SubField() {

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

func (t *FieldTestSuite) TestRow_HasField() {

	row := Row{
		{Id: 1, Name: "test1", Type: "text", OriginalValue: "1", Value: "1"},
		{Id: 2, Name: "test2", Type: "text", OriginalValue: "2", Value: "2"},
		{Id: 3, Name: "test3", Type: "text", OriginalValue: "3", Value: "3"},
	}

	tt := map[string]struct {
		key  string
		want interface{}
	}{
		"True": {
			key:  "test1",
			want: true,
		},
		"False": {
			key:  "wrongval",
			want: false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, row.HasField(test.key))
		})
	}
}

func (t *FieldTestSuite) TestRow_First() {

	tt := map[string]struct {
		row  Row
		want interface{}
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
			row:  Row{},
			want: nil,
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
		row  Row
		want interface{}
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
			row:  Row{},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, test.row.Last())
		})
	}
}
