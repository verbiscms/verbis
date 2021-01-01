package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GetRepeater(t *testing.T) {

	uniq := uuid.New()

	tt := map[string]struct {
		fields []domain.PostField
		key    string
		want   interface{}
	}{
		"Success": {
			fields: []domain.PostField{
				{Id: 1, Type: "repeater", UUID: uniq, Key: "key1", Value: 1, Parent: nil},
				{Id: 2, Type: "text", Key: "key2", Value: 2, Parent: &uniq,},
				{Id: 3, Type: "text", Key: "key3", Value: 3, Parent: &uniq},
				{Id: 4, Type: "text", Key: "key4", Value: 4, Parent: &uniq},
			},
			key:  "key1",
			want: Repeater{
				{Id: 2, Type: "text", Key: "key2", Value: 2, Parent: &uniq},
				{Id: 3, Type: "text", Key: "key3", Value: 3, Parent: &uniq},
				{Id: 4, Type: "text", Key: "key4", Value: 4, Parent: &uniq},
			},
		},
		"Sorted Index": {
			fields: []domain.PostField{
				{Id: 1, Type: "repeater", UUID: uniq, Key: "key1", Value: 1, Parent: nil},
				{Id: 2, Type: "text", Key: "key2", Value: 2, Parent: &uniq, Index: 2},
				{Id: 3, Type: "text", Key: "key3", Value: 3, Parent: &uniq, Index: 0},
				{Id: 4, Type: "text", Key: "key4", Value: 4, Parent: &uniq, Index: 1},
			},
			key:  "key1",
			want: Repeater{
				{Id: 3, Type: "text", Key: "key3", Value: 3, Parent: &uniq, Index: 0},
				{Id: 4, Type: "text", Key: "key4", Value: 4, Parent: &uniq, Index: 1},
				{Id: 2, Type: "text", Key: "key2", Value: 2, Parent: &uniq, Index: 2},
			},
		},
		"Not Found": {
			fields: []domain.PostField{},
			key:  "wrongval",
			want: "Fields.findByKey: no field exists with the key: wrongval",
		},
		"Invalid Type": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", UUID: uniq, Key: "key1", Value: 1, Parent: nil},
			},
			key:  "key1",
			want: "Fields.GetRepeater: field with the key: key1, is not a repeater",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			s := &Service{
				Fields: test.fields,
			}

			got, err := s.GetRepeater(test.key)
			if err != nil {
				assert.Equal(t, err.Error(), test.want)
				return
			}

			assert.Equal(t, test.want, got)
		})
	}
}

func TestRepeater_HasRows(t *testing.T) {

	tt := map[string]struct {
		repeater Repeater
		want     interface{}
	}{
		"With Rows": {
			repeater: Repeater{
				{Id: 1}, {Id: 2}, {Id: 3},
			},
			want: true,
		},
		"Without Rows": {
			repeater: Repeater{},
			want:     false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.want, test.repeater.HasRows())
		})
	}
}

func TestRepeater_SubField(t *testing.T) {

	repeater := Repeater{
		{Id: 1, Key: "test1", Value: 1},
		{Id: 2, Key: "test2", Value: 2},
		{Id: 3, Key: "test3", Value: 3},
	}

	tt := map[string]struct {
		key  string
		want interface{}
	}{
		"Found": {
			key:  "test1",
			want: 1,
		},
		"Not Found": {
			key:  "wrongval",
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.want, repeater.SubField(test.key))
		})
	}
}

func TestRepeater_First(t *testing.T) {

	tt := map[string]struct {
		repeater Repeater
		want     interface{}
	}{
		"Found": {
			repeater: Repeater{
				{Id: 1, Key: "test1", Value: 1},
				{Id: 2, Key: "test2", Value: 2},
				{Id: 3, Key: "test3", Value: 3},
			},
			want: domain.PostField{Id: 1, Key: "test1", Value: 1},
		},
		"Not Found": {
			repeater: Repeater{},
			want:     nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.want, test.repeater.First())
		})
	}
}

func TestRepeater_Last(t *testing.T) {

	tt := map[string]struct {
		repeater Repeater
		want     interface{}
	}{
		"Found": {
			repeater: Repeater{
				{Id: 1, Key: "test1", Value: 1},
				{Id: 2, Key: "test2", Value: 2},
				{Id: 3, Key: "test3", Value: 3},
			},
			want: domain.PostField{Id: 3, Key: "test3", Value: 3},
		},
		"Not Found": {
			repeater: Repeater{},
			want:     nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.want, test.repeater.Last())
		})
	}
}
