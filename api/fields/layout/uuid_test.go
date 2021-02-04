// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package layout

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalkerByUUID(t *testing.T) {

	id := uuid.New()

	repeater := []domain.Field{
		{UUID: id},
	}

	nested := []domain.Field{
		{UUID: uuid.New(), SubFields: &repeater},
	}

	doubleNested := []domain.Field{
		{UUID: uuid.New(), SubFields: &nested},
	}

	tt := map[string]struct {
		uuid  uuid.UUID
		field domain.Field
		found bool
		want  domain.Field
	}{
		"Normal Field": {
			uuid:  id,
			field: domain.Field{UUID: id},
			found: true,
			want:  domain.Field{UUID: id},
		},
		"Repeater": {
			uuid:  id,
			field: domain.Field{UUID: uuid.New(), SubFields: &repeater},
			found: true,
			want:  domain.Field{UUID: id},
		},
		"Nested Repeater": {
			uuid:  id,
			field: domain.Field{UUID: uuid.New(), SubFields: &nested},
			found: true,
			want:  domain.Field{UUID: id},
		},
		"Double Nested Repeater": {
			uuid:  id,
			field: domain.Field{UUID: uuid.New(), SubFields: &doubleNested},
			found: true,
			want:  domain.Field{UUID: id},
		},
		"Nil Sub Fields": {
			uuid:  id,
			field: domain.Field{UUID: uuid.New(), SubFields: nil},
			found: false,
			want:  domain.Field{},
		},
		"Flexible Content": {
			uuid: id,
			field: domain.Field{UUID: uuid.New(), Layouts: map[string]domain.FieldLayout{
				"layout": {
					SubFields: &repeater,
				},
			}},
			found: true,
			want:  domain.Field{UUID: id},
		},
		"Flexible Content Repeater": {
			uuid: id,
			field: domain.Field{UUID: uuid.New(), Layouts: map[string]domain.FieldLayout{
				"layout": {
					SubFields: &repeater,
				},
			}},
			found: true,
			want:  domain.Field{UUID: id},
		},
		"Flexible Content Nested Repeater": {
			uuid: id,
			field: domain.Field{UUID: uuid.New(), Layouts: map[string]domain.FieldLayout{
				"layout": {
					SubFields: &nested,
				},
			}},
			found: true,
			want:  domain.Field{UUID: id},
		},
		"Flexible Content Double Nested Repeater": {
			uuid: id,
			field: domain.Field{UUID: uuid.New(), Layouts: map[string]domain.FieldLayout{
				"layout": {
					SubFields: &doubleNested,
				},
			}},
			found: true,
			want:  domain.Field{UUID: id},
		},
		"Nil Flexible Content": {
			uuid:  id,
			field: domain.Field{UUID: uuid.New(), Layouts: map[string]domain.FieldLayout{}},
			found: false,
			want:  domain.Field{},
		},
		"Not Found": {
			uuid:  id,
			field: domain.Field{UUID: uuid.New()},
			found: false,
			want:  domain.Field{},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, found := walkerByUUID(test.uuid, test.field)
			assert.Equal(t, test.found, found)
			assert.Equal(t, test.want, got)
		})
	}
}
