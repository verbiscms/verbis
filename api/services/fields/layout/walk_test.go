// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package layout

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/verbiscms/verbis/api/domain"
	"testing"
)

func TestByUUID(t *testing.T) {
	id := uuid.New()
	field := domain.Field{UUID: id, Type: "text"}
	fields := domain.Fields{field}

	tt := map[string]struct {
		uuid   uuid.UUID
		groups domain.FieldGroups
		want   interface{}
	}{
		"Found": {
			uuid:   id,
			groups: domain.FieldGroups{{Fields: fields}},
			want:   field,
		},
		"Not Found": {
			uuid:   uuid.New(),
			groups: domain.FieldGroups{{Fields: fields}},
			want:   "unable to find field with UUID of",
		},
		"No Layouts": {
			uuid:   uuid.New(),
			groups: nil,
			want:   "no groups exist, unable to range over groups and find fields",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ByUUID(test.uuid, test.groups)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestByName(t *testing.T) {
	field := domain.Field{Name: "test", Type: "text"}
	fields := domain.Fields{field}

	tt := map[string]struct {
		name   string
		groups domain.FieldGroups
		want   interface{}
	}{
		"Found": {
			name:   "test",
			groups: domain.FieldGroups{{Fields: fields}},
			want:   field,
		},
		"Not Found": {
			name:   "wrong",
			groups: domain.FieldGroups{{Fields: fields}},
			want:   "unable to find field with name of",
		},
		"No Layouts": {
			name:   "test",
			groups: nil,
			want:   "no groups exist, unable to range over groups and find fields",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ByName(test.name, test.groups)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
