package walker

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByUUID(t *testing.T) {

	id := uuid.New()
	field := domain.Field{UUID: id, Type: "text"}
	fields := []domain.Field{field}

	tt := map[string]struct {
		uuid   uuid.UUID
		groups []domain.FieldGroup
		want   interface{}
	}{
		"Found": {
			uuid:   id,
			groups: []domain.FieldGroup{{Fields: &fields}},
			want:   field,
		},
		"Not Found": {
			uuid:   uuid.New(),
			groups: []domain.FieldGroup{{Fields: &fields}},
			want:   fmt.Sprintf("unable to find field with UUID of"),
		},
		"No Layouts": {
			uuid:   uuid.New(),
			groups: nil,
			want:   fmt.Sprintf("no groups exist, unable to range over groups and find fields"),
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
	fields := []domain.Field{field}

	tt := map[string]struct {
		name   string
		groups []domain.FieldGroup
		want   interface{}
	}{
		"Found": {
			name:   "test",
			groups: []domain.FieldGroup{{Fields: &fields}},
			want:   field,
		},
		"Not Found": {
			name:   "wrong",
			groups: []domain.FieldGroup{{Fields: &fields}},
			want:   fmt.Sprintf("unable to find field with name of"),
		},
		"No Layouts": {
			name:   "test",
			groups: nil,
			want:   fmt.Sprintf("no groups exist, unable to range over groups and find fields"),
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
