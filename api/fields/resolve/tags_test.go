package resolve

import "github.com/ainsleyclark/verbis/api/domain"

func (t *ResolverTestSuite) TestValue_Tags() {

	tt := map[string]struct {
		value domain.FieldValue
		want  tags
	}{
		"Success": {
			value: "1,2,3,4,5",
			want:  tags{"1", "2", "3", "4", "5"},
		},
		"Single": {
			value: "1",
			want:  tags{"1"},
		},
		"Trailing Comma": {
			value: "1,2,3,4,5,",
			want:  tags{"1", "2", "3", "4", "5"},
		},
		"Leading Commas": {
			value: ",1,2,3,4,5",
			want:  tags{"1", "2", "3", "4", "5"},
		},
		"Commas Everywhere": {
			value: ",,,,1,,,,2,,3,,4,,,,5,,,,,",
			want:  tags{"1", "2", "3", "4", "5"},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()

			got, err := v.tags(test.value)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *ResolverTestSuite) TestValue_TagsResolve() {

	tt := map[string]struct {
		field domain.PostField
		want  domain.PostField
	}{
		"Tags": {
			field: domain.PostField{OriginalValue: "1,2,3", Type: "tags"},
			want:  domain.PostField{OriginalValue: "1,2,3", Type: "tags", Value: []interface{}{"1","2","3"}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()

			got := v.resolve(test.field)

			t.Equal(test.want, got)
		})
	}
}