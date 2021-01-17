package resolve

import "github.com/ainsleyclark/verbis/api/domain"

func (t *ResolverTestSuite) TestValue_Choice() {

	tt := map[string]struct {
		value domain.FieldValue
		want  interface{}
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
			want:  choice{},
		},
		"Failed": {
			value: `wrongval`,
			want:  "invalid character",
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

func (t *ResolverTestSuite) TestValue_ChoiceResolve() {

	tt := map[string]struct {
		field domain.PostField
		want  domain.PostField
	}{
		"Button Group Value": {
			field: domain.PostField{OriginalValue: "test", Key: "value", Type: "button_group"},
			want:  domain.PostField{OriginalValue: "test", Key: "value", Type: "button_group", Value: "test"},
		},
		"Button Group Key": {
			field: domain.PostField{OriginalValue: "test", Key: "key", Type: "button_group"},
			want:  domain.PostField{OriginalValue: "test", Key: "key", Type: "button_group", Value: "test"},
		},
		"Button Group Map": {
			field: domain.PostField{OriginalValue: `{"key": "key1", "value": "value1"}`, Key: "map", Type: "button_group"},
			want:  domain.PostField{OriginalValue: `{"key": "key1", "value": "value1"}`, Key: "map", Type: "button_group", Value: choice{
				Key:   "key1",
				Value: "value1",
			}},
		},
		"Radio Value": {
			field: domain.PostField{OriginalValue: "test", Key: "value", Type: "radio"},
			want:  domain.PostField{OriginalValue: "test", Key: "value", Type: "radio", Value: "test"},
		},
		"Radio Key": {
			field: domain.PostField{OriginalValue: "test", Key: "key", Type: "radio"},
			want:  domain.PostField{OriginalValue: "test", Key: "key", Type: "radio", Value: "test"},
		},
		"Radio Map": {
			field: domain.PostField{OriginalValue: `{"key": "key1", "value": "value1"}`, Key: "map", Type: "radio"},
			want:  domain.PostField{OriginalValue: `{"key": "key1", "value": "value1"}`, Key: "map", Type: "radio", Value: choice{
				Key:   "key1",
				Value: "value1",
			}},
		},
		"Select Value": {
			field: domain.PostField{OriginalValue: "test", Key: "value", Type: "select"},
			want:  domain.PostField{OriginalValue: "test", Key: "value", Type: "select", Value: "test"},
		},
		"Select Key": {
			field: domain.PostField{OriginalValue: "test", Key: "key", Type: "select"},
			want:  domain.PostField{OriginalValue: "test", Key: "key", Type: "select", Value: "test"},
		},
		"Select Map": {
			field: domain.PostField{OriginalValue: `{"key": "key1", "value": "value1"}`, Key: "map", Type: "select"},
			want:  domain.PostField{OriginalValue: `{"key": "key1", "value": "value1"}`, Key: "map", Type: "select", Value: choice{
				Key:   "key1",
				Value: "value1",
			}},
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
