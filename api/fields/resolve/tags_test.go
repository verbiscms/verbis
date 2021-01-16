package resolve

import "github.com/ainsleyclark/verbis/api/domain"

func (t *ResolverTestSuite) TestFieldValue_Tags() {

	tt := map[string]struct {
		value  domain.FieldValue
		want   tags
	}{
		"Success": {
			value: "1,2,3,4,5",
			want: tags{"1","2","3","4","5"},
		},
		"Trailing Comma": {
			value: "1,2,3,4,5,",
			want: tags{"1","2","3","4","5"},
		},
		"Leading Commas": {
			value: ",1,2,3,4,5",
			want: tags{"1","2","3","4","5"},
		},
		"Commas Everywhere": {
			value: ",,,,1,,,,2,,3,,4,,,,5,,,,,",
			want: tags{"1","2","3","4","5"},
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