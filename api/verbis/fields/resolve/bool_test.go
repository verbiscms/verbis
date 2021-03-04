// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import "github.com/ainsleyclark/verbis/api/domain"

func (t *ResolverTestSuite) TestValue_Checkbox() {
	tt := map[string]struct {
		value domain.FieldValue
		want  interface{}
	}{
		"True": {
			value: domain.FieldValue("true"),
			want:  true,
		},
		"False": {
			value: domain.FieldValue("false"),
			want:  false,
		},
		"Failed": {
			value: `wrongval`,
			want:  "invalid syntax",
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
