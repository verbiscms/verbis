package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
)

func (t *FieldTestSuite) TestService_GetField() {

	tt := map[string]struct {
		fields []domain.PostField
		key    string
		mock   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)
		args   []interface{}
		want   interface{}
	}{
		"Success": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "key1", Value: 1},
			},
			key:  "key1",
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args: nil,
			want: 1,
		},
		"No Field": {
			fields: nil,
			key:    "wrongval",
			mock:   func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {},
			args:   nil,
			want:   "no field exists with the name: wrongval",
		},
		"Post": {
			fields: []domain.PostField{
				{Id: 1, Type: "text", Name: "key1"},
			},
			key: "key2",
			mock: func(f *mocks.FieldsRepository, c *mocks.CategoryRepository) {
				f.On("GetByPost", 2).Return([]domain.PostField{{Id: 2, Type: "text", Name: "key2", Value: 2}}, nil)
			},
			args: []interface{}{2},
			want: 2,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetMockService(test.fields, test.mock)

			got, err := s.GetField(test.key, test.args...)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}
