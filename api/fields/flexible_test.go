package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GetFlexible(t *testing.T) {

	uniq := uuid.New()

	tt := map[string]struct {
		fields []domain.PostField
		key    string
		want   interface{}
	}{
		"Success": {
			fields: []domain.PostField{
				{Id: 1, Type: "flexible", UUID: uniq, Key: "key1", Value: 1, Parent: nil},
				{Id: 2, Type: "text", Key: "key2", Value: 2, Parent: &uniq},
				{Id: 3, Type: "text", Key: "key3", Value: 3, Parent: &uniq},
				{Id: 4, Type: "text", Key: "key4", Value: 4, Parent: &uniq},
			},
			key: "key1",
			want: Repeater{
				{Id: 2, Type: "text", Key: "key2", Value: 2, Parent: &uniq},
				{Id: 3, Type: "text", Key: "key3", Value: 3, Parent: &uniq},
				{Id: 4, Type: "text", Key: "key4", Value: 4, Parent: &uniq},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			s := &Service{
				Fields: test.fields,
			}

			got, err := s.GetFlexible(test.key)
			if err != nil {
				assert.Equal(t, err.Error(), test.want)
				return
			}

			assert.Equal(t, test.want, got)
		})
	}
}
