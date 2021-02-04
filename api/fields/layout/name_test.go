// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package layout

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalkerByName(t *testing.T) {

	repeater := []domain.Field{
		{Name: "test"},
	}

	nested := []domain.Field{
		{Name: "wrong", SubFields: &repeater},
	}

	doubleNested := []domain.Field{
		{Name: "wrong", SubFields: &nested},
	}

	tt := map[string]struct {
		field domain.Field
		found bool
		want  domain.Field
	}{
		"Normal Field": {
			field: domain.Field{Name: "test"},
			found: true,
			want:  domain.Field{Name: "test"},
		},
		"Repeater": {
			field: domain.Field{Name: "wrong", SubFields: &repeater},
			found: true,
			want:  domain.Field{Name: "test"},
		},
		"Nested Repeater": {
			field: domain.Field{Name: "wrong", SubFields: &nested},
			found: true,
			want:  domain.Field{Name: "test"},
		},
		"Double Nested Repeater": {
			field: domain.Field{Name: "wrong", SubFields: &doubleNested},
			found: true,
			want:  domain.Field{Name: "test"},
		},
		"Nil Sub Fields": {
			field: domain.Field{Name: "wrong", SubFields: nil},
			found: false,
			want:  domain.Field{},
		},
		"Flexible Content": {
			field: domain.Field{Name: "wrong", Layouts: map[string]domain.FieldLayout{
				"layout": {
					SubFields: &repeater,
				},
			}},
			found: true,
			want:  domain.Field{Name: "test"},
		},
		"Flexible Content Repeater": {
			field: domain.Field{Name: "wrong", Layouts: map[string]domain.FieldLayout{
				"layout": {
					SubFields: &repeater,
				},
			}},
			found: true,
			want:  domain.Field{Name: "test"},
		},
		"Flexible Content Nested Repeater": {
			field: domain.Field{Name: "wrong", Layouts: map[string]domain.FieldLayout{
				"layout": {
					SubFields: &nested,
				},
			}},
			found: true,
			want:  domain.Field{Name: "test"},
		},
		"Flexible Content Double Nested Repeater": {
			field: domain.Field{Name: "wrong", Layouts: map[string]domain.FieldLayout{
				"layout": {
					SubFields: &doubleNested,
				},
			}},
			found: true,
			want:  domain.Field{Name: "test"},
		},
		"Not Found": {
			field: domain.Field{Name: "wrong"},
			found: false,
			want:  domain.Field{},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, found := walkerByName("test", test.field)
			assert.Equal(t, test.found, found)
			assert.Equal(t, test.want, got)
		})
	}
}
