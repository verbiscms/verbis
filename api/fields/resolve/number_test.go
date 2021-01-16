package resolve

import "github.com/ainsleyclark/verbis/api/domain"

func (t *ResolverTestSuite) TestFieldValue_Number() {

	tt := map[string]struct {
		value  domain.FieldValue
		want   interface{}
	}{
		"Success": {
			value: "1",
			want: int64(1),
		},
		"Large": {
			value: "99999999999999999",
			want: int64(99999999999999999),
		},
		"Bad Cast": {
			value: "wrongval",
			want: "unable to cast",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			v := t.GetValue()

			got, err := v.number(test.value)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}