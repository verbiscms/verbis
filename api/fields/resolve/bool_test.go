package resolve

import "github.com/ainsleyclark/verbis/api/domain"

func (t *ResolverTestSuite) TestFieldValue_Checkbox() {

	tt := map[string]struct {
		value  domain.FieldValue
		want   interface{}
	}{
		"True": {
			value: domain.FieldValue("true"),
			want: true,
		},
		"False": {
			value: domain.FieldValue("false"),
			want: false,
		},
		"Failed": {
			value: `wrongval`,
			want: "invalid syntax",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()

			got, err := v.checkbox(test.value)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}