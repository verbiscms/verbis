// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import "github.com/ainsleyclark/verbis/api/domain"

func (t *ResolverTestSuite) TestValue_Number() {
	tt := map[string]struct {
		value domain.FieldValue
		want  interface{}
	}{
		"Success": {
			value: "1",
			want:  int64(1),
		},
		"Large": {
			value: "99999999999999999",
			want:  int64(99999999999999999),
		},
		"Bad Cast": {
			value: "wrongval",
			want:  "unable to cast",
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

func (t *ResolverTestSuite) TestValue_NumberResolve() {
	tt := map[string]struct {
		field domain.PostField
		want  domain.PostField
	}{
		"Number": {
			field: domain.PostField{OriginalValue: "999", Type: "number"},
			want:  domain.PostField{OriginalValue: "999", Type: "number", Value: int64(999)},
		},
		"Range": {
			field: domain.PostField{OriginalValue: "999", Type: "range"},
			want:  domain.PostField{OriginalValue: "999", Type: "range", Value: int64(999)},
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
