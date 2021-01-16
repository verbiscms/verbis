package resolve

import "github.com/ainsleyclark/verbis/api/domain"

func (t *ResolverTestSuite) TestFieldValue_Choice() {

	tt := map[string]struct {
		value  domain.FieldValue
		want   interface{}
	}{
		"Success": {
			value: domain.FieldValue(`{"key": "key1", "value": "value1"}`),
			want: choice{
				Key:   "key1",
				Value: "value1",
			},
		},
		"Empty": {
			value: `{}`,
			want: choice{},
		},
		"Failed": {
			value: `wrongval`,
			want: "invalid character",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()

			got, err := v.choice(test.value)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}